package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)


func validateJwt(token string, key []byte)(*jwt.Token, error){
	// parse(token, key and nil is of options)
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error){
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok{
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return key, nil
	})


	if err != nil{
		fmt.Println("[VALIDATION ERROR]: ", err)
		return nil, err
	}

	// if jwtToken is not valid
	if !jwtToken.Valid{
		return nil, fmt.Errorf("Invalid token.")
	}
	return jwtToken, nil
}

func ValidateToken(token string) (int,error){
	// validate the token obtained

	// get the secret key
	err := godotenv.Load()
	if err != nil{
		// show the error
		fmt.Println("[VALIDATING ERROR]: Error in loading the env file")
		return -1, err
	}

	// Now get the secret key
	secretKey := []byte(os.Getenv("SECRET_KEY"))

	// now validate the jwt token and get the claims
	jwtToken, err := validateJwt(token, secretKey)
	if err != nil{
		return -1, err
	}

	// get all of the claims
	claims := jwtToken.Claims.(jwt.MapClaims)
	sub, ok := claims["sub"].(string)
	if !ok {
    	return 0, errors.New("invalid subject")
	}

	userId, err := strconv.ParseInt(sub, 10, 64)
	expFloat, ok := claims["exp"].(float64)

	if !ok {
		fmt.Println("[UTILS: Token Validation] Not correct exp format")
		return -1, errors.New("Invalid exp format.")
	}

	if time.Now().Unix() > int64(expFloat) {
		fmt.Println("[UTILS: Token Validation] Token expired")
		return -1, errors.New("Token expired.")
	}

	return int(userId), nil
}