package service

import (
	"github.com/abbos-ron2/go_postgresql/storage"
)

type ServiceManager interface {
	UserService() UserServiceI
	TeacherService() TeacherServiceI
}

type serviceManager struct {
	storage        storage.StorageI
	userService    UserServiceI
	teacherService TeacherServiceI
}

func NewServiceManager(s storage.StorageI) ServiceManager {
	return serviceManager{
		storage:        s,
		userService:    NewUserService(s),
		teacherService: NewTeacherService(s),
	}
}

func (s serviceManager) UserService() UserServiceI {
	return &userService{storage: s.storage}
}

func (s serviceManager) TeacherService() TeacherServiceI {
	return &teacherService{storage: s.storage}
}
