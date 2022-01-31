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
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.basic  BasicAuth
func New(serviceManager service.ServiceManager) *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "*")

	router.Use(cors.New(config))

	handler := handlers.NewHandler(serviceManager)

	api := router.Group("/api")
	{
		api.POST("/user", handler.User.Create)
		api.PUT("/user", handler.User.Update)
		api.GET("/user/:id", handler.User.Get)
	}

	// swagger
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
