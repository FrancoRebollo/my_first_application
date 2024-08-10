package main

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func CreateJWT(req *JWTRequest) (string, error) {
	expiresAt := time.Now().Add(time.Minute * time.Duration(req.MinutesDuration)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": req.UserEmail,
			"exp":      expiresAt,
		})

	req.ExpiresAt = time.Unix(expiresAt, 0)

	tokenString, err := token.SignedString([]byte(os.Getenv("SEED_ENCRIPTATION")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWTAccess(tokenReceived string) error {
	_, err := jwt.Parse(tokenReceived, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SEED_ENCRIPTATION")), nil
	})

	if err != nil {
		return err
	}
	return nil

}

func RefreshJWTAccess(req JWTRequest) (string, error) {
	return "", nil
}
