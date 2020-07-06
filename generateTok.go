package main

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

//Создает токены
func generateTokenPair(guid string) (map[string]string, error) {

	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = 1
	claims["name"] = "User1"
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}
	refreshToken := jwt.New(jwt.SigningMethodHS512)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	mongoCreateUser(guid,rt)

	return map[string]string{
		"access_token":  t,
		"refresh_token": rt,
	}, nil
}
