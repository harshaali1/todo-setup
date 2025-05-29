package handlers

import (
	"tms-backend/src/database"
	"tms-backend/src/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error hashing password"})
		return
	}
	user.Password = string(hashedPassword)

	// Create user
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(201, gin.H{"message": "User created successfully"})
}

/*
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Email already exists"})
		return
	}

	c.JSON(201, gin.H{"id": user.ID, "email": user.Email})
}

func LoginHandler(c *gin.Context) (interface{}, error) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		return nil, jwt.ErrMissingLoginValues
	}

	var user models.User
	database.DB.Where("email = ?", loginData.Email).First(&user)

	if user.ID == 0 || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)) != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	return &user, nil
}*/
