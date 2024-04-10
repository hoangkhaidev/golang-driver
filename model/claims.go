package model

import "github.com/golang-jwt/jwt"


type JwtCustomClaims struct {
	UserId 		string `json:"user_id"`
	Role 		string `json:"role"`
	jwt.StandardClaims 
}