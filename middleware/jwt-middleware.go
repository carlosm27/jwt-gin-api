package middleware

import (
	"fmt"
	"net/http"

	"github.com/carlosm27/jwt-gin-api/utils"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		err := utils.ValidateToken(c)
		fmt.Println(err)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Unauthorized": "Authentication required"})
			fmt.Println(err)
			c.Abort()
			return
		}
		c.Next()
	}
}
