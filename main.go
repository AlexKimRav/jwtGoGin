package main

import (
	"jwtgogin/config"
	"jwtgogin/controller"
	"jwtgogin/helper"
	"jwtgogin/model"
	"jwtgogin/repository"
	"jwtgogin/router"
	"jwtgogin/service"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load environment variables", err)
	}
	db := config.ConnectToDb(&loadConfig)
	validate := validator.New()
	db.Table("users").AutoMigrate(&model.Users{})

	usersRepository := repository.NewUsersRepositoryImpl(db)
	authenticationService := service.NewAuthenticationService(usersRepository, validate)
	authenticationController := controller.NewAutheticationController(authenticationService)
	routes := router.NewRouter(authenticationController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	server_err := server.ListenAndServe()
	helper.ErrorPanic(server_err)

}
