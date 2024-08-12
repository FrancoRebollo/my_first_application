package main

import "time"

type UserLoginRes struct {
	SessionID             string    `json:"session_id"`
	AccessToken           string    `json:"access_token"`
	RefreshToken          string    `json:"refresh_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	User                  User      `json:"user"`
}

type UserLogin struct {
	UserEmail    string `json:"userEmail"`
	UserPassword string `json:"userPassword"`
}

type UserSignUp struct {
	UserFirstName    string `json:"userFirstName"`
	UserLastName     string `json:"userLastName"`
	UserName         string `json:"userName"`
	UserEmail        string `json:"userEmail"`
	UserPassword     string `json:"userPassword"`
	UserPersonalID   string `json:"userPersonalID"`
	UserBirthdayDate string `json:"userBirthdayDate"`
}

type User struct {
	UserID    int    `json:"userID"`
	UserName  string `json:"userName"`
	UserEmail string `json:"userEmail"`
}

type JWTRequest struct {
	UserEmail       string
	MinutesDuration int
	JWTToken        string
	ExpiresAt       time.Time
}

type NotificationContact struct {
	NotificationSender string
	NotificationMssg   string
}

type JWTCheckRefresh struct {
	UserEmail    string
	RefreshToken string
	IsValidYet   bool
}
