package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
)

// HasRole is a middleware that validates if the token contains the expected role in its claims
func HasRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		claims := token.CustomClaims.(*CustomClaims)
		if !claims.hasScope(role) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": fmt.Sprintf("you must have %s role to access this endpoint", role),
			})
		} else {
			c.Next()
		}
	}
}

// HasScope checks whether our claims have a specific scope.
func (c CustomClaims) hasScope(expectedScope string) bool {
	result := strings.Split(c.Scope, " ")
	for i := range result {
		if result[i] == expectedScope {
			return true
		}
	}

	return false
}
