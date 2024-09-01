package controllers

import (
	"net/http"
	"users/models"
	"users/services"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": gin.H{"code": 400, "message": "Invalid request"}})
		return
	}

	if err := services.RegisterUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": gin.H{"code": 500, "message": "Failed to register"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": gin.H{"code": 200, "message": "Success Register"}})
}

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": gin.H{"code": 400, "message": "Invalid request"}})
		return
	}

	authUser, err := services.LoginUser(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": gin.H{"code": 401, "message": "Invalid credentials"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"code":    200,
			"message": "Success Login",
		},
		"data": gin.H{
			"email":    authUser.Email,
			"fullName": authUser.FullName,
		},
	})
}
