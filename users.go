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

	return nil
}
