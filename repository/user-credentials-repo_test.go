package repository

import (
	"database/sql"
	"final-project-enigma-clean/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserCredentialsRepositorySuite struct {
	suite.Suite
	db   *sql.DB
	mock sqlmock.Sqlmock
	repo UserCredentialsRepository
}

func (suite *UserCredentialsRepositorySuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.db = db
	suite.mock = mock
	suite.repo = NewUserDetailsRepository(db)
}

func (suite *UserCredentialsRepositorySuite) TearDownTest() {
	suite.db.Close()
}

func (suite *UserCredentialsRepositorySuite) TestUserRegister() {
	// Prepare test data
	user := model.UserRegisterRequest{
		ID:       "1",
		Email:    "test@example.com",
		Password: "password",
		Name:     "John Doe",
		IsActive: true,
	}

	// Expectation: SQLMock will expect an INSERT query with specific arguments
	suite.mock.ExpectExec("insert into user_credential (.+)").
		WithArgs(user.ID, user.Email, user.Password, user.Name, user.IsActive).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the method to be tested
	err := suite.repo.UserRegister(user)

	// Assertion: Ensure that there is no error and that SQLMock expectations are met
	assert.NoError(suite.T(), err)
	assert.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

func (suite *UserCredentialsRepositorySuite) TestUserRegister_Failure_EmptyEmail() {
	// Prepare test data with empty email
	user := model.UserRegisterRequest{
		ID:       "1",
		Email:    "", // Empty email (invalid)
		Password: "password",
		Name:     "John Doe",
		IsActive: true,
	}

	// Expectation: No expectations set, as it's an invalid test case

	// Call the method to be tested
	err := suite.repo.UserRegister(user)

	// Assertion: Ensure that there is an error and SQLMock expectations are not met
	assert.Error(suite.T(), err)
	assert.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

// test user login
func (suite *UserCredentialsRepositorySuite) TestUserLogin() {
	user := model.UserLoginRequest{
		Email: "test@example.com",
	}

	suite.mock.ExpectQuery("select password from user_credential where email = (.+)").
		WithArgs(user.Email).
		WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow("hashed_password"))

	hashedPassword, err := suite.repo.UserLogin(user)

	// Assertion:
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "hashed_password", hashedPassword)
	assert.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

// invalid email
func (suite *UserCredentialsRepositorySuite) TestUserLogin_EmailNotFound() {
	user := model.UserLoginRequest{
		Email: "nonexistent@example.com", // Email yang tidak ada dalam database
	}

	suite.mock.ExpectQuery("select password from user_credential where email = (.+)").
		WithArgs(user.Email).
		WillReturnError(sql.ErrNoRows) // Set error to sql.ErrNoRows

	hashedPassword, err := suite.repo.UserLogin(user)

	// Assertion
	assert.EqualError(suite.T(), err, "Invalid Credential sql: no rows in result set")
	assert.Empty(suite.T(), hashedPassword)
	assert.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

// email found
func (suite *UserCredentialsRepositorySuite) TestFindUserEmail_EmailFound() {
	// Prepare test data
	email := "test@example.com"
	expectedUser := model.UserLoginRequest{
		ID:       "1",
		Email:    email,
		Password: "hashed_password",
	}

	// Expectation: SQLMock will expect a SELECT query with specific arguments
	suite.mock.ExpectQuery("SELECT id, email, password FROM user_credential WHERE email = (.+)").
		WithArgs(email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password"}).
			AddRow(expectedUser.ID, expectedUser.Email, expectedUser.Password))

	user, err := suite.repo.FindUserEmail(email)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedUser, user)
	assert.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

func (suite *UserCredentialsRepositorySuite) TestFindUserEmail_EmailNotFound() {
	// Prepare test data
	email := "nonexistent@example.com"

	// Expectation: SQLMock will expect a SELECT query with specific arguments
	suite.mock.ExpectQuery("SELECT id, email, password FROM user_credential WHERE email = (.+)").
		WithArgs(email).
		WillReturnError(sql.ErrNoRows) // Set error to sql.ErrNoRows

	// Call the method to be tested
	user, err := suite.repo.FindUserEmail(email)

	assert.EqualError(suite.T(), err, "Invalid Credentials")
	assert.Empty(suite.T(), user)
	assert.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}
func TestUserDetailsRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserCredentialsRepositorySuite))
}
