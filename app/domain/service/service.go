package service

import "go_sample/app/domain/repository"

type DomainService struct {
	User  userService
	Group groupService
}

func NewDomainService(tx repository.Transaction) DomainService {
	return DomainService{
		User:  userService{tx: tx},
		Group: groupService{tx: tx},
	}
}
