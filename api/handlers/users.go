package handlers

import (
	"fmt"

	"github.com/abbos-ron2/go_postgresql/api/models"
	"github.com/abbos-ron2/go_postgresql/service"
	"github.com/gin-gonic/gin"
)

type user struct {
	s service.ServiceManager
}

func NewUser(s service.ServiceManager) *user {
	return &user{s: s}
}

// @Summary      Create user
// @Description  Create user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        data   body    models.CreateUser  true  "User"
// @Success      200  {object}  models.User
// @Router       /user [POST]
func (h *user) Create(c *gin.Context) {
	var body models.CreateUser
	err := c.BindJSON(&body)
	if err != nil {
		fmt.Printf("Error :%s", err)
	}

	resp, err := h.s.UserService().CreateUser(body)
	if err != nil {
		c.JSON(400, fmt.Sprintf("Error: %s", err))
		return
	}
	c.JSON(200, resp)
}

// @Summary      Update user
// @Description  Update user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        data   body    models.UpdateUser  true  "entity"
// @Success      200  {object}  models.User
// @Router       /user [PUT]
func (h *user) Update(c *gin.Context) {
	var body models.UpdateUser
	err := c.BindJSON(&body)
	if err != nil {
		fmt.Printf("Error while binding json:%s", err)
	}

	resp, err := h.s.UserService().UpdateUser(body)
	if err != nil {
		c.JSON(400, fmt.Sprintf("Error: %s", err))
		return
	}

	c.JSON(200, resp)
}

// @Summary      Get user
// @Description  Get user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "id"
// @Success      200  {object}  models.User
// @Router       /user/{id} [GET]
func (h *user) Get(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.s.UserService().GetUser(id)

	if err != nil {
		c.JSON(400, fmt.Sprintf("Error: %s", err))
		return
	}

	c.JSON(200, resp)
}
