package handlers

import "github.com/abbos-ron2/go_postgresql/service"

type handler struct {
	User *user
}

func NewHandler(s service.ServiceManager) *handler {
	return &handler{
		User: NewUser(s),
	}
}
