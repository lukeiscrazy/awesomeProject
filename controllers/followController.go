package controllers

import (
	"net/http"

	"awesomeProject/database"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
)

func FollowUser(c *gin.Context) {
	var input struct {
		FolloweeID uint `json:"followee_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("userID").(uint)
	follow := models.Follow{FollowerID: userID, FolloweeID: input.FolloweeID}

	if err := database.DB.Create(&follow).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to follow user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Followed user successfully!"})
}

func UnfollowUser(c *gin.Context) {
	var input struct {
		FolloweeID uint `json:"followee_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("userID").(uint)
	if err := database.DB.Where("follower_id = ? AND followee_id = ?", userID, input.FolloweeID).Delete(&models.Follow{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unfollow user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Unfollowed user successfully!"})
}
