package routes

import (
	
	"github.com/gin-gonic/gin"
	controller "github.com/lucapierini/UserAuthentication/controllers"
)

func AuthRouters(incomingRoutes *gin.Engine) {
	auth := incomingRoutes.Group("/users")
	{
		auth.POST("/singup", controller.Signup())
		auth.POST("/login", controller.Login())
	}
}