package router

import (
	"Paste-Echo/app/controller"
	_ "Paste-Echo/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func Run(e *echo.Echo){



	e.GET("/swagger/*", echoSwagger.WrapHandler)

	api:=e.Group("/api")

	v1:=api.Group("/v1")

	v1.GET("/",controller.Index)

	v1.POST("/login",controller.Login)

	v1.GET("/user/:userid",controller.GetUserByUserId)

	v1.POST("/createUser",controller.CreateUser)

	v1.POST("/paste/expire",controller.CreateExpirePaste)

	v1.GET("/paste/expire/:key",controller.GetExpirePasteByKey)

	v1.POST("/paste",controller.CreateForeverPaste)

	v1.GET("/paste/:key",controller.GetLongTimePasteByKey)

}
