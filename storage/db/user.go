package db

import (
	"github.com/abbos-ron2/go_postgresql/api/models"
	"github.com/abbos-ron2/go_postgresql/storage/repo"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) repo.UserRepoI {
	return &userRepo{db: db}
}

func (i userRepo) Create(entity models.CreateUser) (id string, err error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return id, err
	}

	err = i.db.Get(&id, `insert into users(id,first_name, last_name, email, age) values($1,$2,$3,$4,$5) returning id;`,
		uuid.String(), entity.Firstname, entity.Lastname, entity.Email, entity.Age)
	return
}

func (i userRepo) Update(entity models.UpdateUser) (id string, err error) {
	err = i.db.Get(&id, `update users set first_name = $1, last_name = $2, email = $3, age = $4 where id::text = $5 returning id;`,
		entity.Firstname, entity.Lastname, entity.Email, entity.Age, entity.ID)
	return
}

func (i userRepo) Get(id string) (resp models.User, err error) {
	err = i.db.Get(&resp, `select * from users where id = $1;`, id)
	return
}
