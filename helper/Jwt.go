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
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(Config().JwtSecret))

	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("token signed Error")
	}
	return tokenString, nil
}
