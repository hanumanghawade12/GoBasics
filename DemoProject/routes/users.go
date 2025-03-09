package routes

import (
	"demoProject/models"
	"demoProject/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the data"})
		return
	}
	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save the user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created", "user": user})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}
	token, err := utils.GenereteToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}
