package repo

import (
	"github.com/abbos-ron2/go_postgresql/api/models"
)

type UserRepoI interface {
	Create(entity models.CreateUser) (id string, err error)
	Update(entity models.UpdateUser) (id string, err error)
	Get(id string) (resp models.User, err error)
}
