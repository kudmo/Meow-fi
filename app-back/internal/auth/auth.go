package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type JWTClaims struct {
	UserId         int `json:"user-id"`
	TokenId        string
	RefreshTokenId string
	Type           string
	jwt.RegisteredClaims
}

// From context field "user" gets jwt and gets user id from it
func TokenGetUserId(c echo.Context) int {
	aa := c.Get("user")
	return aa.(*jwt.Token).Claims.(*JWTClaims).UserId
}
