package main

import (
	"cf-service/database/dmImpl"
	"cf-service/handler/handlerImpl"
	"cf-service/repository/repoImpl"
	"cf-service/service/authImpl"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
//authHandler = handlerImpl.NewAuthHandler()
)

var (
	dbConnection   = dmImpl.NewPostgresConnection()
	authRepository = repoImpl.NewAuthRepository(dbConnection)
	authService    = authImpl.NewAuthService(authRepository)
	authHandler    = handlerImpl.NewAuthHandler(authService)
)

func main() {
	apiVersion := "v1"
	// logging into the file
	f, err := os.OpenFile("server.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)

	// creating a mux router
	router := mux.NewRouter()
	// router := chi.NewRouter()
	const port = ":8000"

	router.HandleFunc(fmt.Sprintf("/api/%s/test", apiVersion), authHandler.UserEmailLoginHandler).Methods("GET")

	log.Println("Server listening on port ", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
