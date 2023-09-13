package usecase

import (
	"errors"
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/repository"
	"final-project-enigma-clean/util/helper"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gookit/slog"
	"regexp"
	"strconv"
)

type UserCredentialUsecase interface {
	RegisterUser(user model.UserRegisterRequest) error
	LoginUser(user model.UserLoginRequest) (string, error)
	LoginUserForgotPass(user model.ChangePasswordRequest) (string, error)
	FindingUserEmail(email string) (userlogin model.UserLoginRequest, err error)
	FindingUserEmailPass(email string) (userlogin model.ChangePasswordRequest, err error)
	ForgotPassword(email, newpass string) error
	GetUserPassword(email string) (string, error)
	EmailExist(email string) bool
}

type userDetailUsecase struct {
	udetailsRepo repository.UserCredentialsRepository
}

func (u *userDetailUsecase) FindingUserEmailPass(email string) (userlogin model.ChangePasswordRequest, err error) {
	//TODO implement me
	return u.udetailsRepo.FindUserEmailPass(email)
}

func (u *userDetailUsecase) LoginUserForgotPass(user model.ChangePasswordRequest) (string, error) {
	//TODO implement me

	// Find user email
	user, err := u.FindingUserEmailPass(user.Email)
	if err != nil {
		return "", err
	}

	//logic otp
	otp, _ := helper.GenerateOTP()
	helper.SendOTPForgotPass(user.Email, strconv.Itoa(otp))
	OTPMap[user.Email] = otp
	slog.Infof("Sending otp to %v", user.Email)

	return user.NewPassword, nil
}

// register user business logic
func (u *userDetailUsecase) RegisterUser(user model.UserRegisterRequest) error {
	//TODO implement me

	//validate struct
	val := validator.New()
	err := val.Struct(user)
	if err != nil {
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

	// Check if email is valid (e.g., gmail.com)
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
	validEmail := regexp.MustCompile(emailPattern)
	if !validEmail.MatchString(user.Email) {
		return errors.New("Invalid email")
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

	//email
	helper.SendEmailRegister(user.Email, user.Name)
	return nil
}

var OTPMap = make(map[string]int)

// login business logic
func (u *userDetailUsecase) LoginUser(userlogin model.UserLoginRequest) (string, error) {
	// TODO implement me

	// Find user email
	user, err := u.FindingUserEmail(userlogin.Email)
	if err != nil {
		return "", err
	}

	// Compare password
	if err = helper.ComparePassword(user.Password, userlogin.Password); err != nil {
		return "", fmt.Errorf("Invalid credential")
	}

	//logic otp
	otp, _ := helper.GenerateOTP()
	helper.SendEmailWithOTP(user.Email, strconv.Itoa(otp))
	OTPMap[user.Email] = otp
	slog.Infof("Sending otp to %v", user.Email)

	// return id
	return user.ID, nil
}

func (u *userDetailUsecase) FindingUserEmail(email string) (user model.UserLoginRequest, err error) {
	//TODO implement me
	return u.udetailsRepo.FindUserEmail(email)
}

func (u *userDetailUsecase) EmailExist(email string) bool {
	//TODO implement me

	var count int
	if !u.udetailsRepo.CheckEmailExist(email) {
		return false
	}

	return count > 0
}

func (u *userDetailUsecase) GetUserPassword(email string) (string, error) {
	//TODO implement me

	return u.udetailsRepo.GetUserPassword(email)
}

func (u *userDetailUsecase) ForgotPassword(email, newpass string) error {
	//TODO implement me

	//update password disini
	u.udetailsRepo.ForgotPassword(email, newpass)
	return nil
}

func NewUserCredentialUsecase(udetailsRepo repository.UserCredentialsRepository) UserCredentialUsecase {
	return &userDetailUsecase{
		udetailsRepo: udetailsRepo,
	}
}
