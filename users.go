package main

import (
	"encoding/json"
	"net/http"
)

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

func (s *APIServer) logUser(w http.ResponseWriter, r *http.Request) error {

	userLogin := new(UserLogin)

	if err := json.NewDecoder(r.Body).Decode(userLogin); err != nil {
		return err
	}
	if err := userLoginValidation(*userLogin); err != nil {
		return err
	}
	if err := s.store.UserLogin(*userLogin); err != nil {
		return err
	}

	return nil
}
