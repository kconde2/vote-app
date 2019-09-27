package controllers

import (
	"net/http"

	"github.com/kconde2/vote-app/api/db"
	"github.com/kconde2/vote-app/api/models"

	jwt "github.com/appleboy/gin-jwt/v2"

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

	// Check admin right creation
	if user.IsAdmin() {
		// Retrive user information
		jwtClaims := jwt.ExtractClaims(c)
		authUserAccessLevel := jwtClaims["access_level"].(float64)
		if authUserAccessLevel != 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Sorry but you can't created admin user",
			})
			return
		}
	}

	// Create user if no errors throws by Validate method in user model
	if err := db.Create(&user); err.Error != nil {

		// convert array of errors to JSON
		errs := err.GetErrors()
		strErrors := make([]string, len(errs))
		for i, err := range errs {
			strErrors[i] = err.Error()
		}

		// return errors
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": strErrors,
		})

		return
	}

	c.JSON(http.StatusOK, &user)
}

// UpdateUser u
func UpdateUser(c *gin.Context) {
	uuid := c.Param("uuid")
	var user models.User

	db := db.GetDB()
	if err := db.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	db.Where("uuid = ?", uuid)

	if user.ID != 0 {

		jwtClaims := jwt.ExtractClaims(c)
		authUserAccessLevel := jwtClaims["access_level"].(float64)
		if authUserAccessLevel != 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Sorry but you can't Update, ONLY admin user can",
			})
			return
		}

		var newUser models.User
		c.Bind(&newUser)

		// Update multiple attributes with `struct`, will only update those changed
		result := db.Model(&user).Update(map[string]interface{}{"first_name": newUser.FirstName, "last_name": newUser.LastName, "email": newUser.Email, "pass": newUser.Password, "actived": false})
		// Display modified data in JSON message "success"
		c.JSON(200, gin.H{"success": result})

	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "User not found"})
	}

}

// DeleteUser delete specific user
func DeleteUser(c *gin.Context) {
	uuid := c.Params.ByName("uuid")
	var user models.User
	db := db.GetDB()
	if uuid != "" {

		jwtClaims := jwt.ExtractClaims(c)
		authUserAccessLevel := jwtClaims["access_level"].(float64)
		if authUserAccessLevel != 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Sorry but you can't Delete, ONLY admin user can",
			})
			return
		}
		// DELETE FROM users WHERE uuid= user.uuid
		// exemple : UPDATE users SET deleted_at=date.now WHERE uuid = user.uuid;
		db.Where("uuid = ?", uuid).Delete(&user)

		// Display JSON result
		c.JSON(200, gin.H{"success": "User #" + uuid + " deleted"})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "User not found"})
	}

}
