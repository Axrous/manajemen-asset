package usecasemock

import (
	"final-project-enigma-clean/model"
	"github.com/stretchr/testify/mock"
)

type UserCredentialsMock struct {
	mock.Mock
}

func (u *UserCredentialsMock) LoginUserChangePass(user model.ChangePasswordRequest) (string, error) {
	//TODO implement me
	return "user_token", nil
}

func (u *UserCredentialsMock) ChangePassword(email, newpass string) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserCredentialsMock) ForgotPass(email string) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserCredentialsMock) ForgotPassRequest(email, newPassword, confirmPassword string) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserCredentialsMock) RegisterUser(user model.UserRegisterRequest) error {
	//TODO implement me
	return u.Called(user).Error(0)
}

func (u *UserCredentialsMock) LoginUser(user model.UserLoginRequest) (string, error) {
	// You can return a sample string and nil error for testing purposes.
	// In a real use case, you would provide appropriate values.
	return "user_token", nil
}

func (u *UserCredentialsMock) LoginUserForgotPass(user model.ChangePasswordRequest) (string, error) {
	//TODO implement me
	return "We have sent you an email to ", nil
}

func (u *UserCredentialsMock) FindingUserEmail(email string) (userlogin model.UserLoginRequest, err error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserCredentialsMock) FindingUserEmailPass(email string) (userlogin model.ChangePasswordRequest, err error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserCredentialsMock) ForgotPassword(email, newpass string) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserCredentialsMock) GetUserPassword(email string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserCredentialsMock) EmailExist(email string) bool {
	// Use the mock to determine whether the email exists or not.
	args := u.Called(email)
	return args.Bool(0)
}
