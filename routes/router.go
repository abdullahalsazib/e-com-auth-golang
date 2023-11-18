package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/neerajbg/go-gin-auth/controller"
	"github.com/neerajbg/go-gin-auth/middleware"
)

func SetupRoutes(r *gin.Engine) {

	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)
	// r.GET("/logout", controller.Logout)

	private := r.Group("/private")

	private.Use(middleware.Authenticate)

	private.GET("/refreshtoken", controller.RefreshToken)

	// r.GET("/refreshtoken", controller.RefreshToken)
}
