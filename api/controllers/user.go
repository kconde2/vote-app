package controllers

import (
	"net/http"
	"log"
	"golang.org/x/crypto/bcrypt"

	"github.com/kconde2/vote-app/api/db"
	"github.com/kconde2/vote-app/api/models"

	"github.com/gin-gonic/gin"
)

// GetUsers g
func GetUsers(c *gin.Context) {

	var users []models.User
	db := db.GetDB()
	db.Find(&users)
	c.JSON(200, users)
}

// CreateUser c
func CreateUser(c *gin.Context) {

	var user models.User
	var db = db.GetDB()

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// Validate request fields
	if err := user.Valid(); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	// Hash password before storing into database
	pwd := user.Password

	if hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost); err != nil{
		log.Println(err)
		return
	} else{
		user.Password = string(hash)
	}

	err := db.Create(&user)
	if err != nil {
		log.Println("DB is nil", db)
		return
	}
	c.JSON(http.StatusOK, &user)
}

// UpdateUser u
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	db := db.GetDB()
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.BindJSON(&user)
	db.Save(&user)
	c.JSON(http.StatusOK, &user)
}

// DeleteUser delete specific user
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	db := db.GetDB()

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Delete(&user)
}
