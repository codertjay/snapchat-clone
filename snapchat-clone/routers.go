package snapchat_clone

import (
	"github.com/gin-gonic/gin"
	"log"
)

func SnapChatCloneRoutes() *gin.Engine {
	router := gin.Default()
	err := router.SetTrustedProxies([]string{"0.0.0.0"})
	if err != nil {
		log.Panicln("Error setting trusted proxies", err)
	}
	/* Log */
	router.Use(gin.Logger())
	return router
}
