package handlerImpl

import (
	"cf-service/handler"
	"cf-service/model"
	"cf-service/service"
	"encoding/json"
	"log"
	"net/http"
)

/*
@author galab pokharel
*/
type handlerAuth struct {
	AuthService service.AuthService
}

func (h handlerAuth) UserEmailLoginHandler(resp http.ResponseWriter, _ *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	cred := model.Credential{Email: "galabwot@gmail.com"}

	data, err := h.AuthService.UserEmailLoginService(cred)
	if err != nil {
		log.Printf("Got error %c", err)
		resp.WriteHeader(http.StatusUnauthorized)
		errorMessage := model.ErrorTextMessage{ErrorMsg: err.Error()}
		json.NewEncoder(resp).Encode(errorMessage)
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(data)
}

func NewAuthHandler(AuthService service.AuthService) handler.AuthHandler {
	return &handlerAuth{
		AuthService: AuthService,
	}
}
