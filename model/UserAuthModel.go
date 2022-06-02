package model

import "time"

/*
@author galab pokharel
*/

type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	Status    int       `json:"status"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Token struct {
	Token        string `json:"jwtToken"`
	RefreshToken string `json:"refreshToken"`
}

type UserView struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	UserToken Token  `json:"token"`
}

type Credential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
