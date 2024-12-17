package token

import (
	"github.com/golang-jwt/jwt"
	"time"
)

var secret = []byte("hello, this is my secret")
var JWTExpireDuration = time.Minute * 30
var JWTIssure = []byte("changtian")

type MyClaims struct {
	UserId int64 `json:"userId"`
	jwt.StandardClaims
}

func GenerateJWT(userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(JWTExpireDuration).Unix(),
			Issuer:    string(JWTIssure),
		},
	})
	return token.SignedString(secret)
}

func ParseJWT(tokenString string) (*MyClaims, error) {
	myClaims := new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, myClaims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
