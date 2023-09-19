package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"final-project-enigma-clean/__mock__/usecasemock"
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type ManageAssetsControllerSuite struct {
	suite.Suite
	controller *ManageAssetController
	usecase    *usecasemock.ManageAssetsMock
	r          *gin.Engine
}

func (suite *ManageAssetsControllerSuite) SetupTest() {
	suite.usecase = new(usecasemock.ManageAssetsMock)
	suite.r = gin.New()
	rg := suite.r.Group("/api/v1")
	suite.controller = NewManageAssetController(suite.usecase, rg)
}

func TestManageAssetControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ManageAssetsControllerSuite))
}

func (suite *ManageAssetsControllerSuite) TestShowAllAssetsSuccess() {
	mockData := []model.ManageAsset{
		{
			Id: "2",
			User: model.UserCredentials{
				ID:       "1",
				Email:    "ellizavad@gmail.com",
				Password: "Nwawd@124",
				Name:     "pal",
				IsActive: true,
			},
			Staff: model.Staff{
				Nik_Staff:    "12312412512",
				Name:         "awdawd",
				Phone_number: "0812461242",
				Address:      "awdawd",
				Birth_date:   time.Time{},
				Img_url:      "jpg",
				Divisi:       "IT",
			},
			SubmissionDate: time.Time{},
			ReturnDate:     time.Time{},
			Detail:         nil,
		},
	}
	suite.usecase.On("ShowAllAsset").Return(mockData, nil)

	suite.controller.Route()
	//record http
	record := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/api/v1/manage-assets/show-all", nil)
	assert.NoError(suite.T(), err)

	suite.r.ServeHTTP(record, request)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}

//
//func (suite *ManageAssetsControllerSuite) TestShowAllAssetsFailed() {
//	mockData := []model.ManageAsset{{
//		Id: "123124",
//		User: model.UserCredentials{
//			ID:       "awdawd",
//			Email:    "awd@gmail.com",
//			Password: "31d22",
//			Name:     "awdawd",
//			IsActive: true,
//		},
//		Staff: model.Staff{
//			Nik_Staff:    "12312412512",
//			Name:         "awdawd",
//			Phone_number: "0812461242",
//			Address:      "awdawd",
//			Birth_date:   time.Time{},
//			Img_url:      "jpg",
//			Divisi:       "IT",
//		},
//		SubmissionDate: time.Time{},
//		ReturnDate:     time.Time{},
//		Detail:         nil,
//	},
//	}
//
//	suite.usecase.On("ShowAllAsset", mockData).Return(errors.New("Failed to get data"))
//	suite.controller.Route()
//
//	record := httptest.NewRecorder()
//
//	request, err := http.NewRequest(http.MethodGet, "/api/v1/manage-assets/show-all", nil)
//	assert.NoError(suite.T(), err)
//
//	suite.r.ServeHTTP(record, request)
//	assert.Equal(suite.T(), http.StatusInternalServerError, record.Code)
//}

func (suite *ManageAssetsControllerSuite) TestCreateNewManageAssetsSuccess() {
	mockData := dto.ManageAssetRequest{
		Id:                   "123213",
		IdUser:               "124214",
		NikStaff:             "12312312",
		SubmisstionDate:      time.Time{},
		ReturnDate:           time.Time{},
		Duration:             2,
		ManageAssetDetailReq: nil,
	}

	suite.usecase.On("CreateTransaction", mockData).Return(nil)
	suite.controller.Route()

	//record http
	record := httptest.NewRecorder()

	marshal, err := json.Marshal(mockData)
	assert.NoError(suite.T(), err)

	request, err := http.NewRequest("POST", "/api/v1/manage-assets/create-new", bytes.NewBuffer(marshal))
	assert.NoError(suite.T(), err)

	suite.r.ServeHTTP(record, request)
	response := record.Body.Bytes()

	var manageAssetResponse model.ManageAsset
	json.Unmarshal(response, &manageAssetResponse)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}

func (suite *ManageAssetsControllerSuite) TestCreate_ErrorJSON() {
	suite.controller.Route()

	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "/api/v1/manage-assets/create-new", nil)
	suite.r.ServeHTTP(record, request)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, record.Code)
}

