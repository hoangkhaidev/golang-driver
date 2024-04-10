package security

import (
	"my-driver/model"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const SECRET_KEY = "MYSECRETKEY";

func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
        Claims:     &model.JwtCustomClaims{},
        SigningKey: []byte(SECRET_KEY),
    }
    return middleware.JWTWithConfig(config);
}

func GenToken(user model.User) (string, error) {
	claims := &model.JwtCustomClaims{
		UserId: 		user.UserId,
		Role: 			user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}

	return result, nil
}