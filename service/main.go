package service

import (
	"github.com/abbos-ron2/go_postgresql/storage"
)

type ServiceManager interface {
	UserService() UserServiceI
}

type serviceManager struct {
	storage     storage.StorageI
	userService UserServiceI
}

func NewServiceManager(s storage.StorageI) ServiceManager {
	return serviceManager{
		storage:     s,
		userService: NewUserService(s),
	}
}

func (s serviceManager) UserService() UserServiceI {
	return &userService{storage: s.storage}
}
