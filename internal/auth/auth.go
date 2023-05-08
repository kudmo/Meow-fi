package auth

import (
	"Meow-fi/internal/config"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"

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

func CalculateToken(userId int) (string, error) {
	claims := &JwtCustomClaims{
		Id: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.SecretKeyJwt))
	return t, err
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
