package utils

import (
	"strconv"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type Token struct{
	SecretKey 	[]byte
}

func(t *Token) GenerateTokens(userId int64, email string) (string, error) {
	// role might be like user and admin


	// define the claims
	claims := jwt.MapClaims{
		"sub":  strconv.FormatInt(userId, 10),
		"email": email,
		"role": "user",
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // expires in 10 mins
		"iat":  time.Now().Unix(),
	}

	// this signing method means like same secret key performs finding and generating the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// create the string with the secret key
	accessToken, err := token.SignedString(t.SecretKey)

	if err != nil {
		return "", err
	}
	return accessToken, nil
}