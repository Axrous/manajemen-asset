package controller

import (
	"bytes"
	"encoding/json"
	"final-project-enigma-clean/__mock__/usecasemock"
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"
	"final-project-enigma-clean/util/helper"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AssetControllerTestSuite struct {
	suite.Suite
	controller *AssetController
	usecase    *usecasemock.AssetUsecaseMock
	router     *gin.Engine
}

func (suite *AssetControllerTestSuite) SetupTest() {
	suite.usecase = new(usecasemock.AssetUsecaseMock)
	suite.router = gin.Default()
}

func TestAssetusecaseTestSuite(t *testing.T) {
	suite.Run(t, new(AssetControllerTestSuite))
}

func (suite *AssetControllerTestSuite) TestCreateHandler_Success() {
	mockData := model.AssetRequest{
		Id:          helper.GenerateUUID(),
		CategoryId:  "TEST1",
		AssetTypeId: "TEST1",
		Name:        "Laptop",
		Status:      "Ready",
		EntryDate:   time.Time{},
		ImgUrl:      "hehe",
		Total:       5,
	}

	suite.usecase.On("Create", mockData).Return(nil)
	mockRg := suite.router.Group("/api/v1")
	NewAssetController(suite.usecase, mockRg).Route()

	record := httptest.NewRecorder()

	marshal, err := json.Marshal(mockData)
	assert.NoError(suite.T(), err)

	request, err := http.NewRequest(http.MethodPost, "/api/v1/assets", bytes.NewBuffer(marshal))
	assert.NoError(suite.T(), err)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	response := record.Body.Bytes()

	var assetResponse model.AssetRequest
	json.Unmarshal(response, &assetResponse)
	assert.Equal(suite.T(), http.StatusCreated, record.Code)
}

//func (suite *AssetControllerTestSuite) TestCreateHandler_Failed() {
//	mockData := model.AssetRequest{
//		Id:          helper.GenerateUUID(),
//		CategoryId:  "TEST1",
//		AssetTypeId: "TEST1",
//		Name:        "Laptop",
//		Status:      "Ready",
//		EntryDate:   time.Time{},
//		ImgUrl:      "hehe",
//		Total:       5,
//	}
//
//	suite.usecase.On("Create", mockData).Return(errors.New("failed create"))
//	mockRg := suite.router.Group("/api/v1")
//	NewAssetController(suite.usecase, mockRg).Route()
//
//	record := httptest.NewRecorder()
//
//	marshal, err := json.Marshal(mockData)
//	assert.NoError(suite.T(), err)
//
//	request, err := http.NewRequest(http.MethodPost, "/api/v1/assets", bytes.NewBuffer(marshal))
//	assert.NoError(suite.T(), err)
//
//	request.Header.Set("Content-Type", "application/json")
//	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")
//
//	suite.router.ServeHTTP(record, request)
//	response := record.Body.Bytes()
//
//	var assetResponse model.AssetRequest
//	json.Unmarshal(response, &assetResponse)
//	assert.Equal(suite.T(), http.StatusInternalServerError, record.Code)
//}

func (suite *AssetControllerTestSuite) TestCreateHandler_BindingError() {
	mockRg := suite.router.Group("/api/v1")
	NewAssetController(suite.usecase, mockRg).Route()
	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "/api/v1/assets", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, record.Code)
}

func (suite *AssetControllerTestSuite) TestListWithNameHandler_Success() {
	mockData := []model.Asset{{
		Id: helper.GenerateUUID(),
		Category: model.Category{
			Id:   "1",
			Name: "Bergerak",
		},
		AssetType: model.TypeAsset{
			Id:   "1",
			Name: "Elektronik",
		},
		Name:      "Laptop",
		Available: 10,
		Status:    "Ready",
		EntryDate: time.Time{},
		ImgUrl:    "upss",
		Total:     10},
	}

	suite.usecase.On("FindByName", "laptop").Return(mockData, nil)
	mockRg := suite.router.Group("/api/v1")
	NewAssetController(suite.usecase, mockRg).Route()

	record := httptest.NewRecorder()

	marshal, err := json.Marshal(mockData)
	assert.NoError(suite.T(), err)

	request, err := http.NewRequest(http.MethodGet, "/api/v1/assets?name=laptop", bytes.NewBuffer(marshal))
	assert.NoError(suite.T(), err)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	response := record.Body.Bytes()

	var assetResponse []model.Asset
	json.Unmarshal(response, &assetResponse)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}

