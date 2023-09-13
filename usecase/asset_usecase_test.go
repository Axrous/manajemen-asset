package usecase

import (
	"errors"
	"final-project-enigma-clean/__mock__/repomock"
	"final-project-enigma-clean/__mock__/usecasemock"
	"final-project-enigma-clean/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AssetUsecaseTestSuite struct {
	suite.Suite
	repoMock    *repomock.AssetRepoMock
	usecase     AssetUsecase
	typeAssetUC *usecasemock.TypeAssetUsecaseMock
	categoryUC  *usecasemock.CategoryUsecaseMock
}

func (suite *AssetUsecaseTestSuite) SetupTest() {
	suite.repoMock = new(repomock.AssetRepoMock)
	suite.typeAssetUC = new(usecasemock.TypeAssetUsecaseMock)
	suite.categoryUC = new(usecasemock.CategoryUsecaseMock)
	suite.usecase = NewAssetUsecase(suite.repoMock, suite.typeAssetUC, suite.categoryUC)
}

func TestAssetusecaseTestSuite(t *testing.T) {
	suite.Run(t, new(AssetUsecaseTestSuite))
}

func (suite *AssetUsecaseTestSuite) TestCreate_Success() {
	payload := model.AssetRequest{
		CategoryId:  "1",
		AssetTypeId: "1",
		Name:        "Laptop",
		Amount:      5,
		Status:      "Ready",
		EntryDate:   time.Now(),
		ImgUrl:      "",
	}

	typeAsset := model.TypeAsset{
		Id:   "1",
		Name: "Elektronik",
	}

	category := model.Category{
		Id:   "1",
		Name: "Bergerak",
	}

	suite.typeAssetUC.On("FindById", payload.AssetTypeId).Return(typeAsset, nil)
	suite.categoryUC.On("FindById", payload.CategoryId).Return(category, nil)
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

	//test category Id or asset type Id empty
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

func (suite *AssetUsecaseTestSuite) TestCreate_InvalidTypeAsset() {
	payload := model.AssetRequest{
		CategoryId:  "1",
		AssetTypeId: "1",
		Name:        "Laptop",
		Amount:      5,
		Status:      "Ready",
		EntryDate:   time.Now(),
		ImgUrl:      "",
	}

	suite.typeAssetUC.On("FindById", payload.AssetTypeId).Return(model.TypeAsset{}, errors.New("failed get asset type"))
	gotError := suite.usecase.Create(payload)
	assert.Error(suite.T(), gotError)
	assert.NotNil(suite.T(), gotError)
}

func (suite *AssetUsecaseTestSuite) TestCreate_InvalidCategory() {
	payload := model.AssetRequest{
		CategoryId:  "1",
		AssetTypeId: "1",
		Name:        "Laptop",
		Amount:      5,
		Status:      "Ready",
		EntryDate:   time.Now(),
		ImgUrl:      "",
	}

	typeAsset := model.TypeAsset{
		Id:   "1",
		Name: "Elektronik",
	}

	suite.typeAssetUC.On("FindById", payload.AssetTypeId).Return(typeAsset, nil)
	suite.categoryUC.On("FindById", payload.CategoryId).Return(model.Category{}, errors.New("failed get category"))
	suite.repoMock.On("Save", payload).Return(errors.New("failed to create asset"))
	gotError := suite.usecase.Create(payload)
	assert.Error(suite.T(), gotError)
	assert.NotNil(suite.T(), gotError)
}

func (suite *AssetUsecaseTestSuite) TestCreate_Failed() {
	payload := model.AssetRequest{
		CategoryId:  "1",
		AssetTypeId: "1",
		Name:        "Laptop",
		Amount:      5,
		Status:      "Ready",
		EntryDate:   time.Now(),
		ImgUrl:      "",
	}

	typeAsset := model.TypeAsset{
		Id:   "1",
		Name: "Elektronik",
	}
	category := model.Category{
		Id:   "1",
		Name: "Bergerak",
	}

	suite.typeAssetUC.On("FindById", payload.AssetTypeId).Return(typeAsset, nil)
	suite.categoryUC.On("FindById", payload.CategoryId).Return(category, nil)
	suite.repoMock.On("Save", payload).Return(errors.New("failed to create asset"))
	gotError := suite.usecase.Create(payload)
	assert.Error(suite.T(), gotError)
	assert.NotNil(suite.T(), gotError)
}

func (suite *AssetUsecaseTestSuite) TestFindAll_Success() {
	assetMock := []model.Asset{
		{
			Id: "1",
			Category: model.Category{
				Id:   "1",
				Name: "Bergerak",
			},
			AssetType: model.TypeAsset{
				Id:   "1",
				Name: "Elektronik",
			},
			Name:      "Laptop",
			Amount:    5,
			Status:    "Ready",
			EntryDate: time.Now(),
			ImgUrl:    "",
		},
	}

	suite.repoMock.On("FindAll").Return(assetMock, nil)

	assets, err := suite.usecase.FindAll()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), assetMock, assets)
}

func (suite *AssetUsecaseTestSuite) TestFindAll_Failed() {

	suite.repoMock.On("FindAll").Return(nil, errors.New("Failed get assets"))

	assets, err := suite.usecase.FindAll()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), assets)
}

