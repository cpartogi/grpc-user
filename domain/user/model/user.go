package model

import "time"

type Users struct {
	Id           string
	FullName     string
	Email        string
	PhoneNumber  string
	UserPassword string
}

type UserToken struct {
	Id                    string
	Token                 string
	TokenExpiredAt        time.Time
	RefreshToken          string
	RefreshTokenExpiredAt time.Time
}

type UserLogs struct {
	Id           string
	UserId       string
	IsSuccess    bool
	LoginMessage string
}
