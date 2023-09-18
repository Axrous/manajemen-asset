package usecase

import (
	"errors"
	"final-project-enigma-clean/__mock__/repomock"
	"final-project-enigma-clean/__mock__/usecasemock"
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ManageAssetUsecaseTestSuite struct {
	suite.Suite
	staffUC  *usecasemock.StaffUsecaseMock
	assetUC  *usecasemock.AssetUsecaseMock
	repoMock *repomock.ManageAssetRepoMock
	usecase  ManageAssetUsecase
}

func (suite *ManageAssetUsecaseTestSuite) SetupTest() {
	suite.staffUC = new(usecasemock.StaffUsecaseMock)
	suite.assetUC = new(usecasemock.AssetUsecaseMock)
	suite.repoMock = new(repomock.ManageAssetRepoMock)
	suite.usecase = NewManageAssetUsecase(suite.repoMock, suite.staffUC, suite.assetUC)
}

func TestManageAssetUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(ManageAssetUsecaseTestSuite))
}

func (suite *ManageAssetUsecaseTestSuite) TestTransaction_Success() {

	mockData := dto.ManageAssetRequest{
		Id:       "1",
		IdUser:   "1",
		NikStaff: "1",
		// SubmisstionDate:      time.Now(),
		// ReturnDate:           time.Now().AddDate(0, 0, 2),
		Duration: 2,
		ManageAssetDetailReq: []dto.ManageAssetDetailRequest{{
			Id:            "1",
			IdManageAsset: "1",
			IdAsset:       "1",
			TotalItem:     1,
			Status:        "ready",
		}},
	}
	assetMock := model.Asset{
		Id: "1",
		Category: model.Category{
			Id:   "1",
			Name: "a",
		},
		AssetType: model.TypeAsset{
			Id:   "1",
			Name: "a",
		},
		Name:      "a",
		Available: 10,
		Total:     10,
		Status:    "a",
		EntryDate: time.Now(),
		ImgUrl:    "a",
	}

	staffMock := model.Staff{
		Nik_Staff:    "1",
		Name:         "b",
		Phone_number: "123123123",
		Address:      "asd",
		Birth_date:   time.Now(),
		Img_url:      "asd",
		Divisi:       "asd",
	}

	for _, detail := range mockData.ManageAssetDetailReq {
		suite.assetUC.On("FindById", detail.IdAsset).Return(assetMock, nil)
	}

	suite.staffUC.On("FindById", "1").Return(staffMock, nil)
	suite.repoMock.On("CreateTransaction", mockData).Return(nil)
	for _, detail := range mockData.ManageAssetDetailReq {
		suite.assetUC.On("UpdateAvailable", detail.IdAsset, detail.TotalItem).Return(nil)
	}
	err := suite.usecase.CreateTransaction(mockData)
	assert.NoError(suite.T(), err)
}

