package controller

import (
	"bytes"
	"encoding/json"
	"final-project-enigma-clean/__mock__/usecasemock"
	"final-project-enigma-clean/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CategoryControllerTestSuite struct {
	suite.Suite
	controller *CategoryController
	usecase    *usecasemock.CategoryUsecaseMock
	router     *gin.Engine
}

func (suite *CategoryControllerTestSuite) SetupTest() {
	suite.usecase = new(usecasemock.CategoryUsecaseMock)
	suite.router = gin.New()
	rg := suite.router.Group("/api/v1")
	suite.controller = NewCategoryController(suite.usecase, rg)
}

func TestCategoryControllerTestSuite(t *testing.T) {
	suite.Run(t, new(CategoryControllerTestSuite))
}

func (suite *CategoryControllerTestSuite) TestCreateNewHandler_Success() {
	mockData := model.Category{
		Name: "Bergerak",
	}

	suite.usecase.On("CreateNew", mockData).Return(nil)
	suite.controller.Route()

	marshal, err := json.Marshal(mockData)
	assert.NoError(suite.T(), err)

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "/api/v1/categories", bytes.NewBuffer(marshal))
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), http.StatusCreated, record.Code)
}

func (suite *CategoryControllerTestSuite) TestCreateNewHandler_BindingJson() {

	suite.controller.Route()

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "/api/v1/categories", nil)
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), http.StatusBadRequest, record.Code)
}

//func (suite *CategoryControllerTestSuite) TestCreateNewHandler_Failed() {
//	mockData := model.Category{
//		Name: "Bergerak",
//	}
//
//	suite.usecase.On("CreateNew", mockData).Return(errors.New("failed create category"))
//	suite.controller.Route()
//
//	marshal, err := json.Marshal(mockData)
//	assert.NoError(suite.T(), err)
//
//	record := httptest.NewRecorder()
//	request, err := http.NewRequest(http.MethodPost, "/api/v1/categories", bytes.NewBuffer(marshal))
//	assert.NoError(suite.T(), err)
//
//	request.Header.Set("Content-Type", "application/json")
//	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")
//
//	suite.router.ServeHTTP(record, request)
//	assert.Equal(suite.T(), http.StatusInternalServerError, record.Code)
//}

func (suite *CategoryControllerTestSuite) TestFindAllHandler_Success() {
	mockData := []model.Category{{
		Id:   "1",
		Name: "Bergerak",
	},
	}

	suite.usecase.On("FindAll").Return(mockData, nil)
	suite.controller.Route()

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/api/v1/categories", nil)
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}

//func (suite *CategoryControllerTestSuite) TestFindAllHandler_Failed() {
//
//	suite.usecase.On("FindAll").Return(nil, errors.New("failed"))
//	suite.controller.Route()
//
//	record := httptest.NewRecorder()
//	request, err := http.NewRequest(http.MethodGet, "/api/v1/categories", nil)
//	assert.NoError(suite.T(), err)
//
//	suite.router.ServeHTTP(record, request)
//	assert.Equal(suite.T(), http.StatusInternalServerError, record.Code)
//}

func (suite *CategoryControllerTestSuite) TestFindByIdHandler_Success() {
	mockData := model.Category{
		Id:   "1",
		Name: "Bergerak",
	}
	suite.usecase.On("FindById", "1").Return(mockData, nil)
	suite.controller.Route()

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/api/v1/categories/1", nil)
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}

//func (suite *CategoryControllerTestSuite) TestFindByIdHandler_Failed() {
//
//	suite.usecase.On("FindById", "1").Return(model.Category{}, errors.New("Failed"))
//	suite.controller.Route()
//
//	record := httptest.NewRecorder()
//	request, err := http.NewRequest(http.MethodGet, "/api/v1/categories/1", nil)
//	assert.NoError(suite.T(), err)
//
//	suite.router.ServeHTTP(record, request)
//	assert.Equal(suite.T(), http.StatusInternalServerError, record.Code)
//}

func (suite *CategoryControllerTestSuite) TestUpdateHandler_Success() {
	mockData := model.Category{
		Name: "Bergerak",
	}

	suite.usecase.On("Update", mockData).Return(nil)
	suite.controller.Route()

	marshal, err := json.Marshal(mockData)
	assert.NoError(suite.T(), err)

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPut, "/api/v1/categories", bytes.NewBuffer(marshal))
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}

func (suite *CategoryControllerTestSuite) TestUpdateHandler_BindingJson() {

	suite.controller.Route()

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPut, "/api/v1/categories", nil)
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), http.StatusBadRequest, record.Code)
}

//func (suite *CategoryControllerTestSuite) TestUpdateHandler_Failed() {
//	mockData := model.Category{
//		Name: "Bergerak",
//	}
//
//	suite.usecase.On("Update", mockData).Return(errors.New("failed create category"))
//	suite.controller.Route()
//
//	marshal, err := json.Marshal(mockData)
//	assert.NoError(suite.T(), err)
//
//	record := httptest.NewRecorder()
//	request, err := http.NewRequest(http.MethodPut, "/api/v1/categories", bytes.NewBuffer(marshal))
//	assert.NoError(suite.T(), err)
//
//	suite.router.ServeHTTP(record, request)
//	assert.Equal(suite.T(), http.StatusInternalServerError, record.Code)
//}

func (suite *CategoryControllerTestSuite) TestDeleteHandler_Success() {

	suite.usecase.On("Delete", "1").Return(nil)
	suite.controller.Route()

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodDelete, "/api/v1/categories/1", nil)
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}

//func (suite *CategoryControllerTestSuite) TestDeleteHandler_Failed() {
//
//	suite.usecase.On("Delete", "1").Return(errors.New("Failed"))
//	suite.controller.Route()
//
//	record := httptest.NewRecorder()
//	request, err := http.NewRequest(http.MethodDelete, "/api/v1/categories/1", nil)
//	assert.NoError(suite.T(), err)
//
//	suite.router.ServeHTTP(record, request)
//	assert.Equal(suite.T(), http.StatusInternalServerError, record.Code)
//}
