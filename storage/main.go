package storage

import (
	mongoDB "github.com/abbos-ron2/go_postgresql/storage/db"
	"github.com/abbos-ron2/go_postgresql/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	UserRepo() repo.UserRepoI
}

type storage struct {
	db       *sqlx.DB
	userRepo repo.UserRepoI
}

func NewStorage(db *sqlx.DB) StorageI {
	return storage{
		db:       db,
		userRepo: mongoDB.NewUserRepo(db),
	}
}

func (s storage) UserRepo() repo.UserRepoI {
	return s.userRepo
}
