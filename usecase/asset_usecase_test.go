package usecase

import (
	"errors"
	"final-project-enigma-clean/__mock__/repomock"
	"final-project-enigma-clean/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AssetUsecaseTestSuite struct {
	suite.Suite
	repoMock *repomock.AssetRepoMock
	usecase AssetUsecase
}

func (suite *AssetUsecaseTestSuite) SetupTest() {
	suite.repoMock = new(repomock.AssetRepoMock)
	suite.usecase = NewAssetUsecase(suite.repoMock)
}

func TestAssetusecaseTestSuite(t *testing.T)  {
	suite.Run(t, new(AssetUsecaseTestSuite))
}

func (suite *AssetUsecaseTestSuite) TestCreate_Success() {
	payload := model.AssetRequest{
		CategoryId:  "1",
		AssetTypeId: "1",
		Name:        "Laptop",
		Amount:      5,
		Status:      "Ready",
		EntryDate: time.Now(),
		ImgUrl:      "",
	}

	suite.repoMock.On("Save", payload).Return(nil)
	gotError := suite.usecase.Create(payload)
	assert.NoError(suite.T(), gotError)
	assert.Nil(suite.T(), gotError)
}

func (suite *AssetUsecaseTestSuite) TestCreate_EmptyField() {

	//Test name empty
	gotError := suite.usecase.Create(model.AssetRequest{
		CategoryId:  "1",
		AssetTypeId: "1",
		Name:        "",
		Amount:      5,
		Status:      "Ready",
		ImgUrl:      "",
	})
	assert.Error(suite.T(), gotError)

	//test category id or asset type id empty
	gotError = suite.usecase.Create(model.AssetRequest{
		CategoryId:  "",
		AssetTypeId: "",
		Name:        "Laptop",
		Amount:      5,
		Status:      "Ready",
		ImgUrl:      "",
	})
	assert.Error(suite.T(), gotError)

	//test amount minus
	gotError = suite.usecase.Create(model.AssetRequest{
		CategoryId:  "TEST1",
		AssetTypeId: "TEST1",
		Name:        "Laptop",
		Amount:      -1,
		Status:      "Ready",
		ImgUrl:      "",
	})
	assert.Error(suite.T(), gotError)

	//test status empty
	gotError = suite.usecase.Create(model.AssetRequest{
		CategoryId:  "TEST1",
		AssetTypeId: "TEST1",
		Name:        "Laptop",
		Amount:      5,
		Status:      "",
		ImgUrl:      "",
	})
	assert.Error(suite.T(), gotError)
}

func (suite *AssetUsecaseTestSuite) TestCreate_Failed() {
	payload := model.AssetRequest{
		CategoryId:  "1",
		AssetTypeId: "1",
		Name:        "Laptop",
		Amount:      5,
		Status:      "Ready",
		EntryDate: time.Now(),
		ImgUrl:      "",
	}

	suite.repoMock.On("Save", payload).Return(errors.New("failed to create asset"))
	gotError := suite.usecase.Create(payload)
	assert.Error(suite.T(), gotError)
	assert.NotNil(suite.T(), gotError)
}

func (suite *AssetUsecaseTestSuite) TestFindAll_Success() {
	assetMock := []model.Asset{
		{
			ID:       "1",
			Category: model.Category{
				ID:   "1",
				Name: "Bergerak",
			},
			AssetType: model.AssetType{
				ID:   "1",
				Name: "Elektronik",
			},
			Name:      "Laptop",
			Amount: 5,
			Status: "Ready",
			EntryDate: time.Now(),
			ImgUrl: "",
		},
	}

	suite.repoMock.On("FindAll").Return(assetMock, nil)

	assets, err  := suite.usecase.FindAll()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), assetMock, assets) 
}

func (suite *AssetUsecaseTestSuite) TestFindAll_Failed() {

	suite.repoMock.On("FindAll").Return(nil, errors.New("Failed get assets"))

	assets, err  := suite.usecase.FindAll()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), assets)
}

func (suite *AssetUsecaseTestSuite) TestUpdate_Success() {
	payload := model.AssetRequest{
		ID: "1",
		CategoryId:  "1",
		AssetTypeId: "1",
		Name:        "Laptop",
		Amount:      5,
		Status:      "Ready",
		ImgUrl:      "",
	}
	asset := model.Asset{
		ID:        "1",
		Category:  model.Category{
			ID:   "1",
			Name: "Bergerak",
		},
		AssetType: model.AssetType{
			ID:   "1",
			Name: "Elektronik",
		},
		Name:      "Laptop",
		Amount:    5,
		Status:    "Ready",
		EntryDate: time.Time{},
		ImgUrl:    "",
	}
	suite.repoMock.On("FindById", payload.ID).Return(asset, nil)
	suite.repoMock.On("Update", payload).Return(nil)
	gotError := suite.usecase.Update(payload)
	assert.NoError(suite.T(), gotError)
	assert.Nil(suite.T(), gotError)
}

