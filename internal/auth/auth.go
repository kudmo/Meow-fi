package auth

import (
	"Meow-fi/internal/config"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"math/rand"
	"net/http"
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

func TokenGetUserId(c echo.Context) int {
	aa := c.Get("user")
	return aa.(*jwt.Token).Claims.(*JWTClaims).UserId
}
func CalculateAccessToken(userId int, accessId string) (string, string, error) {
	tokenId := uuid.New()
	claims := &JWTClaims{
		UserId:         userId,
		TokenId:        tokenId.String(),
		RefreshTokenId: accessId,
		Type:           "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 1)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.SecretKeyJWT))
	return t, tokenId.String(), err
}
func CalculateRefreshToken(userId int) (string, string, error) {
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
func CalculateTokenPair(userId int) (map[string]string, error) {
	refreshToken, refreshId, err := CalculateRefreshToken(userId)
	if err != nil {
		return nil, err
	}
	accessToken, _, err := CalculateAccessToken(userId, refreshId)
	if err != nil {
		return nil, err
	}
	return map[string]string{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		},
		err
}
func RefreshJWT(c echo.Context) error {
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
		log.Println(err.Error())
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	rtoken := rt.Claims.(*RTClaims)

	at, err := jwt.ParseWithClaims(tokenReq.AccessToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKeyJWT), nil
	})
	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	atoken := at.Claims.(*JWTClaims)

	if rt.Valid && atoken.RefreshTokenId == rtoken.TokenId {
		newtoken, _, err := CalculateAccessToken(atoken.UserId, rtoken.TokenId)
		if err != nil {
			log.Println(err.Error())
			return c.String(http.StatusInternalServerError, "something goes wrong")
		}
		return c.JSON(http.StatusOK, map[string]string{
			"access_token": newtoken,
		})
	}
	return c.NoContent(http.StatusUnauthorized)
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
