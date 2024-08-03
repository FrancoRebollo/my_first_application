package main

import "time"

type UserLogin struct {
	userName     string `json:userName`
	userPassword string `json:userPassword`
}

type UserSignUp struct {
	userName         string    `json:userName`
	userEmail        string    `json:userEmail`
	userPassword     string    `json:userPassword`
	userPersonalID   string    `json:userPersonalID`
	userBirthdayDate time.Time `json:userBirthdayDate`
}

type User struct {
	userID    int    `json:userID`
	userName  string `json:userName`
	userEmail string `json:userEmail`
}
