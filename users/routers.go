package users

import "github.com/gin-gonic/gin"

func Routers(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("api/v1/auth/signup/", UserSignup())
	incomingRoutes.POST("api/v1/auth/login/", UserLogin())
}
