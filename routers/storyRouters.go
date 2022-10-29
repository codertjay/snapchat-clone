package routers

import (
	"github.com/gin-gonic/gin"
	"snapchat-clone/controllers"
	"snapchat-clone/snapchat-clone/middleware"
)

func StoryRouters(incomingRoutes *gin.Engine) {
	//	Logged in route
	// this returns all the stories  from the user friends
	incomingRoutes.GET("api/v1/stories/", middleware.RequireAuth(), controllers.UserStoryList())
}
