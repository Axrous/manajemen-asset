package repomock

import (
	"final-project-enigma-clean/model"
	"github.com/stretchr/testify/mock"
)

type MockUserCredentialsRepository struct {
	mock.Mock
}

func (m *MockUserCredentialsRepository) FindUserEmailPass(email string) (userPass model.ChangePasswordRequest, err error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockUserCredentialsRepository) ChangePassword(email, newpass string) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockUserCredentialsRepository) GetUserPassword(email string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockUserCredentialsRepository) CheckEmailExist(email string) bool {
	//TODO implement me
	panic("implement me")
}

func (m *MockUserCredentialsRepository) UserLogin(user model.UserLoginRequest) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}

func (m *MockUserCredentialsRepository) FindUserEmail(email string) (user model.UserLoginRequest, err error) {
	args := m.Called(email)
	return args.Get(0).(model.UserLoginRequest), args.Error(1)
}

func (m *MockUserCredentialsRepository) UserRegister(user model.UserRegisterRequest) error {
	args := m.Called(user)
	return args.Error(0)
}
