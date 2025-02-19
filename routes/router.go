package routes

import (
	"github.com/abdullahalsazib/e-com-auth-golang/controller"
	"github.com/abdullahalsazib/e-com-auth-golang/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/test", controller.Test)
	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)
	// r.GET("/logout", controller.Logout)

	private := r.Group("/private")

	private.Use(middleware.Authenticate)

	private.GET("/refreshtoken", controller.RefreshToken)
	private.GET("/profile", controller.Profile)

	// r.GET("/refreshtoken", controller.RefreshToken)
}
