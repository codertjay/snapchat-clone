package main

import (
	"log"
	"os"
	snapchat_clone "snapchat-clone/snapchat-clone"
)

func main() {
	/* get the def*/
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	// Migrate the database
	snapchat_clone.Migrate()

	// access the routes
	routers := snapchat_clone.SnapChatCloneRoutes()

	err := routers.Run(":" + port)
	if err != nil {
		log.Panicln("Error running server", err)
	}
}
