package routes

import (
	"test-auth/constants"
	controller "test-auth/controllers"
	"test-auth/middlewares"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.Use(middlewares.Auth())
	r.GET("/test", middlewares.HasRole(constants.AdminRole), controller.TestController)
	r.Run()
}
