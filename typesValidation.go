package main

import "fmt"

func userSingUpValidation(typeValidate UserSignUp) error {

	if typeValidate.UserName == "" {
		return fmt.Errorf("you have to provide an username")
	}
	if typeValidate.UserEmail == "" {
		return fmt.Errorf("you have to provide an email")
	}
	if typeValidate.UserPassword == "" {
		return fmt.Errorf("you have to provide a password")
	}
	if typeValidate.UserPersonalID == "" {
		return fmt.Errorf("you have to provide your personal ID")
	}

	return nil
}
