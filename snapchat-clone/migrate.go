package snapchat_clone

import (
	"log"
	"snapchat-clone/chats"
	"snapchat-clone/spotlights"
	"snapchat-clone/stories"
	"snapchat-clone/users"
)

/* The file is meant to migrate all models created*/

func Migrate() {
	// Migrate the schema
	db := DBConnection()
	/* User App migration */
	// Migrate the user models
	err := db.AutoMigrate(&users.User{})
	if err != nil {
		log.Panicln("Error migrating User models", err)
	}
	// Migrate the user profile
	err = db.AutoMigrate(&users.Profile{})
	if err != nil {
		log.Panicln("Error migrating User profile", err)
	}

	/* Chat app migration */

	// Migrate the Messages
	err = db.AutoMigrate(&chats.Message{})
	if err != nil {
		log.Panicln("Error migrating Messages", err)
	}
	// Migrate the Conversation
	err = db.AutoMigrate(&chats.Conversation{})
	if err != nil {
		log.Panicln("Error migrating Conversation", err)
	}
	// Migrate the FriendRequest
	err = db.AutoMigrate(&chats.FriendRequest{})
	if err != nil {
		log.Panicln("Error migrating FriendRequest", err)
	}

	/*  Stories app migrations*/
	// Migrate the story app
	err = db.AutoMigrate(&stories.Story{})
	if err != nil {
		log.Panicln("Error migrating Story", err)
	}

	/* Spotlight app */
	err = db.AutoMigrate(&spotlights.Spotlight{})
	if err != nil {
		log.Panicln("Error migrating Spotlight", err)
	}
}
