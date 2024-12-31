package controllers

import (
	"net/http"

	"awesomeProject/database"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
)

func GetFollowingList(c *gin.Context) {
	userID := c.MustGet("userID").(uint) // 从中间件获取用户ID

	var following []models.User
	if err := database.DB.Raw(`
		SELECT u.* FROM users u
		INNER JOIN follows f ON u.id = f.followee_id
		WHERE f.follower_id = ?
	`, userID).Scan(&following).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve following list"})
		return
	}

	c.JSON(http.StatusOK, following)
}
