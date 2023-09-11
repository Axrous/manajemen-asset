package controller

import (
	"final-project-enigma-clean/delivery/middleware"
	"final-project-enigma-clean/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type UserDetailsControllerSuite struct {
	suite.Suite
	controller *UserController
}

type MockUserDetailsUsecase struct {
	UserDetails   model.UserRegisterRequest
	ErrorToReturn error
}

func (m *MockUserDetailsUsecase) RegisterUser(user model.UserRegisterRequest) error {
	m.UserDetails = user // Implement registration logic here

	if m.ErrorToReturn != nil {
		return m.ErrorToReturn
	}

	return nil
}

func (m *MockUserDetailsUsecase) LoginUser(user model.UserLoginRequest) error {
	// Implement login logic here

	if m.ErrorToReturn != nil {
		return m.ErrorToReturn
	}

	return nil
}

func (m *MockUserDetailsUsecase) NewUserDetails(udetails model.UserRegisterRequest) error {
	m.UserDetails = udetails

	// Check if you should return an error
	if m.ErrorToReturn != nil {
		return m.ErrorToReturn
	}

	return nil
}

func (suite *UserDetailsControllerSuite) SetupTest() {
	// Initialize a Gin router
	router := gin.Default()

	// Create a mock implementation of UserDetailsUsecase for testing
	mockUsecase := &MockUserDetailsUsecase{}

	controller := &UserController{
		gin:    router,
		userUC: mockUsecase, // Initialize with the mock implementation
		// Initialize other controller properties here
	}
	controller.Route() // Set up routes

	// Add the AuthMiddleware
	router.Use(middleware.AuthMiddleware())

	suite.controller = controller
}

// valid input
func (suite *UserDetailsControllerSuite) TestSaveUserHandler_ValidInput() {
	// Create a request with valid UserDetails JSON
	userDetailsJSON := `{
      "id": "1",
      "user_id": "user123",
      "name": "John Doe",
      "phone_number": "1234567890",
      "address": "123 Main St",
      "email": "john.doe@example.com",
      "birth_date": "2000-01-01T00:00:00Z",
      "img_url": "http://example.com/image.jpg"
  }`

	req, err := http.NewRequest("POST", "/app/save-user", strings.NewReader(userDetailsJSON))
	suite.Require().NoError(err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	resp := httptest.NewRecorder()

	// Serve request
	suite.controller.gin.ServeHTTP(resp, req)

	// Check the response status code
	suite.Equal(http.StatusOK, resp.Code)

	// Validate
	expectedJSON := `{"Data":` + userDetailsJSON + "}"
	suite.JSONEq(expectedJSON, resp.Body.String())

	suite.Equal(suite.controller.userUC.(*MockUserDetailsUsecase).UserDetails.ID, "1")
	suite.Equal(suite.controller.userUC.(*MockUserDetailsUsecase).UserDetails.Name, "John Doe")
}

// 400 error code
func (suite *UserDetailsControllerSuite) TestSaveUserHandler_ErrorResponse() {
	// Create a mock Gin context for testing
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Call the SaveUserHandler function
	suite.controller.RegisterUserHandler(c)

	// Check the response status code (it should be 400)
	suite.Equal(http.StatusBadRequest, w.Code)

	// Check the response JSON, which should contain the error message
	expectedJSON := `{"Error":"invalid request"}`
	suite.JSONEq(expectedJSON, w.Body.String())
}

func TestUserDetailsControllerSuite(t *testing.T) {
	suite.Run(t, new(UserDetailsControllerSuite))
}
