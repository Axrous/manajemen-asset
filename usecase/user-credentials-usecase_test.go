package usecase_test

import (
	"errors"
	"final-project-enigma-clean/__mock__/repomock"
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/suite"
)

type MockUserCredentialsRepository struct {
	mock.Mock
}

type UserCredentialSuite struct {
	suite.Suite
	repo       *repomock.MockUserCredentialsRepository
	usecase    usecase.UserCredentialUsecase
	repository *MockUserCredentialsRepository
}

func (suite *UserCredentialSuite) SetupTest() {
	suite.repo = new(repomock.MockUserCredentialsRepository)
	suite.usecase = usecase.NewUserCredentialUsecase(suite.repo)
}

func (suite *UserCredentialSuite) TestRegisterUser_Success() {
	expectedInput := mock.AnythingOfType("model.UserRegisterRequest")
	suite.repo.On("UserRegister", expectedInput).Return(nil)

	// Create a user to register
	userToRegister := model.UserRegisterRequest{
		Email:    "test@example.com",
		Password: "Password123!",
		Name:     "John Doe",
	}

	err := suite.usecase.RegisterUser(userToRegister)

	assert.NoError(suite.T(), err)

	suite.repo.AssertExpectations(suite.T())
}

func (suite *UserCredentialSuite) TestRegisterUser_HashPasswordFailed() {
	// Prepare valid user data
	user := model.UserRegisterRequest{
		Email:    "test@example.com",
		Password: "Password123!",
		Name:     "John Doe",
	}

	suite.repo.On("UserRegister", mock.Anything).Return(errors.New("error from repository"))

	err := suite.usecase.RegisterUser(user)

	// Assertions
	suite.Error(err)
	suite.repo.AssertExpectations(suite.T())
}

func (suite *UserCredentialSuite) TestRegisterUser_RepositoryError() {
	// Prepare valid user data
	user := model.UserRegisterRequest{
		Email:    "test@example.com",
		Password: "Password123!",
		Name:     "John Doe",
	}

	suite.repo.On("UserRegister", mock.AnythingOfType("model.UserRegisterRequest")).Return(errors.New("error from repository"))

	err := suite.usecase.RegisterUser(user)

	// Assertions
	suite.Error(err)
	suite.repo.AssertExpectations(suite.T())

}

func (suite *UserCredentialSuite) TestRegisterUser_InvalidEmailFormat() {
	// Membuat data user dengan email yang salah
	user := model.UserRegisterRequest{
		ID:       "1",
		Email:    "invalid-email",
		Password: "password",
		Name:     "John Doe",
		IsActive: true,
	}

	suite.repo.On("UserRegister", user).
		Return(errors.New("Invalid email format"))

	err := suite.usecase.RegisterUser(user)

	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "Invalid email")


	// suite.repo.AssertExpectations(suite.T())

}

func TestUserCredentialSuite(t *testing.T) {
	suite.Run(t, new(UserCredentialSuite))
}
