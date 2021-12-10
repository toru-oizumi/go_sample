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

type DirectMessageInteractor struct {
	Connection repository.Connection
	Presenter  presenter.DirectMessagePresenter
}

func (i *DirectMessageInteractor) FindMessages(request input.FindDirectMessagesRequest) ([]output.DirectMessageResponse, error) {
	if messages, err := i.Connection.DirectMessage().List(
		repository.DirectMessageFilter{FromUserID: request.FromUserID, ToUserID: request.ToUserID},
	); err != nil {
		return nil, err
	} else {
		return i.Presenter.BuildDirectMessagesResponse(messages)
	}
}

func (i *DirectMessageInteractor) CreateMessage(request input.CreateDirectMessageRequest) (*output.DirectMessageResponse, error) {
	message := model.DirectMessage{
		FromUser: model.User{ID: request.FromUserID},
		ToUser:   model.User{ID: request.ToUserID},
		Body:     request.Message,
	}

	if created_message, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			created_message_id, err := tx.DirectMessage().Store(message)
			if err != nil {
				return nil, err
			}

			created_message, err := tx.DirectMessage().FindByID(*created_message_id)
			if err != nil {
				return nil, err
			}
			return *created_message, nil
		},
	); err != nil {
		return nil, err
	} else {
		parsed_message, _ := created_message.(model.DirectMessage)
		return i.Presenter.BuildDirectMessageResponse(parsed_message)
	}
}

func (i *DirectMessageInteractor) UpdateMessage(request input.UpdateDirectMessageRequest) (*output.DirectMessageResponse, error) {
	message, err := i.Connection.DirectMessage().FindByID(request.ID)
	if err != nil {
		return nil, err
	}

	if message.FromUser.ID != request.FromUserID {
		return nil, util_error.NewErrBadRequest("invalid FromUserID")
	}

	if message.ToUser.ID != request.ToUserID {
		return nil, util_error.NewErrBadRequest("invalid ToUserID")
	}

	message.Body = request.Message

	if updated_message, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			updated_message_id, err := tx.DirectMessage().Update(*message)
			if err != nil {
				return nil, err
			}

			updated_message, err := tx.DirectMessage().FindByID(*updated_message_id)
			if err != nil {
				return nil, err
			}
			return *updated_message, nil
		},
	); err != nil {
		return nil, err
	} else {
		parsed_message, _ := updated_message.(model.DirectMessage)
		return i.Presenter.BuildDirectMessageResponse(parsed_message)
	}
}

func (i *DirectMessageInteractor) DeleteMessage(request input.DeleteDirectMessageRequest) (*output.DeletedDirectMessageResponse, error) {
	// 返す結果はRepositoryから取得した結果に依存しないので、先にrequestから生成しておく
	result, err := i.Presenter.BuildDeletedDirectMessageResponse(
		model.DirectMessage{
			ID: request.ID,
		})
	if err != nil {
		return nil, err
	}

	message, err := i.Connection.DirectMessage().FindByID(request.ID)
	if err != nil {
		// 冪等性を重視して、削除の場合はrecord not foundエラーにしない
		if errors.As(err, &util_error.ErrRecordNotFound{}) {

			return result, nil
		}
		return nil, err
	}

	if message.FromUser.ID != request.FromUserID {
		return nil, util_error.NewErrBadRequest("invalid FromUserID")
	}

	if message.ToUser.ID != request.ToUserID {
		return nil, util_error.NewErrBadRequest("invalid ToUserID")
	}

	if _, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			err := tx.DirectMessage().Delete(request.ID)
			return nil, err
		},
	); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
