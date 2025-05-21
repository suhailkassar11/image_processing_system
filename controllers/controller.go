package controllers

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/suhailkassar11/image_processing_system/initializers"
	"github.com/suhailkassar11/image_processing_system/models"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = os.Getenv("SECRET_KEY")

func CreateUser(c *gin.Context) {

	fmt.Println("create user")
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err, "message": "please pass the correct value"})
		return
	}
	var existingUser models.User
	err = initializers.DB.Where("email=?", &user.Email).First(&existingUser).Error

	if err == nil {
		c.JSON(400, gin.H{"message": "user already exist"})
		return
	}

	hassedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(400, gin.H{"error": err, "message": "error in encrypt the passwrod"})
	}
	user.Password = string(hassedPassword)

	err = initializers.DB.Create(&user).Error
	if err != nil {
		c.JSON(400, gin.H{"error": err, "message": "enable to insert user in database"})
		return
	}
	c.JSON(200, gin.H{"user": user, "message": "user is created successfully"})
}

func LoginUser(c *gin.Context) {

	var userLogin models.UserLogin

	err := c.BindJSON(&userLogin)

	fmt.Println(userLogin.Email)
	if err != nil {
		c.JSON(400, gin.H{"error": err, "message": "error is bind json"})
		return
	}

	var user models.User

	err = initializers.DB.Where("email=?", userLogin.Email).First(&user).Error

	if err != nil {
		c.JSON(400, gin.H{"error": err, "message": "user not found"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))
	if err != nil {
		c.JSON(400, gin.H{"error": err, "message": "password is invalid"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		c.JSON(400, gin.H{"error": err, "message": "token is invalid"})
		return
	}

	c.JSON(200, gin.H{"token": tokenString, "message": "user login successfully"})
}
