package services

import (
	"log"
	"os"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
)

// jwt service
type JWTService interface {
	// GenerateToken(email string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtServices struct {
	issuer string
}

// auth-jwt
func JWTAuthService() JWTService {
	return &jwtServices{
		issuer: "Auth0",
	}
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, getJwksKeyFunc().Keyfunc)
}

func getJwksKeyFunc() *keyfunc.JWKS {
	jwksURL := os.Getenv("JWKS_URL")

	if jwksURL == "" {
		log.Fatalln("JWKS_URL environment variable must be populated.")
	}
	jwks, err := keyfunc.Get(jwksURL, keyfunc.Options{})
	if err != nil {
		log.Fatalf("Failed to get the JWKS from the given URL.\nError: %s", err)
	}
	return jwks
}
