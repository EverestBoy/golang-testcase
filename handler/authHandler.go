package handler

import "net/http"

/*
@author galab pokharel
*/

// AuthHandler : auth api interface
type AuthHandler interface {
	UserEmailLoginHandler(resp http.ResponseWriter, req *http.Request)
}
