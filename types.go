package main

type UserLogin struct {
	UserEmail    string `json:userEmail`
	UserPassword string `json:userPassword`
}

type UserSignUp struct {
	UserFirstName    string `json:userFirstName`
	UserLastName     string `json:userLastName`
	UserName         string `json:userName`
	UserEmail        string `json:userEmail`
	UserPassword     string `json:userPassword`
	UserPersonalID   string `json:userPersonalID`
	UserBirthdayDate string `json:userBirthdayDate`
}

type User struct {
	UserID    int    `json:userID`
	UserName  string `json:userName`
	UserEmail string `json:userEmail`
}
