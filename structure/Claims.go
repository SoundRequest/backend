package structure

import "github.com/dgrijalva/jwt-go"

// Claims for jwt
type Claims struct {
	Id int
	jwt.StandardClaims
}
