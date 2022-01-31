package models

import "time"

type User struct {
	ID        string    `json:"id" db:"id"`
	Firstname string    `json:"first_name" db:"first_name"`
	Lastname  string    `json:"last_name" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	Age       uint      `json:"age" db:"age"`
	Created   time.Time `json:"created_at" db:"created_at" default:""`
}

type CreateUser struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
	Age       uint   `json:"age"`
}

type UpdateUser struct {
	ID        string `json:"id" db:"id"`
	Firstname string `json:"first_name" db:"first_name"`
	Lastname  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Age       uint   `json:"age" db:"age"`
}
