package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"snapchat-clone/models"
	"snapchat-clone/snapchat-clone/database"
	"strings"
	"time"
)

var SECRET_KEY = os.Getenv("SECRET_KEY")

/* This is used to get the tag of a field*/
type tagOptions string

func ParseTag(tag string) (string, tagOptions) {
	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx], tagOptions(tag[idx+1:])
	}
	return tag, tagOptions("")
}

func HashPassword(password string) string {
	userHashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panicln(err)
	}
	return string(userHashPassword)

}

func VerifyPassword(providedHashedPassword string, userPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(
		[]byte(providedHashedPassword),
		[]byte(userPassword),
	)
	check := true
	if err != nil {
		check = false
		return check, err
	}
	return check, nil
}

func CreateUser(user *models.User) (models.User, error) {
	/* To create a user we check if the email exist and also the phone */
	log.Println("Creating user", user)
	// hash the password
	password := HashPassword(*user.Password)
	user.Password = &password
	user.ID = uuid.New()
	// Generate the token
	token, refreshToken, err := GenerateAllToken(*user.Name, *user.Email, *user.Password, *&user.ID)
	if err != nil {
		return *user, err
	}
	user.AccessToken = &token
	user.RefreshToken = &refreshToken
	db := database.DBConnection()
	defer database.CloseDB()
	// create the user
	db.Create(&user)
	// Create profile for the user
	profile := &models.Profile{}
	profile.ID = uuid.New()
	profile.UserID = user.ID
	db.Create(&profile)
	log.Println(profile.ID)
	return *user, nil
}

// SignedInDetails /* This is meant to create access token and also the refresh token */
type SignedInDetails struct {
	Name  string
	Email string
	Phone string
	ID    uuid.UUID
	jwt.StandardClaims
}

// GenerateAllToken /* this is used to generate the access token and also the refresh token*/
func GenerateAllToken(Name string, Email string, Phone string, ID uuid.UUID) (signedToken string,
	refreshToken string, error error) {
	// Create the access claims
	claims := &SignedInDetails{
		Name:  Name,
		Email: Email,
		Phone: Phone,
		ID:    ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	// Create the refresh claims
	refreshClaims := &SignedInDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	/* Create the signed token with the claims provided for it and also the secret key*/
	SignedToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panicln("Error creating token", err)
		return "", "", err
	}
	/* Create the refresh token with the refresh claims provided*/
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panicln("Error creating refresh token", err)
		return "", "", err
	}

	return SignedToken, refreshToken, nil
}
