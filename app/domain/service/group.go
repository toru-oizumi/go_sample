package service

import (
	"go_sample/app/domain/model"
	"go_sample/app/domain/repository"
)

type groupService struct {
	tx repository.Transaction
}

// Group作成に関連する処理をまとめたservice
func (s *groupService) Create(group model.Group) (*model.GroupID, error) {
	chat := model.Chat{
		// Group向けChatのNameは、GroupのNameと同じ値を使用する
		Name: model.ChatName(group.Name),
	}

	created_chat_id, err := s.tx.Chat().Store(chat)
	if err != nil {
		return nil, err
	}

	created_chat, err := s.tx.Chat().FindByID(*created_chat_id)
	if err != nil {
		return nil, err
	}

	group.Chat = *created_chat

	created_group_id, err := s.tx.Group().Store(group)
	if err != nil {
		return nil, err
	}

	return created_group_id, nil
}

// Group作成に関連する処理をまとめたservice
func (s *groupService) Update(group model.Group) (*model.GroupID, error) {

	updated_group_id, err := s.tx.Group().Update(group)
	if err != nil {
		return nil, err
	}

	// groupに紐付くChatを取得する為に、FindByIDせざるを得ない
	updated_group, err := s.tx.Group().FindByID(*updated_group_id)
	if err != nil {
		return nil, err
	}

	chat := updated_group.Chat
	// GroupとChatの名前は一致させる
	chat.Name = model.ChatName(updated_group.Name)
	_, err = s.tx.Chat().Update(chat)
	if err != nil {
		return nil, err
	}

	return updated_group_id, nil
}

// Group削除に関連する処理をまとめたservice
func (s *groupService) Delete(id model.GroupID) error {
	// Userのチャット所属をfreeに
	// UserのGroupをgroup_for_free

	group_for_free, err := s.tx.Group().FindByName(model.FreeGroupName)
	if err != nil {
		return err
	}

	group, err := s.tx.Group().FindByID(id)
	if err != nil {
		return err
	}

	// 所属していたUserのGroupを更新
	var user_ids []model.UserID
	users, err := s.tx.User().List(repository.UserFilter{GroupID: id})
	if err != nil {
		return err
	}
	for _, v := range users {
		user_ids = append(user_ids, v.ID)
	}

	// Userのチャット所属をfreeに
	// UserのGroupをgroup_for_free
	if len(user_ids) > 0 {
		if err := s.tx.Chat().LeaveByUserIDs(user_ids, group.Chat.ID); err != nil {
			return err
		}
		if err := s.tx.User().UpdateGroupByIDs(user_ids, *group_for_free); err != nil {
			return err
		}
		if err := s.tx.Chat().JoinByUserIDs(user_ids, group_for_free.Chat.ID); err != nil {
			return err
		}
		if err := s.tx.Group().IncreaseNumberOfMembers(group_for_free.ID, uint(len(user_ids))); err != nil {
			return err
		}
	}

	if err := s.tx.Group().Delete(id); err != nil {
		return err
	}

	// Group向けChatの削除
	if err := s.tx.Chat().Delete(group.Chat.ID); err != nil {
		return err
	}

	// Group向けChatMessageの削除
	if err := s.tx.ChatMessage().DeleteByChatID(group.Chat.ID); err != nil {
		return err
	}

	return nil
}
