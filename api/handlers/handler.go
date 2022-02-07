package handlers

import (
	"github.com/abbos-ron2/go_postgresql/service"
)

type handler struct {
	services service.ServiceManager
}

func NewHandler(s service.ServiceManager) *handler {
	return &handler{services: s}
}
