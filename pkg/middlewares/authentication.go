package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JwtAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		hmacSampleSecret := []byte(os.Getenv("JWT_SECRET_KEY"))
		header := c.Request.Header.Get("Authorization")
		tokenString := strings.Replace(header, "Bearer ", "", 1)
		if len(strings.TrimSpace(header)) > 0 {
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				return hmacSampleSecret, nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				c.Set("user_id", claims["user_id"])
				c.Set("username", claims["username"])
			} else {
				c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": "forbidden", "massage": err.Error()})
			}
			c.Next()
			return
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized", "massage": "not found header"})
	}
}
