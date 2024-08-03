package main

import "fmt"

func userSingUpValidation(typeValidate UserSignUp) error {

	if typeValidate.userName == "" {
		return fmt.Errorf("you have to provide an username")
	}
	if typeValidate.userEmail == "" {
		return fmt.Errorf("you have to provide an email")
	}
	if typeValidate.userPassword == "" {
		return fmt.Errorf("you have to provide a password")
	}
	if typeValidate.userPersonalID == "" {
		return fmt.Errorf("you have to provide your personal ID")
	}

	return nil
}