func (suite *ManageAssetUsecaseTestSuite) TestTransaction_Failed() {

	mockData := dto.ManageAssetRequest{
		Id:       "1",
		IdUser:   "1",
		NikStaff: "1",
		// SubmisstionDate:      time.Now(),
		// ReturnDate:           time.Now().AddDate(0, 0, 2),
		Duration: 2,
		ManageAssetDetailReq: []dto.ManageAssetDetailRequest{{
			Id:            "1",
			IdManageAsset: "1",
			IdAsset:       "1",
			TotalItem:     1,
			Status:        "ready",
		}},
	}
	assetMock := model.Asset{
		Id: "1",
		Category: model.Category{
			Id:   "1",
			Name: "a",
		},
		AssetType: model.TypeAsset{
			Id:   "1",
			Name: "a",
		},
		Name:      "a",
		Available: 10,
		Total:     10,
		Status:    "a",
		EntryDate: time.Now(),
		ImgUrl:    "a",
	}

	staffMock := model.Staff{
		Nik_Staff:    "1",
		Name:         "b",
		Phone_number: "123123123",
		Address:      "asd",
		Birth_date:   time.Now(),
		Img_url:      "asd",
		Divisi:       "asd",
	}

	for _, detail := range mockData.ManageAssetDetailReq {
		suite.assetUC.On("FindById", detail.IdAsset).Return(assetMock, nil)
	}

	suite.staffUC.On("FindById", "1").Return(staffMock, nil)
	suite.repoMock.On("CreateTransaction", mockData).Return(errors.New("failed save transaction"))
	for _, detail := range mockData.ManageAssetDetailReq {
		suite.assetUC.On("UpdateAvailable", detail.IdAsset, detail.TotalItem).Return(nil)
	}
	err := suite.usecase.CreateTransaction(mockData)
	assert.Error(suite.T(), err)
}
func (suite *ManageAssetUsecaseTestSuite) TestTransaction_EmptyField() {

	//nik staff empty
	err := suite.usecase.CreateTransaction(dto.ManageAssetRequest{
		Id:       "1",
		IdUser:   "1",
		NikStaff: "",
		// SubmisstionDate:      time.Now(),
		// ReturnDate:           time.Now().AddDate(0, 0, 2),
		Duration: 2,
		ManageAssetDetailReq: []dto.ManageAssetDetailRequest{{
			Id:            "1",
			IdManageAsset: "1",
			IdAsset:       "1",
			TotalItem:     1,
			Status:        "ready",
		}},
	})
	assert.Error(suite.T(), err)

	//id asset empty
	err = suite.usecase.CreateTransaction(dto.ManageAssetRequest{
		Id:       "1",
		IdUser:   "1",
		NikStaff: "1",
		// SubmisstionDate:      time.Now(),
		// ReturnDate:           time.Now().AddDate(0, 0, 2),
		Duration: 2,
		ManageAssetDetailReq: []dto.ManageAssetDetailRequest{{
			Id:            "1",
			IdManageAsset: "1",
			IdAsset:       "",
			TotalItem:     1,
			Status:        "ready",
		}},
	})
	assert.Error(suite.T(), err)

	//status empty
	err = suite.usecase.CreateTransaction(dto.ManageAssetRequest{
		Id:       "1",
		IdUser:   "1",
		NikStaff: "1",
		// SubmisstionDate:      time.Now(),
		// ReturnDate:           time.Now().AddDate(0, 0, 2),
		Duration: 2,
		ManageAssetDetailReq: []dto.ManageAssetDetailRequest{{
			Id:            "1",
			IdManageAsset: "1",
			IdAsset:       "1",
			TotalItem:     1,
			Status:        "",
		}},
	})
	assert.Error(suite.T(), err)

	//total empty < 0
	err = suite.usecase.CreateTransaction(dto.ManageAssetRequest{
		Id:       "1",
		IdUser:   "1",
		NikStaff: "1",
		// SubmisstionDate:      time.Now(),
		// ReturnDate:           time.Now().AddDate(0, 0, 2),
		Duration: 2,
		ManageAssetDetailReq: []dto.ManageAssetDetailRequest{{
			Id:            "1",
			IdManageAsset: "1",
			IdAsset:       "1",
			TotalItem:     -1,
			Status:        "a",
		}},
	})
	assert.Error(suite.T(), err)
}

func (suite *ManageAssetUsecaseTestSuite) TestTransaction_InvalidIdAsset() {

	mockData := dto.ManageAssetRequest{
		Id:       "1",
		IdUser:   "1",
		NikStaff: "1",
		// SubmisstionDate:      time.Now(),
		// ReturnDate:           time.Now().AddDate(0, 0, 2),
		Duration: 2,
		ManageAssetDetailReq: []dto.ManageAssetDetailRequest{{
			Id:            "1",
			IdManageAsset: "1",
			IdAsset:       "1",
			TotalItem:     1,
			Status:        "ready",
		}},
	}

	staffMock := model.Staff{
		Nik_Staff:    "1",
		Name:         "b",
		Phone_number: "123123123",
		Address:      "asd",
		Birth_date:   time.Now(),
		Img_url:      "asd",
		Divisi:       "asd",
	}

	for _, detail := range mockData.ManageAssetDetailReq {
		suite.assetUC.On("FindById", detail.IdAsset).Return(model.Asset{}, errors.New("failed get asset"))
	}

	suite.staffUC.On("FindById", "1").Return(staffMock, nil)
	suite.repoMock.On("CreateTransaction", mockData).Return(errors.New("failed save transaction"))
	for _, detail := range mockData.ManageAssetDetailReq {
		suite.assetUC.On("UpdateAvailable", detail.IdAsset, detail.TotalItem).Return(nil)
	}
	err := suite.usecase.CreateTransaction(mockData)
	assert.Error(suite.T(), err)
}

