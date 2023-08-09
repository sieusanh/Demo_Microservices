package custom_type

import "github.com/dgrijalva/jwt-go"

type AuthenInfo struct {
	Email string `json:"email"`
	Name string  `json:"name"`
	Role string	 `json:"role,omitempty"`
}

type Claims struct {
	// Email string `json:"email"`
	// Name string  `json:"name"`
	// Role string	 `json:"role,omitempty"`
	AuthenInfo		// Embedding AuthenInfo type
	jwt.StandardClaims
}
