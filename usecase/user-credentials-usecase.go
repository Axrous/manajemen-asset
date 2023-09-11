package usecase

import (
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/repository"
	"final-project-enigma-clean/util/helper"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type UserCredentialUsecase interface {
	RegisterUser(user model.UserRegisterRequest) error
	LoginUser(user model.UserLoginRequest) (string, error)
	FindingUserEmail(email string) (userlogin model.UserLoginRequest, err error)
}

type userDetailUsecase struct {
	udetailsRepo repository.UserCredentialsRepository
}

// register user business logic
func (u *userDetailUsecase) RegisterUser(user model.UserRegisterRequest) error {
	//TODO implement me

	//validate struct
	val := validator.New()
	err := val.Struct(user)
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

	//password requirement area
	if len(user.Password) < 6 {
		return fmt.Errorf("Password must contain at least six number")
	}
	if !helper.ContainsUppercase(user.Password) {
		return fmt.Errorf("Password must contain at least one uppercase letter")
	}

	if !helper.ContainsSpecialChar(user.Password) {
		return fmt.Errorf("Password must contain at least one special character")
	}

	//generate uuid for user id
	user.ID = helper.GenerateUUID()

	//hash password using bcrypt
	hashedPass, err := helper.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPass

	//save
	if err = u.udetailsRepo.UserRegister(user); err != nil {
		return err
	}
	return nil
}

// login business logic
// Di dalam usecase layer, perbarui LoginUser
func (u *userDetailUsecase) LoginUser(userlogin model.UserLoginRequest) (string, error) {
	// TODO implement me

	// If email is empty
	if userlogin.Email == "" {
		return "", fmt.Errorf("Email is required")
	}

	// Find user email
	user, err := u.FindingUserEmail(userlogin.Email)
	if err != nil {
		return "", fmt.Errorf("Failed to find email %v", err.Error())
	}

	// Compare password
	if err = helper.ComparePassword(user.Password, userlogin.Password); err != nil {
		return "", fmt.Errorf("Invalid Credentials")
	}

	// Mengembalikan email pengguna
	return user.ID, nil
}

func (u *userDetailUsecase) FindingUserEmail(email string) (user model.UserLoginRequest, err error) {
	//TODO implement me

	return u.udetailsRepo.FindUserEmail(email)
}

func NewUserCredentialUsecase(udetailsRepo repository.UserCredentialsRepository) UserCredentialUsecase {
	return &userDetailUsecase{
		udetailsRepo: udetailsRepo,
	}
}
