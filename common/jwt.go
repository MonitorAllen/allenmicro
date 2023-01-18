package common

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var Key = "jfslak34r4lfdkjg0394699d9s0"

type MyClaims struct {
	Username string `json:"username"`
	Uid      uint   `json:"id"`
	jwt.StandardClaims
}

func GetToken(username string, uid uint) string {
	claims := MyClaims{
		username, uid, jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(Key))
	if err != nil {
		panic(err)
	}

	return tokenString
}

func ParseToken(tokenString string) (*jwt.Token, MyClaims, error) {
	claims := MyClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(Key), nil
	})

	return token, claims, err
}
