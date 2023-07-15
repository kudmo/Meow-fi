package auth

import (
	"Meow-fi_app-auth/internal/config"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type JWTClaims struct {
	UserId         int `json:"user-id"`
	TokenId        string
	RefreshTokenId string
	Type           string
	jwt.RegisteredClaims
}
type RTClaims struct {
	UserId  int `json:"user-id"`
	TokenId string
	Type    string
	jwt.RegisteredClaims
}

type AuthController struct {
}

// From context field "user" gets jwt and gets user id from it
func TokenGetUserId(c echo.Context) int {
	aa := c.Get("user")
	return aa.(*jwt.Token).Claims.(*JWTClaims).UserId
}

// Returns encoded token, token id, error.
//
// The current duration of the token is 15 minutes
func (controller *AuthController) CalculateAccessToken(userId int, refreshId string) (string, string, error) {
	tokenId := uuid.New()
	claims := &JWTClaims{
		UserId:         userId,
		TokenId:        tokenId.String(),
		RefreshTokenId: refreshId,
		Type:           "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.SecretKeyJWT))
	return t, tokenId.String(), err
}

// Returns encoded token, token id, error.
//
// The current duration of the token is 72 hours
func (controller *AuthController) CalculateRefreshToken(userId int) (string, string, error) {
	tokenId := uuid.New()
	claims := &RTClaims{
		UserId:  userId,
		TokenId: tokenId.String(),
		Type:    "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.SecretKeyRT))
	return t, tokenId.String(), err
}

// Get JWT an RT from json
//
//	json : {
//		"refresh_token" : RT
//		"access_token" : JWT
//	}
//
// if the JWT has expired it is not an error
// if RT has expired returns an appropriate error
func (controller *AuthController) GetTokensFromContext(c echo.Context) (*JWTClaims, *RTClaims, error) {
	type tokenReqBody struct {
		RefreshToken string `json:"refresh_token"`
		AccessToken  string `json:"access_token"`
	}
	tokenReq := tokenReqBody{}
	c.Bind(&tokenReq)
	rt, err := jwt.ParseWithClaims(tokenReq.RefreshToken, &RTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKeyRT), nil
	})

	if err != nil {
		return nil, nil, err
	}
	rtoken := rt.Claims.(*RTClaims)

	at, err := jwt.ParseWithClaims(tokenReq.AccessToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKeyJWT), nil
	})

	if err != nil && !errors.Is(err, jwt.ErrTokenExpired) {
		return nil, nil, err
	}
	atoken := at.Claims.(*JWTClaims)

	return atoken, rtoken, nil
}

func GenerateRandomSalt() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, config.LenRandomSalt)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func HashPassword(password, randomSalt, localSalt string) string {
	sum := sha256.Sum256([]byte(password + randomSalt + localSalt))
	return hex.EncodeToString(sum[:])
}
