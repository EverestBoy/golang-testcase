package authImpl

import (
	"cf-service/model"
	"cf-service/repository"
	mocks2 "cf-service/repository/mocks"
	"database/sql"
	"errors"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func getMockedMethod(expectedUser *model.UserView) mocks2.AuthRepository {
	// mocking the service
	authRepoMock := mocks2.AuthRepository{}
	authRepoMock.On("FindUserByEmail", mock.AnythingOfType("string")).
		Return(&model.User{Id: expectedUser.Id, Email: expectedUser.Email}, nil).
		Once()
	return authRepoMock
}
func getMockedRepoNoRowsError(sqlError bool) mocks2.AuthRepository {
	// mocking the service
	authRepoMock := mocks2.AuthRepository{}
	var err error
	if sqlError {
		err = sql.ErrNoRows
	} else {
		err = errors.New("some other error")
	}
	authRepoMock.On("FindUserByEmail", mock.AnythingOfType("string")).
		Return(nil, err).
		Once()
	return authRepoMock
}

func Test_authService_UserEmailLoginService(t *testing.T) {
	type args struct {
		credential model.Credential
	}

	tests := []struct {
		name      string
		args      args
		want      *model.UserView
		wantErr   bool
		sqlErr    bool
		fields    repository.AuthRepository
		errorType error
	}{
		{
			name: "No error test",
			args: args{credential: model.Credential{Email: "galabwot@gmail.com"}},
			want: &model.UserView{
				Email: "galabwot@gmail.com",
				Id:    1,
			},
			wantErr: false,
		},
		{
			name:      "Sql user not found test",
			args:      args{credential: model.Credential{Email: "galabwot@gmail.com"}},
			wantErr:   true,
			sqlErr:    true,
			errorType: errors.New("no user found with the given email"),
		},
		{
			name:      "Authentication failed test",
			args:      args{credential: model.Credential{Email: "galabwot@gmail.com"}},
			wantErr:   true,
			sqlErr:    false,
			errorType: errors.New("authentication failed"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var authRepo mocks2.AuthRepository
			if tt.wantErr {
				authRepo = getMockedRepoNoRowsError(tt.sqlErr)
			} else {
				authRepo = getMockedMethod(tt.want)
			}

			authService := NewAuthService(&authRepo)
			got, err := authService.UserEmailLoginService(tt.args.credential)
			if err != nil {
				if tt.wantErr {
					if err.Error() != tt.errorType.Error() {
						t.Errorf("UserEmailLoginService() error = %v, wantErr %v", err, tt.errorType)
					}
				} else {
					t.Errorf("UserEmailLoginService() error = %v, wantErr %v", err, tt.errorType)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserEmailLoginService() got = %v, want %v", got, tt.want)
			}
		})
	}
}
