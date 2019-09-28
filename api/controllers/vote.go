package controllers

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/kconde2/vote-app/api/db"
	"github.com/kconde2/vote-app/api/models"

	"github.com/gin-gonic/gin"
)

// GetVotes get all votes
func GetVotes(c *gin.Context) {
	var votes []models.Vote
	db := db.GetDB()
	db.Find(&votes)
	c.JSON(http.StatusOK, votes)
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

	// Retrive auth user information and check if it is authaurized to perform this action
	jwtClaims := jwt.ExtractClaims(c)
	authUserAccessLevel := jwtClaims["access_level"].(float64)
	if authUserAccessLevel != 1 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Sorry but you can't created vote",
		})
		return
	}

	if err := db.Create(&vote); err.Error != nil {

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
	c.JSON(http.StatusOK, &vote)
}

// RetrieveVote retrieve specific vote
func RetrieveVote(c *gin.Context) {
	uuid := c.Param("uuid")
	var vote models.Vote

	// check if vote exists throw an not found error if not
	db := db.GetDB()
	if err := db.Where("uuid = ?", uuid).First(&vote).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// Find vote and associated voter
	db.Model(&vote).Association("UUIDVote").Find(&vote.UUIDVote)

	// Get only all uuid from voter list
	voteUUIDs := []string{}
	for _, v := range vote.UUIDVote {
		voteUUIDs = append(voteUUIDs, v.UUID.String())
	}

	// return json data
	c.JSON(http.StatusOK, gin.H{
		"uuid":       vote.UUID,
		"title":      vote.Title,
		"desc":       vote.Description,
		"uuid_votes": voteUUIDs,
	})
}

// UpdateVote update specific vote
func UpdateVote(c *gin.Context) {
	uuid := c.Param("uuid")
	var vote models.Vote

	// check if vote exists throw an not found error if not
	db := db.GetDB()
	if err := db.Where("uuid = ?", uuid).First(&vote).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Vote not found",
		})
		return
	}

	if err := c.BindJSON(&vote); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Save vote update if no errors occurs
	if err := db.Save(&vote); err != nil {
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

	// Send message as result
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Vote updated successfully",
	})
}
