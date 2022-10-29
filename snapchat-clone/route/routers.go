package route

import (
	"github.com/gin-gonic/gin"
	"log"
	"snapchat-clone/routers"
)

func SnapChatCloneRoutes() *gin.Engine {
	router := gin.Default()
	err := router.SetTrustedProxies([]string{"0.0.0.0"})
	if err != nil {
		log.Panicln("Error setting trusted proxies", err)
	}
	/* Log */
	router.Use(gin.Logger())
	// Users routes
	routers.UserRouters(router)
	routers.ChatRouters(router)
	routers.StoryRouters(router)
	return router
}