// test create tx fail
func (suite *ManageAssetsControllerSuite) TestCreate_Failed() {
	mockData := dto.ManageAssetRequest{
		Id:                   "123213",
		IdUser:               "124214",
		NikStaff:             "12312312",
		SubmisstionDate:      time.Time{},
		ReturnDate:           time.Time{},
		Duration:             2,
		ManageAssetDetailReq: nil,
	}

	suite.usecase.On("CreateTransaction", mockData).Return(errors.New("Failed to create"))
	suite.controller.Route()

	record := httptest.NewRecorder()

	marshal, err := json.Marshal(mockData)
	assert.NoError(suite.T(), err)

	req, err := http.NewRequest("POST", "/api/v1/manage-assets/create-new", bytes.NewBuffer(marshal))
	assert.NoError(suite.T(), err)

	//serve http
	suite.r.ServeHTTP(record, req)
	resp := record.Body.Bytes()

	var manageAssetsResponse dto.ManageAssetRequest
	json.Unmarshal(resp, &manageAssetsResponse)
	assert.Equal(suite.T(), 500, record.Code)
}

// find transaction by id success
func (suite *ManageAssetsControllerSuite) TestFindById_Success() {
	mockData := []model.ManageAsset{
		{
			Id: "13",
			User: model.UserCredentials{
				ID:       "123213",
				Email:    "Nahfeu@gmail.com",
				Password: "H@829jsad2",
				Name:     "Haji",
				IsActive: true,
			},
			Staff: model.Staff{
				Nik_Staff:    "123124212512",
				Name:         "awdaw22awd",
				Phone_number: "08125361242",
				Address:      "awdawd",
				Birth_date:   time.Time{},
				Img_url:      "jpg",
				Divisi:       "IT",
			},
			SubmissionDate: time.Time{},
			ReturnDate:     time.Time{},
			Detail:         nil,
		},
		{
			Id: "13",
			User: model.UserCredentials{
				ID:       "123213",
				Email:    "Nahfeu@gmail.com",
				Password: "H@829jsad2",
				Name:     "Haji",
				IsActive: true,
			},
			Staff: model.Staff{
				Nik_Staff:    "1232124212512",
				Name:         "awdaw2aw2awd",
				Phone_number: "08155361242",
				Address:      "awdawdwww",
				Birth_date:   time.Time{},
				Img_url:      "jpg",
				Divisi:       "IT",
			},
			SubmissionDate: time.Time{},
			ReturnDate:     time.Time{},
			Detail:         nil,
		},
	}

	suite.usecase.On("FindByTransactionID", "13").Return(mockData, nil)
	suite.controller.Route()

	record := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/api/v1/manage-assets/find/13", nil)
	assert.NoError(suite.T(), err)

	suite.r.ServeHTTP(record, req)
	assert.Equal(suite.T(), 200, record.Code)
}

// find transaction by id failed
func (suite *ManageAssetsControllerSuite) TestFindTXById_Fail() {

	suite.usecase.On("FindByTransactionID", "13").Return(model.ManageAsset{}, errors.New("failed get asset by id"))
	suite.controller.Route()

	record := httptest.NewRecorder()

	request, err := http.NewRequest(http.MethodGet, "/api/v1/manage-assets/find/13", nil)
	assert.NoError(suite.T(), err)

	suite.r.ServeHTTP(record, request)
	assert.Equal(suite.T(), http.StatusInternalServerError, record.Code)
}

// find tx by name success
func (suite *ManageAssetsControllerSuite) TestFindTXByName_Success() {
	mockData := []model.ManageAsset{
		{
			Id:             "123",
			User:           model.UserCredentials{},
			Staff:          model.Staff{Name: "Adelia"},
			SubmissionDate: time.Time{},
			ReturnDate:     time.Time{},
			Detail:         nil,
		},
	}

	suite.usecase.On("FindTransactionByName", "Adelia").Return(mockData, nil)

	suite.controller.Route()

	recorder := httptest.NewRecorder()
	requestData := map[string]string{
		"Name": "Adelia",
	}
	marshal, err := json.Marshal(requestData)
	assert.NoError(suite.T(), err)

	req, err := http.NewRequest("POST", "/api/v1/manage-assets/find-asset", bytes.NewBuffer(marshal))
	assert.NoError(suite.T(), err)

	suite.r.ServeHTTP(recorder, req)
	resp := recorder.Body.Bytes()

	var response map[string][]model.ManageAsset
	err = json.Unmarshal(resp, &response)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), 200, recorder.Code)
	assert.Equal(suite.T(), mockData, response["Data"])
}
