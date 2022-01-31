package service

import (
	"github.com/abbos-ron2/go_postgresql/api/models"
	"github.com/abbos-ron2/go_postgresql/storage"
)

type UserServiceI interface {
	CreateUser(entity models.CreateUser) (id string, err error)
	UpdateUser(entity models.UpdateUser) (id string, err error)
	GetUser(id string) (resp models.User, err error)
}

type userService struct {
	storage storage.StorageI
}

func NewUserService(storage storage.StorageI) *userService {
	return &userService{storage: storage}
}

func (s *userService) CreateUser(entity models.CreateUser) (id string, err error) {
	id, err = s.storage.UserRepo().Create(entity)

	return id, err
}

func (s *userService) UpdateUser(entity models.UpdateUser) (id string, err error) {
	id, err = s.storage.UserRepo().Update(entity)
	return id, err
}

func (s *userService) GetUser(id string) (resp models.User, err error) {
	resp, err = s.storage.UserRepo().Get(id)
	return resp, err
}
