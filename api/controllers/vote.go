package controllers

import (
	"net/http"

	"github.com/kconde2/vote-app/api/db"
	"github.com/kconde2/vote-app/api/models"

	"github.com/gin-gonic/gin"
)

// GetVotes get all votes
func GetVotes(c *gin.Context) {

	var votes []models.Vote
	db := db.GetDB()
	db.Find(&votes)
	c.JSON(200, votes)
}

// CreateVote create new vote subject
func CreateVote(c *gin.Context) {
	var vote models.Vote
	var db = db.GetDB()

	if err := c.BindJSON(&vote); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.Create(&vote)
	c.JSON(http.StatusOK, &vote)
}

// UpdateVote update specific vote
func UpdateVote(c *gin.Context) {
	id := c.Param("id")
	var vote models.Vote

	db := db.GetDB()
	if err := db.Where("id = ?", id).First(&vote).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.BindJSON(&vote)
	db.Save(&vote)
	c.JSON(http.StatusOK, &vote)
}

// DeleteVote delete specific vote
func DeleteVote(c *gin.Context) {
	id := c.Param("id")
	var vote models.Vote
	db := db.GetDB()

	if err := db.Where("id = ?", id).First(&vote).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Delete(&vote)
}
