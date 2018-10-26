package models

import jwt "github.com/dgrijalva/jwt-go"

// UserClaims store the JWT Claims
type UserClaims struct {
	UserName string
	Role     string
	jwt.StandardClaims
}
