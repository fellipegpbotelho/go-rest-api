package middlewares

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	jwtManager "github.com/fellipegpbotelho/go-rest-api/jwt"
	gin "github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		const SCHEMA = "Bearer"
		authorizationHeader := c.GetHeader("Authorization")
		encodedToken := authorizationHeader[len(SCHEMA):]
		token, err := jwtManager.BuildJWTManager().ValidateToken(encodedToken)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
