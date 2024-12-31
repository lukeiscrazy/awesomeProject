package routes

import (
	"awesomeProject/controllers"
	"awesomeProject/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		userGroup := api.Group("/user", middleware.AuthMiddleware())
		{
			userGroup.POST("/follow", controllers.FollowUser)
			userGroup.DELETE("/unfollow", controllers.UnfollowUser)
			userGroup.GET("/following", controllers.GetFollowingList)
		}
	}
}
