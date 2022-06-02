package handlerImpl

import (
	"cf-service/model"
	"cf-service/service/mocks"
	"errors"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getMockedMethod(errReturn bool) mocks.AuthService {
	// mocking the service
	authServiceMock := &mocks.AuthService{}

	var errorReturn error
	var expectedUser *model.UserView

	if errReturn {
		errorReturn = errors.New("authentication failed")
	} else {
		expectedUser = &model.UserView{
			Email: "galabwot@gmail.com",
			Id:    1,
		}
	}

	authServiceMock.On("UserEmailLoginService", mock.AnythingOfType("model.Credential")).
		Return(expectedUser, errorReturn).
		Once()
	return *authServiceMock
}

func TestHandlerAuth_UserEmailLoginHandler(t *testing.T) {

	type args struct {
		resp *httptest.ResponseRecorder
		req  *http.Request
	}
	tests := []struct {
		name   string
		err    bool
		args   args
		status int
	}{
		{
			name: "No error test",
			err:  false,
			args: args{
				resp: httptest.NewRecorder(),
				req:  httptest.NewRequest(http.MethodGet, "/", nil),
			},
			status: 200,
		},
		{
			name: "401 status test",
			err:  true,
			args: args{
				resp: httptest.NewRecorder(),
				req:  httptest.NewRequest(http.MethodGet, "/", nil),
			},
			status: 401,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authService := getMockedMethod(tt.err)
			authHandler := NewAuthHandler(&authService)
			resp := tt.args.resp
			req := tt.args.req
			authHandler.UserEmailLoginHandler(resp, req)

			if resp.Code == 200 {
				if tt.err {
					t.Errorf("got response code %d but expected %d", resp.Code, 200)
				}
			} else {
				if !tt.err {
					t.Errorf("got response code %d but expected %d", 200, 401)
				}
			}

		})
	}
}
