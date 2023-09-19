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

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TypeAssetControllerTestSuite struct {
	suite.Suite
	controller *TypeAssetController
	usecase    *usecasemock.TypeAssetUsecaseMock
	router     *gin.Engine
}

func (suite *TypeAssetControllerTestSuite) SetupTest() {
	suite.usecase = new(usecasemock.TypeAssetUsecaseMock)
	suite.router = gin.New()
	rg := suite.router.Group("/api/v1")
	suite.controller = NewTypeAssetController(suite.usecase, rg)
}

func TestTypeAssetControllerTestSuite(t *testing.T) {
	suite.Run(t, new(TypeAssetControllerTestSuite))
}

func (suite *TypeAssetControllerTestSuite) TestCreateNewHandler_Success() {
	mockData := model.TypeAsset{
		Id:   "1",
		Name: "bergerak",
	}

	suite.usecase.On("CreateNew", mockData).Return(nil)
	suite.controller.Route()

	marshal, err := json.Marshal(mockData)
	assert.NoError(suite.T(), err)

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "/api/v1/typeAsset", bytes.NewBuffer(marshal))
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), 200, record.Code)
}

func (suite *TypeAssetControllerTestSuite) TestCreateNewHandler_BindingJson() {

	suite.controller.Route()

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "/api/v1/typeAsset", nil)
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), 400, record.Code)
}

func (suite *TypeAssetControllerTestSuite) TestCreateNewHandler_Failed() {
	mockData := model.TypeAsset{
		Name: "Bergerak",
	}

	suite.usecase.On("CreateNew", mockData).Return(errors.New("failed create type asset"))
	suite.controller.Route()

	marshal, err := json.Marshal(mockData)
	assert.NoError(suite.T(), err)

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "/api/v1/typeAsset", bytes.NewBuffer(marshal))
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), 500, record.Code)
}

func (suite *TypeAssetControllerTestSuite) TestGetByNameHandler_Success() {
	mockData := []model.TypeAsset{{
		Id:   "1",
		Name: "Bergerak",
	},
	}

	suite.usecase.On("FindByName", "Bergerak").Return(mockData, nil)
	suite.controller.Route()

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/api/v1/typeAsset/name/Bergerak", nil)
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), 200, record.Code)
}

//func (suite *TypeAssetControllerTestSuite) TestGetByNameHandler_Failed() {
//
//	suite.usecase.On("FindByName", "Bergerak").Return(nil, errors.New("failed"))
//	suite.controller.Route()
//
//	record := httptest.NewRecorder()
//	request, err := http.NewRequest(http.MethodGet, "/api/v1/typeAsset/name/Bergerak", nil)
//	assert.NoError(suite.T(), err)
//
//	request.Header.Set("Content-Type", "application/json")
//	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")
//
//	suite.router.ServeHTTP(record, request)
//	assert.Equal(suite.T(), 500, record.Code)
//}

func (suite *TypeAssetControllerTestSuite) TestFindByIdHandler_Success() {
	mockData := model.TypeAsset{
		Id:   "1",
		Name: "Bergerak",
	}
	suite.usecase.On("FindById", "1").Return(mockData, nil)
	suite.controller.Route()

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/api/v1/typeAsset/1", nil)
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), 200, record.Code)
}

//func (suite *TypeAssetControllerTestSuite) TestFindByIdHandler_Failed() {
//
//	suite.usecase.On("FindById", "1").Return(model.TypeAsset{}, errors.New("Failed"))
//	suite.controller.Route()
//
//	record := httptest.NewRecorder()
//	request, err := http.NewRequest(http.MethodGet, "/api/v1/typeAsset/1", nil)
//	assert.NoError(suite.T(), err)
//
//	request.Header.Set("Content-Type", "application/json")
//	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")
//
//	suite.router.ServeHTTP(record, request)
//	assert.Equal(suite.T(), 500, record.Code)
//}

func (suite *TypeAssetControllerTestSuite) TestUpdateHandler_Success() {
	mockData := model.TypeAsset{
		Id:   "1",
		Name: "Bergerak",
	}

	suite.usecase.On("Update", mockData).Return(nil)
	suite.controller.Route()

	marshal, err := json.Marshal(mockData)
	assert.NoError(suite.T(), err)

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPut, "/api/v1/typeAsset", bytes.NewBuffer(marshal))
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), 200, record.Code)
}

func (suite *TypeAssetControllerTestSuite) TestUpdateHandler_BindingJson() {

	suite.controller.Route()

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPut, "/api/v1/typeAsset", nil)
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), 400, record.Code)
}

//func (suite *TypeAssetControllerTestSuite) TestUpdateHandler_Failed() {
//	mockData := model.TypeAsset{
//		Id:   "1",
//		Name: "Bergerak",
//	}
//
//	suite.usecase.On("Update", mockData).Return(errors.New("failed create typeAsset"))
//	suite.controller.Route()
//
//	marshal, err := json.Marshal(mockData)
//	assert.NoError(suite.T(), err)
//
//	record := httptest.NewRecorder()
//	request, err := http.NewRequest(http.MethodPut, "/api/v1/typeAsset", bytes.NewBuffer(marshal))
//	assert.NoError(suite.T(), err)
//
//	request.Header.Set("Content-Type", "application/json")
//	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")
//
//	suite.router.ServeHTTP(record, request)
//	assert.Equal(suite.T(), 500, record.Code)
//}

func (suite *TypeAssetControllerTestSuite) TestDeleteHandler_Success() {

	suite.usecase.On("Delete", "1").Return(nil)
	suite.controller.Route()

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodDelete, "/api/v1/typeAsset/1", nil)
	assert.NoError(suite.T(), err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), 200, record.Code)
}

//func (suite *TypeAssetControllerTestSuite) TestDeleteHandler_Failed() {
//
//	suite.usecase.On("Delete", "1").Return(errors.New("Failed"))
//	suite.controller.Route()
//
//	record := httptest.NewRecorder()
//	request, err := http.NewRequest(http.MethodDelete, "/api/v1/typeAsset/1", nil)
//	assert.NoError(suite.T(), err)
//
//	request.Header.Set("Content-Type", "application/json")
//	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")
//
//	suite.router.ServeHTTP(record, request)
//	assert.Equal(suite.T(), 500, record.Code)
//}
