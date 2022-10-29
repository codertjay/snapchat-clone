package routers

import (
	"github.com/gin-gonic/gin"
	"snapchat-clone/controllers"
	"snapchat-clone/snapchat-clone/middleware"
)

func ChatRouters(incomingRoutes *gin.Engine) {
	//	Logged in route
	incomingRoutes.POST("api/v1/chats/send_request/", middleware.RequireAuth(), controllers.SendFriendRequest())
	incomingRoutes.GET("api/v1/chats/received_requests/", middleware.RequireAuth(), controllers.ReceivedFriendRequests())
	incomingRoutes.POST("api/v1/chats/accept_requests/", middleware.RequireAuth(), controllers.AcceptFriendRequest())
}
