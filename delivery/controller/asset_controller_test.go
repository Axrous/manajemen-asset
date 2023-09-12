package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"final-project-enigma-clean/__mock__/usecasemock"
	"final-project-enigma-clean/model"
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
	usecase *usecasemock.AssetUsecaseMock
	router *gin.Engine
}

func (suite *AssetControllerTestSuite) SetupTest() {
	suite.usecase = new(usecasemock.AssetUsecaseMock)
	suite.router = gin.Default()
}

func TestAssetusecaseTestSuite(t *testing.T)  {
	suite.Run(t, new(AssetControllerTestSuite))
}

func (suite *AssetControllerTestSuite) TestCreateHandler_Success() {
	mockData := model.AssetRequest{
		Id:          helper.GenerateUUID(),
		CategoryId:  "TEST1",
		AssetTypeId: "TEST1",
		Name:        "Laptop",
		Amount:      5,
		Status:      "Ready",
		EntryDate: time.Time{},
		ImgUrl:      "hehe",
	}

	suite.usecase.On("Create", mockData).Return(nil)
	mockRg := suite.router.Group("/api/v1")
	NewAssetController(suite.usecase, mockRg).Route()

	record := httptest.NewRecorder()

	marshal, err := json.Marshal(mockData)
	assert.NoError(suite.T(), err)

	request, err := http.NewRequest(http.MethodPost, "/api/v1/assets", bytes.NewBuffer(marshal))
	assert.NoError(suite.T(), err)

	suite.router.ServeHTTP(record, request)
	response := record.Body.Bytes()

	var assetResponse model.AssetRequest
	json.Unmarshal(response, &assetResponse)
	assert.Equal(suite.T(), http.StatusCreated, record.Code)
}

func (suite *AssetControllerTestSuite) TestCreateHandler_Failed() {
	mockData := model.AssetRequest{
		Id:          helper.GenerateUUID(),
		CategoryId:  "TEST1",
		AssetTypeId: "TEST1",
		Name:        "Laptop",
		Amount:      5,
		Status:      "Ready",
		EntryDate: time.Time{},
		ImgUrl:      "hehe",
	}

	suite.usecase.On("Create", mockData).Return(errors.New("failed create"))
	mockRg := suite.router.Group("/api/v1")
	NewAssetController(suite.usecase, mockRg).Route()

	record := httptest.NewRecorder()

	marshal, err := json.Marshal(mockData)
	assert.NoError(suite.T(), err)

	request, err := http.NewRequest(http.MethodPost, "/api/v1/assets", bytes.NewBuffer(marshal))
	assert.NoError(suite.T(), err)

	suite.router.ServeHTTP(record, request)
	response := record.Body.Bytes()

	var assetResponse model.AssetRequest
	json.Unmarshal(response, &assetResponse)
	assert.Equal(suite.T(), http.StatusInternalServerError, record.Code)
}

func (suite *AssetControllerTestSuite) TestCreateHandler_BindingError() {
	mockRg := suite.router.Group("/api/v1")
	NewAssetController(suite.usecase, mockRg).Route()
	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "/api/v1/assets", nil)
	suite.router.ServeHTTP(record, request)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, record.Code)
}
