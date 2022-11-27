package utils

import (
	"WebProject/ginEssential/mapping"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId string
	jwt.StandardClaims
}

// 发放token
func ReleaseToken(user mapping.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "oceanlearn.tech",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Clasims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, Clasims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, Clasims, err
}
