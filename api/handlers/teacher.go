package handlers

import (
	"fmt"

	"github.com/abbos-ron2/go_postgresql/api/models"
	"github.com/gin-gonic/gin"
)

// @Router /teacher [POST]
// @Tags teacher
// @Summary create teacher
// @Description create teacher
// @Accept json
// @Param teacher body models.Teacher true "teacher"
// @Produce json
// @Success 200 {object} models.SuccessModel{data=models.Teacher.ID} "Success"
// @Response 422 {object} models.ErrorModel{error=string} "Validation Error"
// @Response 400 {object} models.ErrorModel "Bad Request"
// @Failure 500 {object} models.ErrorModel "Server Error"
func (h *handler) CreateTeacher(c *gin.Context) {
	var body models.CreateTeacher
	err := c.BindJSON(&body)
	if err != nil {
		fmt.Printf("Error :%s", err)
	}

	resp, err := h.services.TeacherService().CreateTeacher(body)
	if err != nil {
		c.JSON(400, fmt.Sprintf("Error: %s", err))
		return
	}
	c.JSON(200, resp)
}

// @Router /teacher/{guid} [PUT]
// @Tags teacher
// @Summary update teacher
// @Description update teacher
// @Accept json
// @Param teacher body models.UpdateTeacher true "teacher"
// @Param guid path string true "guid"
// @Produce json
// @Success 200 {object} models.SuccessModel{data=models.Teacher} "Success"
// @Response 422 {object} models.ErrorModel{error=string} "Validation Error"
// @Response 400 {object} models.ErrorModel "Bad Request"
// @Failure 500 {object} models.ErrorModel "Server Error"
func (h *handler) UpdateTeacher(c *gin.Context) {
	guid := c.Param("guid")

	var body models.UpdateTeacher
	err := c.BindJSON(&body)
	if err != nil {
		fmt.Printf("Error while binding json:%s", err)
	}

	resp, err := h.services.TeacherService().UpdateTeacher(guid, body)
	if err != nil {
		c.JSON(400, fmt.Sprintf("Error: %s", err))
		return
	}

	c.JSON(200, resp)
}

// @Router /teacher/{guid} [GET]
// @Tags teacher
// @Summary get teacher
// @Description get teacher
// @Accept json
// @Param guid path string true "guid"
// @Produce json
// @Success 200 {object} models.SuccessModel{data=models.Teacher} "Success"
// @Response 422 {object} models.ErrorModel{error=string} "Validation Error"
// @Response 400 {object} models.ErrorModel "Bad Request"
// @Failure 500 {object} models.ErrorModel "Server Error"
func (h *handler) GetTeacher(c *gin.Context) {
	guid := c.Param("guid")

	resp, err := h.services.UserService().GetUser(guid)

	if err != nil {
		c.JSON(400, fmt.Sprintf("Error: %s", err))
		return
	}

	c.JSON(200, resp)
}

// @Router /teacher [GET]
// @Tags teacher
// @Summary get all teachers
// @Description get all teachers
// @Accept json
// @Produce json
// @Success 200 {object} models.SuccessModel{data=[]models.Teacher} "Success"
// @Response 422 {object} models.ErrorModel{error=string} "Validation Error"
// @Response 400 {object} models.ErrorModel "Bad Request"
// @Failure 500 {object} models.ErrorModel "Server Error"
func (h *handler) GetAllTeachers(c *gin.Context) {
	resp, err := h.services.TeacherService().GetAllTeachers()
	if err != nil {
		c.JSON(400, fmt.Sprintf("Error: %s", err))
		return
	}

	c.JSON(200, resp)
}

// @Router /teacher/{guid} [DELETE]
// @Tags teacher
// @Summary delete teacher
// @Description delete teacher
// @Accept json
// @Param guid path string true "guid"
// @Produce json
// @Success 200 {object} models.SuccessModel{data=string} "Success"
// @Response 422 {object} models.ErrorModel{error=string} "Validation Error"
// @Response 400 {object} models.ErrorModel "Bad Request"
// @Failure 500 {object} models.ErrorModel "Server Error"
func (h *handler) DeleteTeacher(c *gin.Context) {
	guid := c.Param("guid")

	err := h.services.TeacherService().DeleteTeacher(guid)
	if err != nil {
		c.JSON(400, fmt.Sprintf("Error: %s", err))
		return
	}

	c.JSON(200, "Success")
}
