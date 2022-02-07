package api

import (
	_ "github.com/abbos-ron2/go_postgresql/api/docs"
	handlers "github.com/abbos-ron2/go_postgresql/api/handlers"
	"github.com/abbos-ron2/go_postgresql/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Test Task
// @version         1.0
// @description     This is a sample server celler server.
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @host      localhost:8080
// @BasePath  /
func New(serviceManager service.ServiceManager) *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "*")

	router.Use(cors.New(config))

	handler := handlers.NewHandler(serviceManager)

	api := router.Group("/")
	{
		api.POST("/login", handler.Login)

		api.POST("/user", handler.CreateUser)
		api.GET("/user", handler.GetAllUsers)
		api.GET("/user/:id", handler.GetUser)
		// api.PUT("/user", handler.UpdateUser)
		// api.DELETE("/user/:id", handler.DeleteUser)

		api.POST("/teacher", handler.CreateTeacher)
		api.GET("/teacher", handler.GetAllTeachers)
		api.GET("/teacher/:id", handler.GetTeacher)
		api.PUT("/teacher", handler.UpdateTeacher)
		api.DELETE("/teacher/:id", handler.DeleteTeacher)

	}

	// swagger
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
