package controllers

import (
	"net/http"

	config "playnex-api/configs"
	"playnex-api/dtos"
	"playnex-api/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register handles new user registration.
// POST /api/auth/register
func RegisterUser(c *gin.Context) {
	var req dtos.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if email already exists
	var existingUser models.Users
	if result := config.DB.Where("email = ?", req.Email).First(&existingUser); result.RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create user
	user := models.Users{
		FullName:  req.FullName,
		Gender:    req.Gender,
		AvatarUrl: req.AvatarUrl,
		Email:     req.Email,
		Password:  string(hashedPassword),
	}

	if result := config.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, dtos.RegisterSuccessful{
		FullName: user.FullName,
		Gender:   user.Gender,
		Email:    user.Email,
		Message:  "Registered successfully",
	})
}
