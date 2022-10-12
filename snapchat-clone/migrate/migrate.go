package migrate

import (
	"log"
	"snapchat-clone/models"
	"snapchat-clone/snapchat-clone/database"
)

/* The file is meant to migrate all models created*/

func Migrate() {
	// Migrate the schema
	db := database.DBConnection()
	/* User App migration */
	// Migrate the user models
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Panicln("Error migrating User models", err)
	}
	// Migrate the user profile
	err = db.AutoMigrate(&models.Profile{})
	if err != nil {
		log.Panicln("Error migrating User profile", err)
	}
	/* Chat app migration */

	// Migrate the Messages
	err = db.AutoMigrate(&models.Message{})
	if err != nil {
		log.Panicln("Error migrating Messages", err)
	}
	// Migrate the Conversation
	err = db.AutoMigrate(&models.Conversation{})
	if err != nil {
		log.Panicln("Error migrating Conversation", err)
	}
	// Migrate the FriendRequest
	err = db.AutoMigrate(&models.FriendRequest{})
	if err != nil {
		log.Panicln("Error migrating FriendRequest", err)
	}

	/*  Stories app migrations*/
	// Migrate the story app
	err = db.AutoMigrate(&models.Story{})
	if err != nil {
		log.Panicln("Error migrating Story", err)
	}

	/* Spotlight app */
	err = db.AutoMigrate(&models.Spotlight{})
	if err != nil {
		log.Panicln("Error migrating Spotlight", err)
	}
}
