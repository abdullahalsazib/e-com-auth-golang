package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/neerajbg/go-gin-auth/controller"
)

func SetupRoutes(r *gin.Engine) {

	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)
	// r.GET("/logout", controller.Logout)
	// r.GET("/refreshtoken", controller.RefreshToken)
}