func (suite *ManageAssetUsecaseTestSuite) TestTransaction_InvalidStaffId() {

	mockData := dto.ManageAssetRequest{
		Id:       "1",
		IdUser:   "1",
		NikStaff: "1",
		// SubmisstionDate:      time.Now(),
		// ReturnDate:           time.Now().AddDate(0, 0, 2),
		Duration: 2,
		ManageAssetDetailReq: []dto.ManageAssetDetailRequest{{
			Id:            "1",
			IdManageAsset: "1",
			IdAsset:       "1",
			TotalItem:     1,
			Status:        "ready",
		}},
	}
	assetMock := model.Asset{
		Id: "1",
		Category: model.Category{
			Id:   "1",
			Name: "a",
		},
		AssetType: model.TypeAsset{
			Id:   "1",
			Name: "a",
		},
		Name:      "a",
		Available: 10,
		Total:     10,
		Status:    "a",
		EntryDate: time.Now(),
		ImgUrl:    "a",
	}

	for _, detail := range mockData.ManageAssetDetailReq {
		suite.assetUC.On("FindById", detail.IdAsset).Return(assetMock, nil)
	}

	suite.staffUC.On("FindById", "1").Return(model.Staff{}, errors.New("failed get staff"))
	suite.repoMock.On("CreateTransaction", mockData).Return(nil)
	for _, detail := range mockData.ManageAssetDetailReq {
		suite.assetUC.On("UpdateAvailable", detail.IdAsset, detail.TotalItem).Return(nil)
	}
	err := suite.usecase.CreateTransaction(mockData)
	assert.Error(suite.T(), err)
}

func (suite *ManageAssetUsecaseTestSuite) TestTransaction_FailUpdateAsset() {

	mockData := dto.ManageAssetRequest{
		Id:       "1",
		IdUser:   "1",
		NikStaff: "1",
		// SubmisstionDate:      time.Now(),
		// ReturnDate:           time.Now().AddDate(0, 0, 2),
		Duration: 2,
		ManageAssetDetailReq: []dto.ManageAssetDetailRequest{{
			Id:            "1",
			IdManageAsset: "1",
			IdAsset:       "1",
			TotalItem:     1,
			Status:        "ready",
		}},
	}
	assetMock := model.Asset{
		Id: "1",
		Category: model.Category{
			Id:   "1",
			Name: "a",
		},
		AssetType: model.TypeAsset{
			Id:   "1",
			Name: "a",
		},
		Name:      "a",
		Available: 10,
		Total:     10,
		Status:    "a",
		EntryDate: time.Now(),
		ImgUrl:    "a",
	}

	staffMock := model.Staff{
		Nik_Staff:    "1",
		Name:         "b",
		Phone_number: "123123123",
		Address:      "asd",
		Birth_date:   time.Now(),
		Img_url:      "asd",
		Divisi:       "asd",
	}

	for _, detail := range mockData.ManageAssetDetailReq {
		suite.assetUC.On("FindById", detail.IdAsset).Return(assetMock, nil)
	}

	suite.staffUC.On("FindById", "1").Return(staffMock, nil)
	suite.repoMock.On("CreateTransaction", mockData).Return(nil)
	for _, detail := range mockData.ManageAssetDetailReq {
		suite.assetUC.On("UpdateAvailable", detail.IdAsset, detail.TotalItem).Return(errors.New("failed update available"))
	}
	err := suite.usecase.CreateTransaction(mockData)
	assert.Error(suite.T(), err)
}

