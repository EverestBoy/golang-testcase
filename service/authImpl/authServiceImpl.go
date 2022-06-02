package authImpl

import (
	"cf-service/model"
	"cf-service/repository"
	"cf-service/service"
	"database/sql"
	"errors"
	"fmt"
)

/*
@author galab pokharel
*/

type authService struct {
	authRepo repository.AuthRepository
}

func (a authService) UserEmailLoginService(credential model.Credential) (*model.UserView, error) {
	fmt.Println("Do some service things here ")
	foundUser, err := a.authRepo.FindUserByEmail(credential.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("no user found with the given email")
		} else {
			err = errors.New("authentication failed")
		}
		return nil, err
	}
	userView := model.UserView{
		Id:    foundUser.Id,
		Email: foundUser.Email,
	}
	return &userView, nil
}

func NewAuthService(authRepo repository.AuthRepository) service.AuthService {
	return &authService{authRepo: authRepo}
}
