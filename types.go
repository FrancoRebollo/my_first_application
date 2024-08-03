package main

import "time"

type UsersLogStruct struct {
	userName     string `json:userName`
	userPassword string `json:userPassword`
}

type UsersSignUp struct {
	userName         string    `json:userName`
	userEmail        string    `json:userEmail`
	userPassword     string    `json:userPassword`
	userPersonalID   string    `json:userPersonalID`
	userBirthdayDate time.Time `json:userBirthdayDate`
}