func (suite *ManageAssetUsecaseTestSuite) TestShowList_Success() {
	mockData := []model.ManageAsset{{
		Id: "1",
		User: model.UserCredentials{
			ID:   "",
			Name: "",
		},
		Staff: model.Staff{
			Nik_Staff: "",
			Name:      "",
		},
		SubmissionDate: time.Time{},
		ReturnDate:     time.Time{},
		Detail: []model.ManageDetailAsset{{
			Id:            "1",
			ManageAssetId: "1",
			Asset: model.Asset{
				Id:   "1",
				Name: "a",
			},
			TotalItem: 1,
			Status:    "a",
		}},
	}}

	suite.repoMock.On("FindAllTransaction").Return(mockData, nil)
	got, err := suite.usecase.ShowAllAsset()
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), got)
}

func (suite *ManageAssetUsecaseTestSuite) TestShowById_Success() {
	mockData := []model.ManageAsset{{
		Id: "1",
		User: model.UserCredentials{
			ID:   "",
			Name: "",
		},
		Staff: model.Staff{
			Nik_Staff: "",
			Name:      "",
		},
		SubmissionDate: time.Time{},
		ReturnDate:     time.Time{},
	}}
	mockDataDetail := []model.ManageDetailAsset{{
		Id:            "1",
		ManageAssetId: "1",
		Asset: model.Asset{
			Id:   "1",
			Name: "a",
		},
		TotalItem: 1,
		Status:    "a",
	}}

	suite.repoMock.On("FindAllByTransId", "1").Return(mockData, mockDataDetail, nil)
	got, err := suite.usecase.FindByTransactionID("1")
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), got)
}

func (suite *ManageAssetUsecaseTestSuite) TestShowById_Failed() {

	suite.repoMock.On("FindAllByTransId", "1").Return(nil, nil, errors.New("failed get data"))
	got, err := suite.usecase.FindByTransactionID("1")
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), got)
}
func (suite *ManageAssetUsecaseTestSuite) TestShowById_IdEmpty() {

	got, err := suite.usecase.FindByTransactionID("")
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), got)
}

func (suite *ManageAssetUsecaseTestSuite) TestShowByName_Success() {
	mockData := []model.ManageAsset{{
		Id: "1",
		User: model.UserCredentials{
			ID:   "123",
			Name: "Jhon",
		},
		Staff: model.Staff{
			Nik_Staff: "123",
			Name:      "jarjit",
		},
		SubmissionDate: time.Time{},
		ReturnDate:     time.Time{},
	}}
	mockDataDetail := []model.ManageDetailAsset{{
		Id:            "1",
		ManageAssetId: "1",
		Asset: model.Asset{
			Id:   "1",
			Name: "a",
		},
		TotalItem: 1,
		Status:    "a",
	}}

	suite.repoMock.On("FindByNameTransaction", "jarjit").Return(mockData, mockDataDetail, nil)
	got, err := suite.usecase.FindTransactionByName("jarjit")
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), got)
}

func (suite *ManageAssetUsecaseTestSuite) TestShowByName_Failed() {

	suite.repoMock.On("FindByNameTransaction", "jarjit").Return(nil, nil, errors.New("failed get data"))
	got, err := suite.usecase.FindTransactionByName("jarjit")
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), got)
}
func (suite *ManageAssetUsecaseTestSuite) TestShowByName_IdEmpty() {

	got, err := suite.usecase.FindTransactionByName("")
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), got)
}

func (suite *ManageAssetUsecaseTestSuite) TestDownload_Success() {
	mockData := []model.ManageAsset{{
		Id: "1",
		User: model.UserCredentials{
			ID:   "",
			Name: "",
		},
		Staff: model.Staff{
			Nik_Staff: "",
			Name:      "",
		},
		SubmissionDate: time.Time{},
		ReturnDate:     time.Time{},
		Detail: []model.ManageDetailAsset{{
			Id:            "1",
			ManageAssetId: "1",
			Asset: model.Asset{
				Id:   "1",
				Name: "a",
			},
			TotalItem: 1,
			Status:    "a",
		}},
	}}

	suite.repoMock.On("FindAllTransaction").Return(mockData, nil)
	got, err := suite.usecase.DownloadAssets()
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), got)
}

func (suite *ManageAssetUsecaseTestSuite) TestDownload_Failed() {
	suite.repoMock.On("FindAllTransaction").Return(nil, errors.New("failed datas transaction"))
	got, err := suite.usecase.DownloadAssets()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), got)
}
