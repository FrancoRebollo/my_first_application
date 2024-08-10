package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func NewJWTRequest(userEmail string, minutesDuration int) (*JWTRequest, error) {
	a := new(JWTRequest)

	if userEmail == "" || minutesDuration == 0 {
		return nil, fmt.Errorf("token creation failed")
	}

	a.UserEmail = userEmail
	a.MinutesDuration = minutesDuration

	return a, nil
}

func (s *APIServer) createUser(w http.ResponseWriter, r *http.Request) error {

	userSingUp := new(UserSignUp)

	if err := json.NewDecoder(r.Body).Decode(userSingUp); err != nil {
		return err
	}
	if err := userSingUpValidation(*userSingUp); err != nil {
		return err
	}
	if err := s.store.UserSignUp(*userSingUp); err != nil {
		return err
	}

	return nil
}

func (s *APIServer) logUser(w http.ResponseWriter, r *http.Request) (*UserLoginRes, error) {

	userLogin := new(UserLogin)

	if err := json.NewDecoder(r.Body).Decode(userLogin); err != nil {
		return nil, err
	}
	if err := userLoginValidation(*userLogin); err != nil {
		return nil, err
	}
	if err := s.store.UserLogin(*userLogin); err != nil {
		return nil, err
	}
	fmt.Println("Debugging")

	accesTokenReq, err := NewJWTRequest(userLogin.UserEmail, 15)

	if err != nil {
		return nil, err
	}

	accesToken, err := CreateJWT(accesTokenReq)
	if err != nil {
		return nil, fmt.Errorf("error creating JWT")
	}
	userResponse, err := s.store.UserGetByEmail(userLogin.UserEmail)
	if err != nil {
		return nil, err
	}

	userLoginRes := &UserLoginRes{
		AccessToken:          accesToken,
		AccessTokenExpiresAt: accesTokenReq.ExpiresAt,
		User: User{
			UserID:    userResponse.UserID,
			UserName:  userResponse.UserName,
			UserEmail: userResponse.UserEmail,
		},
	}

	return userLoginRes, nil
}
