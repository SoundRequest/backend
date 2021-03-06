package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/SoundRequest/backend/helper"
	"github.com/SoundRequest/backend/structure"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CheckAuth handles authentication information
func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parsing Token From Header
		clientToken := c.GetHeader("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization Token is required"})
			c.Abort()
			return
		}

		extractedToken := strings.Split(clientToken, "Bearer ")

		// Verify if the format of the token is correct
		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect Format of Authorization Token "})
			c.Abort()
			return
		}

		// Parsing JWT To struct
		claims := &structure.Claims{}
		_, errParseWithClaims := jwt.ParseWithClaims(clientToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(helper.Config().JwtSecret), nil
		})

		// Check Correct OR Has Error
		if errParseWithClaims != nil {
			log.Println(errParseWithClaims)
			if errParseWithClaims.Error() == jwt.ErrSignatureInvalid.Error() {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Incorrect Format of Authorization Token or Failed to authorize token."})
			} else {
				log.Println(errParseWithClaims)
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error. This is not invalid token error. So this is not fault. I think there was an Oopsie in Internal Server."})
			}
			c.Abort()
			return
		}

		c.Set("UserId", claims.ID)
		c.Next()
	}
}
