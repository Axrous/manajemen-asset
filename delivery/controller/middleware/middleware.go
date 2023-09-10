package middleware

import (
	"final-project-enigma-clean/helper"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization")

		//if tokenheader is empty
		if tokenHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"Error": "Unauthorized"})
			return
		}

		//replace bearer nya
		tokenHeader = strings.Replace(tokenHeader, "Bearer ", "", 1)

		//do parse jwt
		token, err := helper.ParseJWT(tokenHeader)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"Error": "Failed to parse jwt"})
			return
		}

		//claim token nya
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", claims["user_id"])
			return
		}
		//next
		c.Next()
	}
}
