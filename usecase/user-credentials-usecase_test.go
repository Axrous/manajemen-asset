package usecase_test

// import (
// 	"errors"
// 	"final-project-enigma-clean/model"
// 	"final-project-enigma-clean/repository"
// 	"final-project-enigma-clean/usecase"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"

// 	"github.com/stretchr/testify/suite"
// )

// type MockUserCredentialsRepository struct {
// 	mock.Mock
// }

// func (m *MockUserCredentialsRepository) FindingUserEmail(email string) (model.UserLoginRequest, error) {
// 	args := m.Called(email)
// 	return args.Get(0).(model.UserLoginRequest), args.Error(1)
// }

// func (m *MockUserCredentialsRepository) UserLogin(user model.UserLoginRequest) (string, error) {
// 	args := m.Called(user)
// 	return args.String(0), args.Error(1)
// }

// type UserCredentialSuite struct {
// 	suite.Suite
// 	repo       *repository.MockUserCredentialsRepository
// 	usecase    usecase.UserCredentialUsecase
// 	repository *MockUserCredentialsRepository
// }

// func (suite *UserCredentialSuite) SetupTest() {
// 	// Initialize usecase and mock repository
// 	suite.repo = new(repository.MockUserCredentialsRepository)
// 	suite.usecase = usecase.NewUserCredentialUsecase(suite.repo)
// }

// func (suite *UserCredentialSuite) TestRegisterUser_Success() {
// 	// Define the expected input for UserRegister as mock.Anything
// 	expectedInput := mock.AnythingOfType("model.UserRegisterRequest")
// 	suite.repo.On("UserRegister", expectedInput).Return(nil)

// 	// Create a user to register
// 	userToRegister := model.UserRegisterRequest{
// 		Email:    "test@example.com",
// 		Password: "Password123!",
// 		Name:     "John Doe",
// 	}

// 	// Call the RegisterUser method of the usecase
// 	err := suite.usecase.RegisterUser(userToRegister)

// 	// Assert that there is no error
// 	assert.NoError(suite.T(), err)

// 	// Assert that the expected method in the mock repository was called
// 	suite.repo.AssertExpectations(suite.T())
// }

// func (suite *UserCredentialSuite) TestRegisterUser_HashPasswordFailed() {
// 	// Prepare valid user data
// 	user := model.UserRegisterRequest{
// 		Email:    "test@example.com",
// 		Password: "Password123!",
// 		Name:     "John Doe",
// 	}

// 	// Expectations
// 	suite.repo.On("UserRegister", mock.Anything).Return(errors.New("error from repository"))

// 	// Call the method to be tested
// 	err := suite.usecase.RegisterUser(user)

// 	// Assertions
// 	suite.Error(err)
// 	suite.repo.AssertExpectations(suite.T())
// }

// func (suite *UserCredentialSuite) TestRegisterUser_RepositoryError() {
// 	// Prepare valid user data
// 	user := model.UserRegisterRequest{
// 		Email:    "test@example.com",
// 		Password: "Password123!",
// 		Name:     "John Doe",
// 	}

// 	// Expectations
// 	suite.repo.On("UserRegister", mock.AnythingOfType("model.UserRegisterRequest")).Return(errors.New("error from repository"))

// 	// Call the method to be tested
// 	err := suite.usecase.RegisterUser(user)

// 	// Assertions
// 	suite.Error(err)
// 	suite.repo.AssertExpectations(suite.T())

// login area
// func TestUserCredentialSuite(t *testing.T) {
// 	suite.Run(t, new(UserCredentialSuite))
// }
// =======
// }

// // login area
// func (suite *UserCredentialSuite) TestLoginUser_EmptyEmail() {
// 	// Prepare an empty email in the userlogin request
// 	userlogin := model.UserLoginRequest{
// 		Email:    "",
// 		Password: "Password123!",
// 	}

// 	// Expect that FindingUserEmail should not be called
// 	suite.repository.On("FindingUserEmail", userlogin.Email).Times(0)

// 	// Call the method to be tested
// 	id, err := suite.usecase.LoginUser(userlogin)

// 	// Assertions
// 	assert.Error(suite.T(), err)
// 	assert.Contains(suite.T(), err.Error(), "Email is required")
// 	assert.Equal(suite.T(), "", id)

// 	// Verify that FindingUserEmail was not called
// 	suite.repository.AssertExpectations(suite.T())
// }

// func TestUserCredentialSuite(t *testing.T) {
// 	suite.Run(t, new(UserCredentialSuite))
// }
