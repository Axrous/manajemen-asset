package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"final-project-enigma-clean/__mock__/usecasemock"
	"final-project-enigma-clean/model"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StaffControllerTestSuite struct {
	suite.Suite
	controller *StaffController
	usecase    *usecasemock.StaffUsecaseMock
	router     *gin.Engine
}

func (suite *StaffControllerTestSuite) SetupTest() {
	suite.usecase = new(usecasemock.StaffUsecaseMock)
	suite.router = gin.New()
	rg := suite.router.Group("/api/v1")
	suite.controller = NewStaffController(suite.usecase, rg)
}

func TestStaffControllerTestSuite(t *testing.T) {
	suite.Run(t, new(StaffControllerTestSuite))
}

func (suite *StaffControllerTestSuite) TestCreateNewHandler_Success() {
	mockData := model.Staff{
		Nik_Staff:    "116511777",
		Name:         "rizki",
		Phone_number: "082767246122",
		Address:      "Pku",
		Birth_date:   time.Time{},
		Img_url:      "hdhagd.png",
		Divisi:       "IT",
	}

	suite.usecase.On("CreateNew", mockData).Return(nil)
	suite.controller.Route()

	marshal, err := json.Marshal(mockData)
	assert.NoError(suite.T(), err)

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "/api/v1/staffs", bytes.NewBuffer(marshal))
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), 200, record.Code)
}

func (suite *StaffControllerTestSuite) TestCreateNewHandler_BindingJson() {

	suite.controller.Route()

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "/api/v1/staffs", nil)
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), 400, record.Code)
}

func (suite *StaffControllerTestSuite) TestCreateNewHandler_Failed() {
	mockData := model.Staff{
		Name: "Bergerak",
	}

	suite.usecase.On("CreateNew", mockData).Return(errors.New("failed create type asset"))
	suite.controller.Route()

	marshal, err := json.Marshal(mockData)
	assert.NoError(suite.T(), err)

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "/api/v1/staffs", bytes.NewBuffer(marshal))
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), 500, record.Code)
}

func (suite *StaffControllerTestSuite) TestGetByNameHandler_Failed() {

	suite.usecase.On("FindByName", "rizki").Return(nil, errors.New("failed"))
	suite.controller.Route()

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/api/v1/staffs/name/rizki", nil)
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), 500, record.Code)
}

func (suite *StaffControllerTestSuite) TestFindByIdHandler_Success() {
	mockData := model.Staff{
		Nik_Staff:    "116511777",
		Name:         "rizki",
		Phone_number: "082767246122",
		Address:      "Pku",
		Birth_date:   time.Time{},
		Img_url:      "hdhagd.png",
		Divisi:       "IT",
	}
	suite.usecase.On("FindById", "1").Return(mockData, nil)
	suite.controller.Route()

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/api/v1/staffs/1", nil)
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), 200, record.Code)
}

func (suite *StaffControllerTestSuite) TestFindByIdHandler_Failed() {

	suite.usecase.On("FindById", "1").Return(model.TypeAsset{}, errors.New("Failed"))
	suite.controller.Route()

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/api/v1/staffs/1", nil)
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), 500, record.Code)
}

func (suite *StaffControllerTestSuite) TestUpdateHandler_Success() {
	mockData := model.Staff{
		Nik_Staff:    "116511777",
		Name:         "rizki",
		Phone_number: "082767246122",
		Address:      "Pku",
		Birth_date:   time.Time{},
		Img_url:      "hdhagd.png",
		Divisi:       "IT",
	}

	suite.usecase.On("Update", mockData).Return(nil)
	suite.controller.Route()

	marshal, err := json.Marshal(mockData)
	assert.NoError(suite.T(), err)

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPut, "/api/v1/staffs", bytes.NewBuffer(marshal))
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), 200, record.Code)
}

func (suite *StaffControllerTestSuite) TestUpdateHandler_BindingJson() {

	suite.controller.Route()

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPut, "/api/v1/staffs", nil)
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), 400, record.Code)
}

func (suite *StaffControllerTestSuite) TestUpdateHandler_Failed() {
	mockData := model.Staff{
		Nik_Staff:    "116511777",
		Name:         "rizki",
		Phone_number: "082767246122",
		Address:      "Pku",
		Birth_date:   time.Time{},
		Img_url:      "hdhagd.png",
		Divisi:       "IT",
	}

	suite.usecase.On("Update", mockData).Return(errors.New("failed create typeAsset"))
	suite.controller.Route()

	marshal, err := json.Marshal(mockData)
	assert.NoError(suite.T(), err)

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPut, "/api/v1/staffs", bytes.NewBuffer(marshal))
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), 500, record.Code)
}
func (suite *StaffControllerTestSuite) TestGetByNameHandler_Success() {
	mockData := []model.Staff{{
		Nik_Staff:    "1",
		Name:         "Rizki",
		Phone_number: "9897986787856",
		Address:      "PkU",
		Birth_date:   time.Time{},
		Img_url:      "ghj.jpg",
		Divisi:       "ghjgjh",
	},
	}

	suite.usecase.On("FindByName", "Rizki").Return(mockData, nil)
	suite.controller.Route()

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/api/v1/staffs/name/Rizki", nil)
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), 200, record.Code)
}
func (suite *StaffControllerTestSuite) TestDeleteHandler_Success() {

	suite.usecase.On("Delete", "1").Return(nil)
	suite.controller.Route()

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodDelete, "/api/v1/staffs/1", nil)
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), 200, record.Code)
}

func (suite *StaffControllerTestSuite) TestDeleteHandler_Failed() {

	suite.usecase.On("Delete", "1").Return(errors.New("Failed"))
	suite.controller.Route()

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodDelete, "/api/v1/staffs/1", nil)
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), 500, record.Code)
}
