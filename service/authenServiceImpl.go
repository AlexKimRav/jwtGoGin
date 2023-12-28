package service

import (
	"errors"
	"jwtgogin/config"
	"jwtgogin/data/request"
	"jwtgogin/helper"
	"jwtgogin/model"
	"jwtgogin/repository"
	"jwtgogin/utils"

	"github.com/go-playground/validator/v10"
)

type AuthenticationServiceImpl struct {
	UsersRepository repository.UserRepository
	Validate        *validator.Validate
}

func NewAuthenticationService(userRepository repository.UserRepository,
	validate *validator.Validate) AuthenticationService {
	return &AuthenticationServiceImpl{UsersRepository: userRepository,
		Validate: validate}
}

// Login implements AuthenticationService.
func (a *AuthenticationServiceImpl) Login(users request.LoginRequest) (string, error) {
	newUser, userErr := a.UsersRepository.FindByUsername(users.Username)
	if userErr != nil {
		return "", errors.New("invalid username or password")
	}

	config, _ := config.LoadConfig(".")

	verify_error := utils.VerifyPassword(newUser.Password, users.Password)
	if verify_error != nil {
		return "", errors.New("invalid username or password")
	}

	token, err_token := utils.GenerateToken(config.TokenExpiresIn, newUser.Id, config.TokenSecret)
	helper.ErrorPanic(err_token)
	return token, nil
}

// Register implements AuthenticationService.
func (a *AuthenticationServiceImpl) Register(users request.CreateUserRequest) {
	hashedPassword, err := utils.HashPassword(users.Password)
	helper.ErrorPanic(err)

	newUser := model.Users{
		Username: users.Username,
		Email:    users.Email,
		Password: hashedPassword,
	}

	a.UsersRepository.Save(newUser)
}
