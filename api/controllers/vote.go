package controllers

import (
	"fmt"
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

	if vote.EndDate.Sub(vote.StartDate).Hours() < 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "End time is less than Start time",
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
	var voteDb models.Vote

	// check if vote exists throw an not found error if not
	db := db.GetDB()
	if err := db.Where("uuid = ?", uuid).First(&voteDb).Error; err != nil {
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

	// Check admin rights when editing title or description
	jwtClaims := jwt.ExtractClaims(c)
	authUserAccessLevel := jwtClaims["access_level"].(float64)

	if authUserAccessLevel != 1 {
		if vote.Title != "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Sorry but you can't update title field on vote",
			})
			return

		}

		if vote.Description != "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Sorry but you can't update description (desc) field on vote",
			})
			return
		}
	}

	// check date
	if vote.EndDate.Sub(vote.StartDate).Hours() < 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "End time is less than Start time",
		})
		return
	}

	// update all fields if modified (specified in request)
	if vote.Title != "" {
		voteDb.Title = vote.Title
	}

	if vote.Description != "" {
		voteDb.Description = vote.Description
	}

	if !vote.StartDate.IsZero() {
		voteDb.StartDate = vote.StartDate
	}

	if !vote.EndDate.IsZero() {
		voteDb.EndDate = vote.EndDate
	}

	// save auth user vote
	var user models.User
	uuid = jwtClaims["uuid"].(string)

	if err := db.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Can't find auth user",
		})
		return
	}

	db.Model(&voteDb).Association("UUIDVote").Append(user)

	// Save vote update if no errors occurs
	if err := db.Save(&voteDb); err != nil {
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

	// Send modified data as result
	c.JSON(http.StatusAccepted, gin.H{
		"uuid":       voteDb.UUID,
		"title":      voteDb.Title,
		"desc":       voteDb.Description,
		"uuid_votes": nil,
		"start_date": fmt.Sprintf("%s", voteDb.StartDate.Format("02-01-2006")),
		"end_date":   fmt.Sprintf("%s", voteDb.EndDate.Format("02-01-2006")),
	})
}
