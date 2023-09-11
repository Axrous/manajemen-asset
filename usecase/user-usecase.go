package usecase

import (
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/repository"
	"final-project-enigma-clean/util/helper"
	"github.com/go-playground/validator/v10"
)

type UserUsecase interface {
	UserRegist(request model.UserRegisterRequest) error
	UserLogin(logreq model.UserLoginRequest) (string, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func (u userUsecase) UserRegist(request model.UserRegisterRequest) error {
	//TODO implement me

	//validate struct
	val := validator.New()
	err := val.Struct(request)
	if err != nil {
		// Mengganti pesan error berdasarkan jenis kesalahan
		var errMsg string
		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Email" && err.Tag() == "email" {
				errMsg = "Invalid email format"
				break
			}
		}
		if errMsg == "" {
			errMsg = "Bad request format"
			return err
		}
	}

	//password area
	if !helper.ContainsUppercase(request.Password) {
		return err
	}

	if !helper.ContainsNumber(request.Password) {
		return err
	}

	if !helper.ContainsSpecialChar(request.Password) {
		return err
	}

	//assign and generate uuid
	request.Password = helper.GenerateUUID()

	//generate using bcrypt
	hashedPass, err := helper.HashPassword(request.Password)
	if err != nil {
		return err
	}
	request.Password = hashedPass
	return nil
}

func (u userUsecase) UserLogin(logreq model.UserLoginRequest) (string, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserUsecase(urepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: urepo,
	}
}
