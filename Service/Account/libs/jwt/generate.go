package jwt

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	models "go-module/model"
	jwtConfig "go-module/config/jwt"
)

func GenToken(user models.User) (string, error) {
	var jwtKey = []byte(jwtConfig.SECRET_KEY)
	type Claims struct {
		Email string `json:"email"`
		Name string  `json:"name"`
		Role string	 `json:"role,omitempty"`
		jwt.StandardClaims
	}

	expirationTime := time.Now().
		Add(jwtConfig.EXPIRE_PERIOD * time.Second)
	claims := &Claims{
		Email: user.Email,
		Name: user.Name,
		Role: user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}