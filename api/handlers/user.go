package handlers

import (
	"fmt"

	"github.com/abbos-ron2/go_postgresql/api/models"
	"github.com/gin-gonic/gin"
)

// @Summary      Create user
// @Description  Create user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        data   body    models.CreateUser  true  "User"
// @Success      200  {object}  models.User
// @Router       /user [POST]
func (h *handler) CreateUser(c *gin.Context) {
	var body models.CreateUser
	err := c.BindJSON(&body)
	if err != nil {
		fmt.Printf("Error :%s", err)
	}

	resp, err := h.services.UserService().CreateUser(body)
	if err != nil {
		c.JSON(400, fmt.Sprintf("Error: %s", err))
		return
	}
	c.JSON(200, resp)
}

// @Summary      Update user
// @Description  Update user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        data   body    models.UpdateUser  true  "entity"
// @Success      200  {object}  models.User
// @Router       /user [PUT]
// func (h *handler) UpdateUser(c *gin.Context) {
// 	var body models.UpdateUser
// 	err := c.BindJSON(&body)
// 	if err != nil {
// 		fmt.Printf("Error while binding json:%s", err)
// 	}
// 	resp, err := h.services.UserService().UpdateUser(body)
// 	if err != nil {
// 		c.JSON(400, fmt.Sprintf("Error: %s", err))
// 		return
// 	}
// 	c.JSON(200, resp)
// }

// @Summary      Get user
// @Description  Get user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "id"
// @Success      200  {object}  models.User
// @Router       /user/{id} [GET]
func (h *handler) GetUser(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.UserService().GetUser(id)

	if err != nil {
		c.JSON(400, fmt.Sprintf("Error: %s", err))
		return
	}

	c.JSON(200, resp)
}

// @Router /user [GET]
// @Tags user
// @Summary get all users
// @Description get all users
// @Accept json
// @Produce json
// @Success 200 {object} models.SuccessModel{data=[]models.User} "Success"
// @Response 422 {object} models.ErrorModel{error=string} "Validation Error"
// @Response 400 {object} models.ErrorModel "Bad Request"
// @Failure 500 {object} models.ErrorModel "Server Error"
func (h *handler) GetAllUsers(c *gin.Context) {
	resp, err := h.services.UserService().GetAllUsers(0, 0)

	if err != nil {
		c.JSON(400, fmt.Sprintf("Error: %s", err))
		return
	}

	c.JSON(200, resp)
}

// @Router /user [DELETE]
// @Tags teacher
// @Summary delete user
// @Description delete user
// @Accept json
// @Produce json
// @Response 422 {object} models.ErrorModel{error=string} "Validation Error"
// @Response 400 {object} models.ErrorModel "Bad Request"
// @Failure 500 {object} models.ErrorModel "Server Error"
// func (h *handler) DeleteUser(c *gin.Context) {
// 	id := c.Param("id")
// 	resp, err := h.services.UserService().DeleteUser(id)

// 	if err != nil {
// 		c.JSON(400, fmt.Sprintf("Error: %s", err))
// 		return
// 	}

// 	c.JSON(200, resp)
// }
