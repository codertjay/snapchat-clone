package main

import (
	"log"
	"os"
	snapchat_clone "snapchat-clone/snapchat-clone/migrate"
	"snapchat-clone/snapchat-clone/route"
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
	routers := route.SnapChatCloneRoutes()
	// run on this port
	err := routers.Run(":" + port)
	if err != nil {
		log.Panicln("Error running server", err)
	}
}
