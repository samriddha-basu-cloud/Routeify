package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samriddha-basu-cloud/Routeify/models"
	"github.com/samriddha-basu-cloud/Routeify/repository"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate user input (e.g., check for required fields, format validation, etc.)

	// Save user to MongoDB
	err := repository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var credentials models.Credentials
	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Authenticate user (e.g., check credentials against database)

	// Generate JWT token
	token, err := generateToken(credentials.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
