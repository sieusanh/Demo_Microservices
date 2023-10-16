package jwt

import (
	"github.com/dgrijalva/jwt-go"
	jwtConfig "go-module/config/jwt"
	"go-module/libs/custom_type"
)

func Verify(tokenString string) (*jwt.Token, error) {
	jwtKey := []byte(jwtConfig.SECRET_KEY)
	
	tk := &custom_type.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, tk, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, err
}

