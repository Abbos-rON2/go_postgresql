package service

import (
	"github.com/abbos-ron2/go_postgresql/api/models"
	"github.com/abbos-ron2/go_postgresql/storage"
)

type TeacherServiceI interface {
	CreateTeacher(entity models.CreateTeacher) (id string, err error)
	UpdateTeacher(guid string, entity models.UpdateTeacher) (resp models.Teacher, err error)
	GetTeacher(guid string) (resp models.Teacher, err error)
	GetAllTeachers() (resp []models.Teacher, err error)
	DeleteTeacher(guid string) error
}

type teacherService struct {
	storage storage.StorageI
}

func NewTeacherService(storage storage.StorageI) *teacherService {
	return &teacherService{storage: storage}
}

func (s *teacherService) CreateTeacher(entity models.CreateTeacher) (guid string, err error) {
	// id, err = s.storage.UserRepo().Create(entity)
	// s.storage.Queries().CreateUser(context.Background(), queries.CreateUserParams{
	// 	Username: entity.Username,
	// 	Password: entity.Password,
	// })

	return guid, err
}

func (s *teacherService) UpdateTeacher(guid string, entity models.UpdateTeacher) (resp models.Teacher, err error) {
	// id, err = s.storage.UserRepo().Update(entity)
	return resp, err
}

func (s *teacherService) GetTeacher(id string) (resp models.Teacher, err error) {
	// resp, err = s.storage.UserRepo().Get(id)
	return resp, err
}

func (s *teacherService) GetAllTeachers() (resp []models.Teacher, err error) {
	// resp, err = s.storage.UserRepo().Get(id)
	return resp, err
}

func (s *teacherService) DeleteTeacher(guid string) error {
	// resp, err = s.storage.UserRepo().Get(id)
	return nil
}
