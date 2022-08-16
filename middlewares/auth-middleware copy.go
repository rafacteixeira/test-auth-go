package middlewares

import (
	"fmt"
	"net/http"
	service "test-auth/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthorizeJWT2() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := service.JWTAuthService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
			c.Next()
		} else {
			fmt.Println(err.Error())
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
