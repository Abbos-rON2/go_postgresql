package service

import (
	"context"
	"errors"
	"time"

	"github.com/abbos-ron2/go_postgresql/api/models"
	"github.com/abbos-ron2/go_postgresql/storage"
	queries "github.com/abbos-ron2/go_postgresql/storage/sqlc"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type UserServiceI interface {
	Login(models.LoginModel) (string, error)
	CreateUser(entity models.CreateUser) (id string, err error)
	UpdateUserPassword(entity models.UpdateUserPassword) (id string, err error)
	GetUser(guid string) (resp models.User, err error)
	GetAllUsers(limit, offset int) (resp []models.User, err error)
}

type userService struct {
	storage storage.StorageI
}

func NewUserService(storage storage.StorageI) *userService {
	return &userService{storage: storage}
}
func (s *userService) Login(entity models.LoginModel) (tokenString string, err error) {
	user, err := s.storage.Queries().LoginUser(context.Background(), entity.Username)
	if err != nil {
		return "", errors.New("user not found")
	}
	if !user.IsActive {
		return "", errors.New("user is deactivated")
	}
	if user.Password != entity.Password {
		return "", errors.New("invalid password")
	}

	claims := jwt.MapClaims{}
	claims["user_id"] = user.Guid.String()
	claims["role"] = string(user.Role)
	claims["issued_at"] = time.Now().Unix()
	claims["expires_at"] = time.Now().AddDate(0, 0, 1)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString([]byte("root"))
	if err != nil {
		return "", err
	}

	return tokenString, err
}
func (s *userService) CreateUser(entity models.CreateUser) (id string, err error) {
	birthDate, err := time.Parse("2006-01-02", entity.BirthDate)
	if err != nil {
		return id, err
	}

	uuid, err := s.storage.Queries().CreateUser(context.Background(), queries.CreateUserParams{
		Username:  entity.Username,
		Password:  entity.Password,
		Phone:     entity.PhoneNumber,
		FirstName: entity.Firstname,
		LastName:  entity.Lastname,
		Address:   entity.Address,
		IsActive:  entity.IsActive,
		BirthDate: birthDate,
		Role:      queries.Roles(entity.Role),
	})

	return uuid.String(), err
}
func (s *userService) GetUser(id string) (resp models.User, err error) {
	uid, _ := uuid.Parse(id)
	user, err := s.storage.Queries().GetUser(context.Background(), uid)

	resp = models.User{
		Guid:        user.Guid.String(),
		Username:    user.Username,
		Password:    user.Password,
		PhoneNumber: user.Phone,
		Firstname:   user.FirstName,
		Lastname:    user.LastName,
		IsActive:    user.IsActive,
		BirthDate:   user.BirthDate.Format("2006-01-02"),
		Role:        string(user.Role),
		CreatedAt:   user.CreatedAt.String(),
		UpdatedAt:   user.UpdatedAt.String(),
	}
	return resp, err
}
func (s *userService) GetAllUsers(limit, offset int) (resp []models.User, err error) {
	users, err := s.storage.Queries().GetAllUsers(context.Background(), queries.GetAllUsersParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})

	for _, user := range users {
		entity := models.User{
			Guid:        user.Guid.String(),
			Username:    user.Username,
			Password:    user.Password,
			PhoneNumber: user.Phone,
			Firstname:   user.FirstName,
			Lastname:    user.LastName,
			IsActive:    user.IsActive,
			BirthDate:   user.BirthDate.Format("2006-01-02"),
			Role:        string(user.Role),
			CreatedAt:   user.CreatedAt.String(),
			UpdatedAt:   user.UpdatedAt.String(),
		}
		resp = append(resp, entity)
	}
	return resp, err
}
func (s *userService) DeleteUser(guid string) (err error) {
	uid, _ := uuid.Parse(guid)

	err = s.storage.Queries().DeleteUser(context.Background(), uid)
	return err
}
func (s *userService) UpdateUserPassword(entity models.UpdateUserPassword) (id string, err error) {
	// params := queries.UpdateUserPasswordParams{}
	// id, err = s.storage.Queries().UpdateUserPassword(entity)
	return id, err
}
