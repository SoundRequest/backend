package helper

import (
	"fmt"
	"time"

	"github.com/SoundRequest/backend/structure"
	"github.com/dgrijalva/jwt-go"
)

// GetJwtToken with id
func GetJwtToken(id int) (string, error) {
	expirationTime := time.Now().Add(time.Hour)

	claims := &structure.Claims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, errStringToSignedString := token.SignedString([]byte(Config().JwtSecret))

	if errStringToSignedString != nil {
		fmt.Println(errStringToSignedString)
		return "", fmt.Errorf("token signed Error")
	}
	return tokenString, nil
}
