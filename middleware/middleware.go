package middleware

import (
	"net/http"
    "github.com/LoanMgtGo/helpers"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helpers.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "UnAuthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
