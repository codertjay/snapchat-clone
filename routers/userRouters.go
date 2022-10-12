package routers

import (
	"github.com/gin-gonic/gin"
	"snapchat-clone/controllers"
	"snapchat-clone/snapchat-clone/middleware"
)

func Routers(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("api/v1/auth/signup/", controllers.UserSignup())
	incomingRoutes.POST("api/v1/auth/login/", controllers.UserLogin())
	//	Logged in route
	incomingRoutes.POST("api/v1/auth/user_update/", middleware.RequireAuth(), controllers.UserUpdate())
}