func (suite *AssetUsecaseTestSuite) TestUpdate_EmptyField() {

	//Test name empty
	gotError := suite.usecase.Update(model.AssetRequest{
		CategoryId:  "1",
		AssetTypeId: "1",
		Name:        "",
		Amount:      5,
		Status:      "Ready",
		ImgUrl:      "",
	})
	assert.Error(suite.T(), gotError)

	//test category id or asset type id empty
	gotError = suite.usecase.Update(model.AssetRequest{
		CategoryId:  "",
		AssetTypeId: "",
		Name:        "Laptop",
		Amount:      5,
		Status:      "Ready",
		ImgUrl:      "",
	})
	assert.Error(suite.T(), gotError)

	//test amount minus
	gotError = suite.usecase.Update(model.AssetRequest{
		CategoryId:  "TEST1",
		AssetTypeId: "TEST1",
		Name:        "Laptop",
		Amount:      -1,
		Status:      "Ready",
		ImgUrl:      "",
	})
	assert.Error(suite.T(), gotError)

	//test status empty
	gotError = suite.usecase.Update(model.AssetRequest{
		CategoryId:  "TEST1",
		AssetTypeId: "TEST1",
		Name:        "Laptop",
		Amount:      5,
		Status:      "",
		ImgUrl:      "",
	})
	assert.Error(suite.T(), gotError)
}

func (suite *AssetUsecaseTestSuite) TestUpdate_InvalidId() {
	payload := model.AssetRequest{
		ID: "xx",
		CategoryId:  "1",
		AssetTypeId: "1",
		Name:        "Laptop",
		Amount:      5,
		Status:      "Ready",
		ImgUrl:      "",
	}

	suite.repoMock.On("FindById", "xx").Return(model.Asset{}, errors.New("cannot found asset with id"))
	gotError := suite.usecase.Update(payload)
	assert.NotNil(suite.T(), gotError)
	assert.Error(suite.T(), gotError)
}

func (suite *AssetUsecaseTestSuite) TestUpdate_Failed() {
	payload := model.AssetRequest{
		ID: "1",
		CategoryId:  "1",
		AssetTypeId: "1",
		Name:        "Laptop",
		Amount:      5,
		Status:      "Ready",
		ImgUrl:      "",
	}

	asset := model.Asset{
		ID:        "1",
		Category:  model.Category{
			ID:   "1",
			Name: "Bergerak",
		},
		AssetType: model.AssetType{
			ID:   "1",
			Name: "Elektronik",
		},
		Name:      "Laptop",
		Amount:    5,
		Status:    "Ready",
		EntryDate: time.Time{},
		ImgUrl:    "",
	}

	suite.repoMock.On("FindById", payload.ID).Return(asset, nil)
	suite.repoMock.On("Update", payload).Return(errors.New("failed update asset"))
	gotError := suite.usecase.Update(payload)
	assert.NotNil(suite.T(), gotError)
	assert.Error(suite.T(), gotError)
}

func (suite *AssetUsecaseTestSuite) TestDelete_Success() {
	asset := model.Asset{
		ID:        "1",
		Category:  model.Category{
			ID:   "1",
			Name: "Bergerak",
		},
		AssetType: model.AssetType{
			ID:   "1",
			Name: "Elektronik",
		},
		Name:      "Laptop",
		Amount:    5,
		Status:    "Ready",
		EntryDate: time.Time{},
		ImgUrl:    "",
	}
	suite.repoMock.On("FindById", "1").Return(asset, nil)
	suite.repoMock.On("Delete", "1").Return(nil)
	gotError := suite.usecase.Delete("1")
	assert.NoError(suite.T(), gotError)
	assert.Nil(suite.T(), gotError)
}

func (suite *AssetUsecaseTestSuite) TestDelete_InvalidId() {
	suite.repoMock.On("FindById", "xx").Return(model.Asset{}, errors.New("cannot found asset with id"))
	gotError := suite.usecase.Delete("xx")
	assert.NotNil(suite.T(), gotError)
	assert.Error(suite.T(), gotError)
}

func (suite *AssetUsecaseTestSuite) TestDelete_Failed() {
	asset := model.Asset{
		ID:        "1",
		Category:  model.Category{
			ID:   "1",
			Name: "Bergerak",
		},
		AssetType: model.AssetType{
			ID:   "1",
			Name: "Elektronik",
		},
		Name:      "Laptop",
		Amount:    5,
		Status:    "Ready",
		EntryDate: time.Time{},
		ImgUrl:    "",
	}
	suite.repoMock.On("FindById", "1").Return(asset, nil)
	suite.repoMock.On("Delete", "1").Return(errors.New("failed delete asset"))
	gotError := suite.usecase.Delete("1")
	assert.Error(suite.T(), gotError)
	assert.NotNil(suite.T(), gotError)
}