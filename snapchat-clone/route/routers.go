package route

import (
	"github.com/gin-gonic/gin"
	"log"
	"snapchat-clone/users"
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
	users.Routers(router)
	return router
}
