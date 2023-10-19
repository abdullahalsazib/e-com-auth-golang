package controller

import "github.com/gin-gonic/gin"

// Login handler
func Login(c *gin.Context) {

	returnObject := gin.H{
		"status": "OK",
		"msg":    "Login route",
	}
	c.JSON(200, returnObject)

}

// Register handler
func Register(c *gin.Context) {
	returnObject := gin.H{
		"status": "OK",
		"msg":    "Register route",
	}
	c.JSON(200, returnObject)
}

func Logout() {}

func RefreshToken() {}
