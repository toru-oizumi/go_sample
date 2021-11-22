package interactor

import (
	"go_sample/app/application/input"
	"go_sample/app/application/output"
	"go_sample/app/application/presenter"
	"go_sample/app/domain/model"
	"go_sample/app/domain/repository"
)

type ChatInteractor struct {
	Connection repository.Connection
	Presenter  presenter.ChatPresenter
}

func (i *ChatInteractor) FindAll(request input.FindChatsRequest) ([]output.ChatResponse, error) {
	if chats, err := i.Connection.Chat().List(repository.ChatFilter{UserID: model.UserID("1")}); err != nil {
		return nil, err
	} else {
		return i.Presenter.BuildChatsResponse(chats)
	}
}

func (i *ChatInteractor) FindMessages(request input.FindChatMessagesByIDRequest) ([]output.ChatMessageResponse, error) {
	if messages, err := i.Connection.ChatMessage().List(repository.ChatMessageFilter{ChatID: request.ChatID}); err != nil {
		return nil, err
	} else {
		return i.Presenter.BuildChatMessagesResponse(messages)
	}
}

func (i *ChatInteractor) CreateMessage(request input.CreateChatMessageRequest) (*output.ChatMessageResponse, error) {
	chat_message := model.ChatMessage{
		ChatID:       request.ChatID,
		CreatedBy:    request.UserID,
		Body:         request.Message,
		IsPrivileged: false,
	}

	if created_chat_message, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			if created_chat_message, err := tx.ChatMessage().Store(chat_message); err != nil {
				return nil, err
			} else {
				return created_chat_message, nil
			}
		},
	); err != nil {
		return nil, err
	} else {
		parsed_chat_message, _ := created_chat_message.(model.ChatMessage)
		return i.Presenter.BuildChatMessageResponse(parsed_chat_message)
	}
}

func (i *ChatInteractor) UpdateMessage(request input.UpdateChatMessageRequest) (*output.ChatMessageResponse, error) {
	chat_message := model.ChatMessage{
		ID:        request.ChatMessageID,
		ChatID:    request.ChatID,
		CreatedBy: request.UserID,
		Body:      request.Message,
	}

	if updated_chat_message, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			if updated_chat_message, err := tx.ChatMessage().Update(chat_message); err != nil {
				return nil, err
			} else {
				return updated_chat_message, nil
			}
		},
	); err != nil {
		return nil, err
	} else {
		parsed_chat_message, _ := updated_chat_message.(model.ChatMessage)
		return i.Presenter.BuildChatMessageResponse(parsed_chat_message)
	}
}

func (i *ChatInteractor) DeleteMessage(request input.DeleteChatMessageRequest) error {
	if _, err := i.Connection.ChatMessage().FindByID(request.ChatMessageID); err != nil {
		return err
	}
	if _, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			if err := tx.ChatMessage().Delete(request.ChatMessageID); err != nil {
				return nil, err
			} else {
				return nil, nil
			}
		},
	); err != nil {
		return err
	} else {
		return nil
	}
}
