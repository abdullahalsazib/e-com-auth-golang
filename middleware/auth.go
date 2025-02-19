package middleware

import (
	"log"

	"github.com/abdullahalsazib/e-com-auth-golang/helper"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {

	token := c.GetHeader("token")

	if token == "" {
		c.JSON(401, gin.H{"error": "Token not present."})
		c.Abort()
		return
	}

	claims, msg := helper.ValidateToken(token)

	log.Println(claims)

	if msg != "" {
		c.JSON(401, gin.H{"error": msg})
		c.Abort()
		return
	}

	c.Set("email", claims.Email)

	c.Next()

}
