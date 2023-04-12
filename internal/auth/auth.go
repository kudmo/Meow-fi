package auth

import (
	"Meow-fi/internal/config"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	Id int `json:"id"`
	jwt.RegisteredClaims
}

func TokenGetUserId(c echo.Context) int {
	return c.Get("user").(*jwt.Token).Claims.(*JwtCustomClaims).Id
}

func RandSeq() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, config.LenRandomSalt)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func HashPass(password, randomSalt, localSalt string) string {
	sum := sha256.Sum256([]byte(password + randomSalt + localSalt))
	return hex.EncodeToString(sum[:])
}
