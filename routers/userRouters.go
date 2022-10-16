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
	incomingRoutes.PUT("api/v1/user/user_update/", middleware.RequireAuth(), controllers.UserUpdate())
	incomingRoutes.GET("api/v1/user/user_detail/", middleware.RequireAuth(), controllers.UserDetail())
	incomingRoutes.GET("api/v1/user/profile_detail/", middleware.RequireAuth(), controllers.ProfileDetail())
	incomingRoutes.PUT("api/v1/user/profile_update/", middleware.RequireAuth(), controllers.ProfileUpdate())
}