func (suite *AssetUsecaseTestSuite) TestUpdate_Success() {
	payload := model.AssetRequest{
		Id:          "1",
		CategoryId:  "1",
		AssetTypeId: "1",
		Name:        "Laptop",
		Amount:      5,
		Status:      "Ready",
		ImgUrl:      "",
	}
	asset := model.Asset{
		Id: "1",
		Category: model.Category{
			Id:   "1",
			Name: "Bergerak",
		},
		AssetType: model.TypeAsset{
			Id:   "1",
			Name: "Elektronik",
		},
		Name:      "Laptop",
		Amount:    5,
		Status:    "Ready",
		EntryDate: time.Time{},
		ImgUrl:    "",
	}

	typeAsset := model.TypeAsset{
		Id:   "1",
		Name: "Elektronik",
	}
	category := model.Category{
		Id:   "1",
		Name: "Bergerak",
	}

	suite.typeAssetUC.On("FindById", payload.AssetTypeId).Return(typeAsset, nil)
	suite.categoryUC.On("FindById", payload.CategoryId).Return(category, nil)
	suite.repoMock.On("FindById", payload.Id).Return(asset, nil)
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

	//test category Id or asset type Id empty
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

func (suite *AssetUsecaseTestSuite) TestUpdate_InvalId() {
	payload := model.AssetRequest{
		Id:          "xx",
		CategoryId:  "1",
		AssetTypeId: "1",
		Name:        "Laptop",
		Amount:      5,
		Status:      "Ready",
		ImgUrl:      "",
	}
	typeAsset := model.TypeAsset{
		Id:   "1",
		Name: "Elektronik",
	}
	category := model.Category{
		Id:   "1",
		Name: "Bergerak",
	}

	suite.typeAssetUC.On("FindById", payload.AssetTypeId).Return(typeAsset, nil)
	suite.categoryUC.On("FindById", payload.CategoryId).Return(category, nil)
	suite.repoMock.On("FindById", "xx").Return(model.Asset{}, errors.New("cannot found asset with Id"))
	gotError := suite.usecase.Update(payload)
	assert.NotNil(suite.T(), gotError)
	assert.Error(suite.T(), gotError)
}

func (suite *AssetUsecaseTestSuite) TestUpdate_InvalidTypeAsset() {
	payload := model.AssetRequest{
		Id:          "1",
		CategoryId:  "1",
		AssetTypeId: "1",
		Name:        "Laptop",
		Amount:      5,
		Status:      "Ready",
		ImgUrl:      "",
	}

	suite.typeAssetUC.On("FindById", payload.AssetTypeId).Return(model.TypeAsset{}, errors.New("failed get type asset"))
	gotError := suite.usecase.Update(payload)
	assert.NotNil(suite.T(), gotError)
	assert.Error(suite.T(), gotError)
}

func (suite *AssetUsecaseTestSuite) TestUpdate_InvalidCategory() {
	payload := model.AssetRequest{
		Id:          "1",
		CategoryId:  "1",
		AssetTypeId: "1",
		Name:        "Laptop",
		Amount:      5,
		Status:      "Ready",
		ImgUrl:      "",
	}
	typeAsset := model.TypeAsset{
		Id:   "1",
		Name: "Elektronik",
	}

	suite.typeAssetUC.On("FindById", payload.AssetTypeId).Return(typeAsset, nil)
	suite.categoryUC.On("FindById", payload.CategoryId).Return(typeAsset, errors.New("failed get category"))
	gotError := suite.usecase.Update(payload)
	assert.NotNil(suite.T(), gotError)
	assert.Error(suite.T(), gotError)
}
func (suite *AssetUsecaseTestSuite) TestUpdate_Failed() {
	payload := model.AssetRequest{
		Id:          "1",
		CategoryId:  "1",
		AssetTypeId: "1",
		Name:        "Laptop",
		Amount:      5,
		Status:      "Ready",
		ImgUrl:      "",
	}

	asset := model.Asset{
		Id: "1",
		Category: model.Category{
			Id:   "1",
			Name: "Bergerak",
		},
		AssetType: model.TypeAsset{
			Id:   "1",
			Name: "Elektronik",
		},
		Name:      "Laptop",
		Amount:    5,
		Status:    "Ready",
		EntryDate: time.Time{},
		ImgUrl:    "",
	}
	typeAsset := model.TypeAsset{
		Id:   "1",
		Name: "Elektronik",
	}
	category := model.Category{
		Id:   "1",
		Name: "Bergerak",
	}

	suite.typeAssetUC.On("FindById", payload.AssetTypeId).Return(typeAsset, nil)
	suite.categoryUC.On("FindById", payload.CategoryId).Return(category, nil)
	suite.repoMock.On("FindById", payload.Id).Return(asset, nil)
	suite.repoMock.On("Update", payload).Return(errors.New("failed update asset"))
	gotError := suite.usecase.Update(payload)
	assert.NotNil(suite.T(), gotError)
	assert.Error(suite.T(), gotError)
}

func (suite *AssetUsecaseTestSuite) TestDelete_Success() {
	asset := model.Asset{
		Id: "1",
		Category: model.Category{
			Id:   "1",
			Name: "Bergerak",
		},
		AssetType: model.TypeAsset{
			Id:   "1",
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

func (suite *AssetUsecaseTestSuite) TestDelete_InvalId() {
	suite.repoMock.On("FindById", "xx").Return(model.Asset{}, errors.New("cannot found asset with Id"))
	gotError := suite.usecase.Delete("xx")
	assert.NotNil(suite.T(), gotError)
	assert.Error(suite.T(), gotError)
}

func (suite *AssetUsecaseTestSuite) TestDelete_Failed() {
	asset := model.Asset{
		Id: "1",
		Category: model.Category{
			Id:   "1",
			Name: "Bergerak",
		},
		AssetType: model.TypeAsset{
			Id:   "1",
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
