package repository

import "cf-service/model"

type AuthRepository interface {
	FindUserByEmail(email string) (*model.User, error)
}
