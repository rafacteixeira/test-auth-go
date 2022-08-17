package middlewares

import (
	"net/http"
	"strings"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
)

func IsAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		claims := token.CustomClaims.(*CustomClaims)
		if !claims.HasScope("admin") {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "you must have admin role to access this endpoint",
			})
		} else {
			c.Next()
		}
	}
}

// HasScope checks whether our claims have a specific scope.
func (c CustomClaims) HasScope(expectedScope string) bool {
	result := strings.Split(c.Scope, " ")
	for i := range result {
		if result[i] == expectedScope {
			return true
		}
	}

	return false
}
