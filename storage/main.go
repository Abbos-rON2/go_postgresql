package storage

import (
	"database/sql"

	postgres "github.com/abbos-ron2/go_postgresql/storage/db"
	"github.com/abbos-ron2/go_postgresql/storage/repo"
	queries "github.com/abbos-ron2/go_postgresql/storage/sqlc"
)

type StorageI interface {
	UserRepo() repo.UserRepoI
	Queries() *queries.Queries
}

type storage struct {
	db       *sql.DB
	userRepo repo.UserRepoI
	queries  *queries.Queries
}

func NewStorage(db *sql.DB) StorageI {
	queries := queries.New(db)
	return storage{
		db:       db,
		userRepo: postgres.NewUserRepo(db),
		queries:  queries,
	}
}

func (s storage) UserRepo() repo.UserRepoI {
	return s.userRepo
}

func (s storage) Queries() *queries.Queries {
	return s.queries
}
