package users

import (
	"star/utility/jwt"
)

type Users struct {
	tokenManager *jwt.TokenManager
}

func New() *Users {
	return &Users{
		tokenManager: jwt.NewTokenManager(),
	}
}
func (u *Users) GetTokenManager() *jwt.TokenManager {
	return u.tokenManager
}
