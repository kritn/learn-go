package entities

import (
	"github.com/golang-jwt/jwt/v4"
)

type AuthRepository interface {
	SignUsersAccessToken(req *UsersPassport) (string, error)
}

type AuthUsecase interface {
	Login(req *UsersCredentials) (*UsersLoginRes, error)
}

type UsersCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UsersPassport struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UsersClaims struct {
	Id       int    `json:"user_id"`
	Username string `json:"username"`
	Timeout  string `json:"timeout"`
	jwt.RegisteredClaims
}

type UsersLoginRes struct {
	AccessToken string `json:"access_token"`
}
