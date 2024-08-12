package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type JWTError struct {
	Error string
}

func CreateJWT(req *JWTRequest) (string, error) {
	expiresAt := time.Now().Add(time.Minute * time.Duration(req.MinutesDuration)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userEmail": req.UserEmail,
			"exp":       expiresAt,
		})

	req.ExpiresAt = time.Unix(expiresAt, 0)

	tokenString, err := token.SignedString([]byte(os.Getenv("SEED_ENCRIPTATION")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWTAccess(tokenReceived string) (string, error) {
	token, err := jwt.Parse(tokenReceived, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SEED_ENCRIPTATION")), nil
	})

	if err != nil {
		return "", err
	}

	claims, bool := token.Claims.(jwt.MapClaims)

	fmt.Println(claims)

	if !bool {
		return "", fmt.Errorf("invalid claims")
	}

	tokenMail, bool := claims["userEmail"].(string)
	fmt.Println(tokenMail)
	if !bool {
		return "", fmt.Errorf("can't read mail from claims")
	}

	return tokenMail, nil

}

func JWTValidationMiddleware(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("x-jwt-token")
		if tokenString == "" {
			err := fmt.Errorf("no token found")
			WriteJSON(w, http.StatusBadRequest, JWTError{Error: err.Error()})
			return
		}
		addRequest, err := ValidateJWTAccess(tokenString)

		if err != nil {
			WriteJSON(w, http.StatusForbidden, JWTError{Error: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), "commonIdentification", addRequest)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}

func RefreshJWTAccess(req JWTRequest) (string, error) {
	return "", nil
}

func (s *APIServer) JWTCheckRefreshToken(a *JWTCheckRefresh) (*JWTCheckRefresh, error) {
	return s.store.CheckJWTRefreshToken(a)
}
