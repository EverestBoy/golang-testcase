package service

import "cf-service/model"

type AuthService interface {
	UserEmailLoginService(credential model.Credential) (*model.UserView, error)
}