func (suite *AssetControllerTestSuite) TestListHandler_Success() {
	mockData := []model.Asset{{
		Id: helper.GenerateUUID(),
		Category: model.Category{
			Id:   "1",
			Name: "Bergerak",
		},
		AssetType: model.TypeAsset{
			Id:   "1",
			Name: "Elektronik",
		},
		Name:      "Laptop",
		Available: 10,
		Status:    "Ready",
		EntryDate: time.Time{},
		ImgUrl:    "upss",
		Total:     10},
	}

	mockDto := dto.PageRequest{
		Page: 1,
		Size: 5,
	}

	mockPaging := dto.Paging{
		Page:       1,
		Size:       5,
		TotalRows:  5,
		TotalPages: 1,
	}

	suite.usecase.On("Paging", mockDto).Return(mockData, mockPaging, nil)
	mockRg := suite.router.Group("/api/v1")
	NewAssetController(suite.usecase, mockRg).Route()

	record := httptest.NewRecorder()

	marshal, err := json.Marshal(mockData)
	assert.NoError(suite.T(), err)

	request, err := http.NewRequest(http.MethodGet, "/api/v1/assets", bytes.NewBuffer(marshal))
	assert.NoError(suite.T(), err)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	response := record.Body.Bytes()

	var assetResponse []model.Asset
	json.Unmarshal(response, &assetResponse)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}

//func (suite *AssetControllerTestSuite) TestListHandler_Failed() {
//
//	mockDto := dto.PageRequest{
//		Page: 1,
//		Size: 5,
//	}
//	suite.usecase.On("Paging", mockDto).Return(nil, dto.Paging{}, errors.New("failed get assets"))
//	mockRg := suite.router.Group("/api/v1")
//	NewAssetController(suite.usecase, mockRg).Route()
//
//	record := httptest.NewRecorder()
//
//	request, err := http.NewRequest(http.MethodGet, "/api/v1/assets", nil)
//	assert.NoError(suite.T(), err)
//
//	request.Header.Set("Content-Type", "application/json")
//	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")
//
//	suite.router.ServeHTTP(record, request)
//	assert.Equal(suite.T(), http.StatusInternalServerError, record.Code)
//}

//func (suite *AssetControllerTestSuite) TestListByNameHandler_Failed() {
//
//	suite.usecase.On("FindByName", "laptop").Return(nil, errors.New("failed get assets"))
//	mockRg := suite.router.Group("/api/v1")
//	NewAssetController(suite.usecase, mockRg).Route()
//
//	record := httptest.NewRecorder()
//
//	request, err := http.NewRequest(http.MethodGet, "/api/v1/assets?name=laptop", nil)
//	assert.NoError(suite.T(), err)
//
//	request.Header.Set("Content-Type", "application/json")
//	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")
//
//	suite.router.ServeHTTP(record, request)
//	assert.Equal(suite.T(), http.StatusInternalServerError, record.Code)
//}

func (suite *AssetControllerTestSuite) TestFindByIdHandler_Success() {

	mockData := model.Asset{
		Id: helper.GenerateUUID(),
		Category: model.Category{
			Id:   "1",
			Name: "Bergerak",
		},
		AssetType: model.TypeAsset{
			Id:   "1",
			Name: "Elektronik",
		},
		Name:      "Laptop",
		Available: 50,
		Status:    "Ready",
		EntryDate: time.Time{},
		ImgUrl:    "upss",
		Total:     50,
	}

	suite.usecase.On("FindById", "1").Return(mockData, nil)
	mockRg := suite.router.Group("/api/v1")
	NewAssetController(suite.usecase, mockRg).Route()

	record := httptest.NewRecorder()

	request, err := http.NewRequest(http.MethodGet, "/api/v1/assets/1", nil)
	assert.NoError(suite.T(), err)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}

//func (suite *AssetControllerTestSuite) TestFindByIdHandler_Failed() {
//
//	suite.usecase.On("FindById", "1").Return(model.Asset{}, errors.New("failed get asset by id"))
//	mockRg := suite.router.Group("/api/v1")
//	NewAssetController(suite.usecase, mockRg).Route()
//
//	record := httptest.NewRecorder()
//
//	request, err := http.NewRequest(http.MethodGet, "/api/v1/assets/1", nil)
//	assert.NoError(suite.T(), err)
//
//	request.Header.Set("Content-Type", "application/json")
//	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")
//
//	suite.router.ServeHTTP(record, request)
//	assert.Equal(suite.T(), http.StatusInternalServerError, record.Code)
//}

