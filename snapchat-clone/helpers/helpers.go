package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// This is used globally to prevent side import like ( importing from both file)

// SignedInDetails /* This is meant to create access token and also the refresh token */
type SignedInDetails struct {
	Name  string
	Email string
	Phone string
	ID    uuid.UUID
	jwt.StandardClaims
}
