package service

import (
	"go_sample/app/domain/model"
	"go_sample/app/domain/repository"
)

type userService struct {
	tx repository.Transaction
}

// User作成に関連する処理をまとめたservice
func (s *userService) Create(user model.User) (*model.UserID, error) {
	chat_for_all, err := s.tx.Chat().FindByName(model.AllChatName)
	if err != nil {
		return nil, err
	}

	group_for_free, err := s.tx.Group().FindByName(model.FreeGroupName)
	if err != nil {
		return nil, err
	}

	if err := s.tx.Group().IncreaseNumberOfMembers(group_for_free.ID, 1); err != nil {
		return nil, err
	}

	user.Group = *group_for_free

	created_user_id, err := s.tx.User().Store(user)
	if err != nil {
		return nil, err
	}

	if err := s.tx.Chat().Join(*created_user_id, chat_for_all.ID); err != nil {
		return nil, err
	}

	if err := s.tx.Chat().Join(*created_user_id, group_for_free.Chat.ID); err != nil {
		return nil, err
	}

	return created_user_id, nil
}

// UserのGroup所属に関連する処理をまとめたservice
func (s *userService) JoinGroup(user model.User, group model.Group) error {
	// Group Chatの変更
	if err := s.tx.Chat().Leave(user.ID, user.Group.Chat.ID); err != nil {
		return err
	}

	if err := s.tx.Chat().Join(user.ID, group.Chat.ID); err != nil {
		return err
	}

	// Groupの変更と、変更前後のGroupのNumberOfMembersを変更
	if err := s.tx.Group().DecreaseNumberOfMembers(user.Group.ID, 1); err != nil {
		return err
	}

	user.Group = group
	if _, err := s.tx.User().Update(user); err != nil {
		return err
	}

	if err := s.tx.Group().IncreaseNumberOfMembers(user.Group.ID, 1); err != nil {
		return err
	}

	return nil
}

// UserのGroup脱退に関連する処理をまとめたservice
func (s *userService) LeaveGroup(user model.User) error {
	// 未所属向けGroupへの移動となる

	group_for_free, err := s.tx.Group().FindByName(model.FreeGroupName)
	if err != nil {
		return err
	}

	// Group Chatの変更
	if err := s.tx.Chat().Leave(user.ID, user.Group.Chat.ID); err != nil {
		return err
	}

	if err := s.tx.Chat().Join(user.ID, group_for_free.Chat.ID); err != nil {
		return err
	}

	// Groupの変更と、変更前後のGroupのNumberOfMembersを変更
	if err := s.tx.Group().DecreaseNumberOfMembers(user.Group.ID, 1); err != nil {
		return err
	}

	user.Group = *group_for_free
	if _, err := s.tx.User().Update(user); err != nil {
		return err
	}

	if err := s.tx.Group().IncreaseNumberOfMembers(user.Group.ID, 1); err != nil {
		return err
	}

	return err
}

// User削除に関連する処理をまとめたservice
func (s *userService) Delete(user model.User) error {
	if err := s.tx.User().Delete(user.ID); err != nil {
		return err
	}

	// Userが所属していたGroupのNumberOfMembersを減らす
	if err := s.tx.Group().DecreaseNumberOfMembers(user.Group.ID, 1); err != nil {
		return err
	}

	// UserのChatMessageを削除
	if err := s.tx.ChatMessage().DeleteByCreatedByID(user.ID); err != nil {
		return err
	}

	// DirectMessageを削除
	if err := s.tx.DirectMessage().DeleteByFromUserID(user.ID); err != nil {
		return err
	}
	if err := s.tx.DirectMessage().DeleteByToUserID(user.ID); err != nil {
		return err
	}

	return nil
}