func (suite *AssetControllerTestSuite) TestUpdateHandler_Success() {
	mockData := model.AssetRequest{
		Id:          "1",
		CategoryId:  "TEST1",
		AssetTypeId: "TEST1",
		Name:        "Laptop",
		Available:   5,
		Status:      "Ready",
		EntryDate:   time.Time{},
		ImgUrl:      "hehe",
		Total:       5,
	}

	suite.usecase.On("Update", mockData).Return(nil)
	mockRg := suite.router.Group("/api/v1")
	NewAssetController(suite.usecase, mockRg).Route()

	record := httptest.NewRecorder()

	marshal, err := json.Marshal(mockData)
	assert.NoError(suite.T(), err)

	request, err := http.NewRequest(http.MethodPut, "/api/v1/assets", bytes.NewBuffer(marshal))
	assert.NoError(suite.T(), err)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)

	assert.Equal(suite.T(), http.StatusCreated, record.Code)
}

//func (suite *AssetControllerTestSuite) TestUpdateHandler_Failed() {
//	mockData := model.AssetRequest{
//		Id:          "1",
//		CategoryId:  "TEST1",
//		AssetTypeId: "TEST1",
//		Name:        "Laptop",
//		Available:   5,
//		Status:      "Ready",
//		EntryDate:   time.Time{},
//		ImgUrl:      "hehe",
//		Total:       5,
//	}
//
//	suite.usecase.On("Update", mockData).Return(errors.New("failedddd"))
//	mockRg := suite.router.Group("/api/v1")
//	NewAssetController(suite.usecase, mockRg).Route()
//
//	record := httptest.NewRecorder()
//
//	marshal, err := json.Marshal(mockData)
//	assert.NoError(suite.T(), err)
//
//	request, err := http.NewRequest(http.MethodPut, "/api/v1/assets", bytes.NewBuffer(marshal))
//	assert.NoError(suite.T(), err)
//
//	request.Header.Set("Content-Type", "application/json")
//	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")
//
//	suite.router.ServeHTTP(record, request)
//
//	assert.Equal(suite.T(), http.StatusInternalServerError, record.Code)
//}

//func (suite *AssetControllerTestSuite) TestUpdateHandler_BindingError() {
//	mockRg := suite.router.Group("/api/v1")
//	NewAssetController(suite.usecase, mockRg).Route()
//	record := httptest.NewRecorder()
//	request, err := http.NewRequest(http.MethodPut, "/api/v1/assets", nil)
//	suite.router.ServeHTTP(record, request)
//
//	request.Header.Set("Content-Type", "application/json")
//	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")
//
//	assert.NoError(suite.T(), err)
//	assert.Equal(suite.T(), http.StatusBadRequest, record.Code)
//}

func (suite *AssetControllerTestSuite) TestDeletehandler_Success() {

	suite.usecase.On("Delete", "1").Return(nil)
	mockRg := suite.router.Group("/api/v1")
	NewAssetController(suite.usecase, mockRg).Route()

	record := httptest.NewRecorder()

	request, err := http.NewRequest(http.MethodDelete, "/api/v1/assets/1", nil)
	assert.NoError(suite.T(), err)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")

	suite.router.ServeHTTP(record, request)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}

//func (suite *AssetControllerTestSuite) TestDeletehandler_Failed() {
//
//	suite.usecase.On("Delete", "1").Return(errors.New("failed delete asset"))
//	mockRg := suite.router.Group("/api/v1")
//	NewAssetController(suite.usecase, mockRg).Route()
//
//	record := httptest.NewRecorder()
//
//	request, err := http.NewRequest(http.MethodDelete, "/api/v1/assets/1", nil)
//	assert.NoError(suite.T(), err)
//
//	request.Header.Set("Content-Type", "application/json")
//	request.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjoxNjk0MjgyNzQyLCJleHBfYXQiOiIyMDIzLTA5LTEwVDA3OjA1OjQyLjkzNDc3ODkrMDc6MDAiLCJ1c2VyX2VtYWlsIjoiZWxsaXphdmFkQHBhbC5jb20ifQ.TeRaZw60Rrtp6wHpP5oL7BAHSLxDMBxVcZNtJPHkXYM")
//
//	suite.router.ServeHTTP(record, request)
//	assert.Equal(suite.T(), http.StatusInternalServerError, record.Code)
//}
