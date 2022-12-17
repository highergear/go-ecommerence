package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/highergear/go-ecommerence/utils"
	"net/http"
)

func JwtAuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized access")
			c.Abort()
			return
		}
		c.Next()
	}
}
