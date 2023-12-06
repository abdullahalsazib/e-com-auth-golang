package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/neerajbg/go-gin-auth/database"
	"github.com/neerajbg/go-gin-auth/helper"
	"github.com/neerajbg/go-gin-auth/model"
	"golang.org/x/crypto/bcrypt"
)

type formData struct {
	Email    string `json:email`
	Password string `json:password`
}

// Login handler
func Login(c *gin.Context) {

	returnObject := gin.H{
		"status": "",
		"msg":    "Something went wrong.",
	}

	// 1. Check user for the given credentials

	var formData formData

	if err := c.ShouldBind(&formData); err != nil {
		log.Println("Form binding error.")

		c.JSON(400, returnObject)
		return
	}

	var user model.User

	database.DBConn.First(&user, "email=?", formData.Email)

	if user.ID == 0 {
		returnObject["msg"] = "User not found."

		c.JSON(400, returnObject)
		return
	}

	// Validate password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formData.Password))

	if err != nil {
		log.Println("Invalid password.")

		returnObject["msg"] = "Password doesnt match"
		c.JSON(401, returnObject)
		return
	}

	// 2. Create token

	token, err := helper.GenerateToken(user)

	if err != nil {
		returnObject["msg"] = "Token creation error."
		c.JSON(401, returnObject)
		return
	}

	returnObject["token"] = token
	returnObject["user"] = user
	returnObject["status"] = "OK"
	returnObject["msg"] = "User authenticated"
	c.JSON(200, returnObject)

}

// Register handler
func Register(c *gin.Context) {
	returnObject := gin.H{
		"status": "OK",
		"msg":    "Register route",
	}

	// Collect form data
	var formData formData

	if err := c.ShouldBind(&formData); err != nil {
		log.Println("Error in json binding.")
		returnObject["msg"] = "Error occurs."
		c.JSON(400, returnObject)
		return
	}

	// Add formdata to model
	var user model.User

	user.Email = formData.Email
	user.Password = helper.HashPassword(formData.Password)

	result := database.DBConn.Create(&user)

	if result.Error != nil {
		log.Println(result.Error)
		returnObject["msg"] = "User already exists."
		c.JSON(400, returnObject)
		return
	}

	returnObject["msg"] = "User added successfully."
	c.JSON(201, returnObject)
}

func Logout() {}

func RefreshToken(c *gin.Context) {
	returnObject := gin.H{
		"status": "OK",
		"msg":    "Refresh Token  route",
	}

	email, exists := c.Get("email")

	if !exists {
		log.Println("Email key not found.")

		returnObject["msg"] = "Email not found."
		c.JSON(401, returnObject)
		return
	}

	var user model.User

	database.DBConn.First(&user, "email=?", email)

	if user.ID == 0 {
		returnObject["msg"] = "User not found."

		c.JSON(400, returnObject)
		return
	}

	token, err := helper.GenerateToken(user)

	if err != nil {
		returnObject["msg"] = "Token creation error."
		c.JSON(401, returnObject)
		return
	}

	returnObject["token"] = token
	returnObject["user"] = user

	c.JSON(200, returnObject)
}

// User profile route
func Profile(c *gin.Context) {

	returnObject := gin.H{
		"status": "OK",
		"msg":    "User profile route. This is protected route accessible to authenticated user only.",
	}

	c.JSON(200, returnObject)
}
