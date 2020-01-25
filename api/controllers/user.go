package controllers

import (
	"net/http"

	"github.com/kconde2/vote-app/api/db"
	"github.com/kconde2/vote-app/api/models"

	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/gin-gonic/gin"
)

// GetUsers retrieve all users registered
func GetUsers(c *gin.Context) {
	var users []models.User
	db := db.GetDB()
	db.Find(&users)
	c.JSON(200, users)
}

// CreateUser create new user and save it into database
func CreateUser(c *gin.Context) {
	var user models.User
	var db = db.GetDB()

	// Check admin right creation
	// Retrive user information
	jwtClaims := jwt.ExtractClaims(c)
	authUserAccessLevel := jwtClaims["access_level"].(float64)

	if authUserAccessLevel != 1 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Sorry 🤔 but only admins can create user",
		})
		return
	}

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create user if no errors thrown by Validate method in user model
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

// UpdateUser update user information according to business logic
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
		authUserUUID := jwtClaims["uuid"].(string)
		if authUserAccessLevel != 1 {
			if authUserUUID != uuid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "Sorry but you can't Update, ONLY admin user can",
				})
				return
			}
		}

		var newUser models.User
		c.Bind(&newUser)

		if newUser.FirstName != "" {
			user.FirstName = newUser.FirstName
		}

		if newUser.LastName != "" {
			user.FirstName = newUser.LastName
		}

		if newUser.Email != "" {
			user.Email = newUser.Email
		}

		if newUser.Email != "" {
			user.Email = newUser.Email
		}

		// Update multiple attributes with `struct`, will only update those changed

		if err := db.Save(&user); err != nil {
			// convert array of errors to JSON
			errs := err.GetErrors()

			if len(errs) > 0 {
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
		}
		// Display modified data in JSON message "success"
		c.JSON(http.StatusOK, &user)

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
		authUserUUID := jwtClaims["uuid"].(string)
		if authUserAccessLevel != 1 {
			if authUserUUID != uuid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "Sorry but you can't delete user, ONLY admins can",
				})
				return
			}
		}
		// DELETE FROM users WHERE uuid= user.uuid
		// exemple : UPDATE users SET deleted_at=date.now WHERE uuid = user.uuid;
		db.Where("uuid = ?", uuid).Delete(&user)

		// Display JSON result
		// c.JSON(200, gin.H{"success": "User #" + uuid + " deleted"})
		c.JSON(200, gin.H{"success": "User successfully deleted"})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "User not found"})
	}

}
