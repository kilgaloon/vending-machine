package auth

import "github.com/dgrijalva/jwt-go"

// Credentials struct represents request to signin
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claims struct represents request to signin
type Claims struct {
	ID       uint `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
