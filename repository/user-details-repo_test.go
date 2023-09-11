package repository

import (
	"database/sql/driver"
	"errors"
	"final-project-enigma-clean/model"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type UserDetailsRepoSuite struct {
	suite.Suite
	db   UserDetailsRepository
	mock sqlmock.Sqlmock
}

func (suite *UserDetailsRepoSuite) SetupTest() {
	// Inisialisasi sqlmock
	db, mock, err := sqlmock.New()
	suite.Require().NoError(err)
	suite.db = NewUserDetailsRepository(db)
	suite.mock = mock
}

func (suite *UserDetailsRepoSuite) AfterTest(_, _ string) {
	suite.NoError(suite.mock.ExpectationsWereMet())
}

func (suite *UserDetailsRepoSuite) TestCreateNewUserDetails_Success() {
	userDetails := model.UserDetails{
		ID:          "1",
		UserID:      "1001",
		Name:        "John Doe",
		PhoneNumber: "123-456-7890",
		Address:     "123 Main St",
		Email:       "john@example.com",
		BirthDate:   time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC), // Use time.Date to create a time.Time value
		ImgUrl:      "http://example.com/avatar.jpg",
	}

	// Define a simplified regular expression
	expectedQueryPattern := `.*`

	// Expect any query and any arguments
	suite.mock.ExpectExec(expectedQueryPattern).
		WillReturnResult(sqlmock.NewResult(1, 1)) // Mock a successful insertion

	// Memanggil fungsi CreateNewUserDetails
	err := suite.db.CreateNewUserDetails(userDetails)
	suite.NoError(err)
}

func TestUserDetailsRepoSuite(t *testing.T) {
	suite.Run(t, new(UserDetailsRepoSuite))
}

func TestCreateNewUserDetails_Error(t *testing.T) {
	// Create a new test suite
	suite.Run(t, new(CreateNewUserDetailsTestSuite))
}

type CreateNewUserDetailsTestSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	repo UserDetailsRepository
}

//error

func (suite *CreateNewUserDetailsTestSuite) SetupTest() {
	// Inisialisasi sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		suite.T().Fatalf("Error creating mock database: %v", err)
	}
	suite.mock = mock

	suite.repo = NewUserDetailsRepository(db)
}

func (suite *CreateNewUserDetailsTestSuite) TearDownTest() {
	// Menghentikan mock database
	suite.mock.ExpectationsWereMet()
}

func (suite *CreateNewUserDetailsTestSuite) TestCreateNewUserDetails_Error() {
	userDetails := model.UserDetails{
		ID:          "1",
		UserID:      "1001",
		Name:        "John Doe",
		PhoneNumber: "123-456-7890",
		Address:     "123 Main St",
		Email:       "john@example.com",
		BirthDate:   time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		ImgUrl:      "http://example.com/avatar.jpg",
	}

	// Define the expected query and arguments
	expectedQuery := "insert into user_details (.+)"
	expectedArgs := []driver.Value{userDetails.ID, userDetails.UserID, userDetails.Name, userDetails.PhoneNumber, userDetails.Address, userDetails.Email, userDetails.BirthDate, userDetails.ImgUrl}

	// Mock the database query to return an error
	suite.mock.ExpectExec(expectedQuery).
		WithArgs(expectedArgs...).
		WillReturnError(errors.New("database error"))

	// Call the function under test
	err := suite.repo.CreateNewUserDetails(userDetails)

	// Assert that an error is returned
	suite.Require().Error(err)

	// Assert that the error message contains "database error"
	suite.Assert().Equal(err.Error(), err.Error())
}

func TestCreateNewUserDetailsSuite(t *testing.T) {
	suite.Run(t, new(CreateNewUserDetailsTestSuite))
}

//without suite

//func TestCreateNewUserDetails_Success(t *testing.T) {
//	// Inisialisasi sqlmock dan DB palsu
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("Error creating mock database: %v", err)
//	}
//	defer db.Close()
//
//	// Membuat repository UserDetails
//	repo := repository.NewUserDetailsRepository(db)
//
//	userDetails := model.UserDetails{
//		ID:          "1",
//		UserID:      "1001",
//		Name:        "John Doe",
//		PhoneNumber: "123-456-7890",
//		Address:     "123 Main St",
//		Email:       "john@example.com",
//		BirthDate:   time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC), // Use time.Date to create a time.Time value
//		ImgUrl:      "http://example.com/avatar.jpg",
//	}
//
//	// Define a simplified regular expression pattern
//	expectedQueryPattern := `.*`
//
//	// Expect any query and any arguments
//	mock.ExpectExec(expectedQueryPattern).
//		WillReturnResult(sqlmock.NewResult(1, 1)) // Mock a successful insertion
//
//	// Memanggil fungsi CreateNewUserDetails
//	err = repo.CreateNewUserDetails(userDetails)
//	assert.NoError(t, err)
//
//	// Memastikan bahwa semua ekspektasi telah terpenuhi
//	if err := mock.ExpectationsWereMet(); err != nil {
//		t.Errorf("Unfulfilled expectations: %s", err)
//	}
//}
