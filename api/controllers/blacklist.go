package controllers

import (
	"net/http"

	"github.com/kconde2/vote-app/api/db"
	"github.com/kconde2/vote-app/api/models"

	"github.com/gin-gonic/gin"
)

// AddToBlacklist add ip address to blacklist
func AddToBlacklist(c *gin.Context) {
	var blacklist models.Blacklist
	var db = db.GetDB()

	if err := c.BindJSON(&blacklist); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.Create(&blacklist)
	c.JSON(http.StatusOK, &blacklist)
}

// RemoveFromBlacklist remove ip address from blacklist
func RemoveFromBlacklist(c *gin.Context) {
	id := c.Param("id")
	var blacklist models.Blacklist
	db := db.GetDB()

	if err := db.Where("id = ?", id).First(&blacklist).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Delete(&blacklist)
}
