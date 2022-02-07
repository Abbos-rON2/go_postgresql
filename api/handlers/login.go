package handlers

import (
	"fmt"

	"github.com/abbos-ron2/go_postgresql/api/models"
	"github.com/gin-gonic/gin"
)

// @Router /login [POST]
// @Tags user
// @Summary login
// @Description login
// @Param data body models.LoginModel true "data"
// @Accept json
// @Produce json
// @Response 422 {object} models.ErrorModel{error=string} "Validation Error"
// @Response 400 {object} models.ErrorModel "Bad Request"
// @Failure 500 {object} models.ErrorModel "Server Error"
func (h *handler) Login(c *gin.Context) {
	var body models.LoginModel
	c.BindJSON(&body)
	token, err := h.services.UserService().Login(body)
	if err != nil {
		c.JSON(400, fmt.Sprintf("Error: %s", err))
		return
	}
	c.SetCookie("token", token, 60*60*24, "/", "localhost", true, true)

	c.JSON(200, token)
}
