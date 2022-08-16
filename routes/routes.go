package routes

import (
	controller "test-auth/controllers"
	"test-auth/middlewares"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.Use(middlewares.Auth0Middleware())
	r.Use()
	r.GET("/test", controller.TestController)
	r.Run()
}
