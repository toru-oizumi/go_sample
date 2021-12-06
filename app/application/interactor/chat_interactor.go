package interactor

import (
	"errors"
	"go_sample/app/application/input"
	"go_sample/app/application/output"
	"go_sample/app/application/presenter"
	"go_sample/app/domain/model"
	"go_sample/app/domain/repository"
	util_error "go_sample/app/utility/error"
)

type ChatInteractor struct {
	Connection repository.Connection
	Presenter  presenter.ChatPresenter
}

func (i *ChatInteractor) FindAll(request input.FindChatsRequest) ([]output.ChatResponse, error) {
	if chats, err := i.Connection.Chat().List(repository.ChatFilter{UserID: request.UserID}); err != nil {
		return nil, err
	} else {
		return i.Presenter.BuildChatsResponse(chats)
	}
}

func (i *ChatInteractor) FindChatMembers(request input.FindChatMembersRequest) (output.ChatMembersResponse, error) {
	if members, err := i.Connection.Chat().FindMembersByID(request.ChatID); err != nil {
		return nil, err
	} else {
		return i.Presenter.BuildChatMembersResponse(members)
	}
}

func (i *ChatInteractor) FindMessages(request input.FindChatMessagesRequest) ([]output.ChatMessageResponse, error) {
	if ok, _ := i.Connection.Chat().Exists(request.ChatID); !ok {
		return nil, util_error.NewErrRecordNotFound()
	}

	if ok, err := i.Connection.Chat().DoseJoinChat(request.UserID, request.ChatID); err != nil {
		return nil, err
	} else {
		if !ok {
			return nil, util_error.NewErrBadRequest("dose not join this chat")
		}
	}

	if messages, err := i.Connection.ChatMessage().List(repository.ChatMessageFilter{ChatID: request.ChatID}); err != nil {
		return nil, err
	} else {
		return i.Presenter.BuildChatMessagesResponse(messages)
	}
}

func (i *ChatInteractor) CreateMessage(request input.CreateChatMessageRequest) (*output.ChatMessageResponse, error) {
	if ok, _ := i.Connection.Chat().Exists(request.ChatID); !ok {
		return nil, util_error.NewErrRecordNotFound()
	}

	chat_message := model.ChatMessage{
		ChatID:       request.ChatID,
		CreatedBy:    request.UserID,
		Body:         request.Message,
		IsPrivileged: false,
	}

	if created_chat_message, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			created_chat_message_id, err := tx.ChatMessage().Store(chat_message)
			if err != nil {
				return nil, err
			}

			created_chat_message, err := tx.ChatMessage().FindByID(*created_chat_message_id)
			if err != nil {
				return nil, err
			}
			return *created_chat_message, nil
		},
	); err != nil {
		return nil, err
	} else {
		parsed_chat_message, _ := created_chat_message.(model.ChatMessage)
		return i.Presenter.BuildChatMessageResponse(parsed_chat_message)
	}
}

func (i *ChatInteractor) UpdateMessage(request input.UpdateChatMessageRequest) (*output.ChatMessageResponse, error) {
	message, err := i.Connection.ChatMessage().FindByID(request.ChatMessageID)
	if err != nil {
		return nil, err
	}

	if message.ChatID != request.ChatID {
		return nil, util_error.NewErrBadRequest("invalid ChatID")
	}

	if message.CreatedBy != request.UserID {
		return nil, util_error.NewErrBadRequest("invalid UserID")
	}

	chat_message := model.ChatMessage{
		ID:        request.ChatMessageID,
		ChatID:    request.ChatID,
		CreatedBy: request.UserID,
		Body:      request.Message,
	}

	if updated_chat_message, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			updated_chat_message_id, err := tx.ChatMessage().Update(chat_message)
			if err != nil {
				return nil, err
			}

			updated_chat_message, err := tx.ChatMessage().FindByID(*updated_chat_message_id)
			if err != nil {
				return nil, err
			}
			return *updated_chat_message, nil
		},
	); err != nil {
		return nil, err
	} else {
		parsed_chat_message, _ := updated_chat_message.(model.ChatMessage)
		return i.Presenter.BuildChatMessageResponse(parsed_chat_message)
	}
}

func (i *ChatInteractor) DeleteMessage(request input.DeleteChatMessageRequest) (*output.DeletedChatMessageResponse, error) {
	// 返す結果はRepositoryから取得した結果に依存しないので、先にrequestから生成しておく
	result, err := i.Presenter.BuildDeletedChatMessageResponse(
		model.ChatMessage{
			ID:     request.ChatMessageID,
			ChatID: request.ChatID,
		})
	if err != nil {
		return nil, err
	}

	message, err := i.Connection.ChatMessage().FindByID(request.ChatMessageID)
	if err != nil {
		// 冪等性を重視して、削除の場合はrecord not foundエラーにしない
		if errors.As(err, &util_error.ErrRecordNotFound{}) {

			return result, nil
		}
		return nil, err
	}

	if message.ChatID != request.ChatID {
		return nil, util_error.NewErrBadRequest("invalid ChatID")
	}

	if message.CreatedBy != request.UserID {
		return nil, util_error.NewErrBadRequest("invalid UserID")
	}

	if _, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			err := tx.ChatMessage().Delete(request.ChatMessageID)
			return nil, err
		},
	); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
