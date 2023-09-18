package usecase

import (
	"errors"
	"final-project-enigma-clean/__mock__/repomock"
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TypeAssetUsecaseTestSuite struct {
	suite.Suite
	repo    *repomock.TypeAssetRepoMock
	usecase TypeAssetUseCase
}

func (suite *TypeAssetUsecaseTestSuite) SetupTest() {
	suite.repo = new(repomock.TypeAssetRepoMock)
	suite.usecase = NewTypeAssetUseCase(suite.repo)
}

func TestTypeAssetUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(TypeAssetUsecaseTestSuite))
}

func (suite *TypeAssetUsecaseTestSuite) TestCreate_Success() {
	mockData := model.TypeAsset{
		Id:   "1",
		Name: "Product A",
	}
	suite.repo.On("Save", mockData).Return(nil)
	err := suite.usecase.CreateNew(mockData)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)

}

func (suite *TypeAssetUsecaseTestSuite) TestCreate_EmptyField() {

	gotErr := suite.usecase.CreateNew(model.TypeAsset{
		Id:   "1",
		Name: "",
	})
	assert.Error(suite.T(), gotErr)
}

func (suite *TypeAssetUsecaseTestSuite) TestCreate_Failed() {
	mockData := model.TypeAsset{
		Id:   "1",
		Name: "Bergerak",
	}

	suite.repo.On("Save", mockData).Return(errors.New("failed save type asset"))
	gotErr := suite.usecase.CreateNew(mockData)
	assert.Error(suite.T(), gotErr)
}

func (suite *TypeAssetUsecaseTestSuite) TestFindAll_Success() {
	mockData := []model.TypeAsset{{
		Id:   "1",
		Name: "Bergerak",
	},
	}

	suite.repo.On("FindAll").Return(mockData, nil)
	typeAsset, err := suite.usecase.FindAll()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockData, typeAsset)
}

func (suite *TypeAssetUsecaseTestSuite) TestFindAll_Failed() {
	suite.repo.On("FindAll").Return(nil, errors.New("failed"))
	typeAsset, err := suite.usecase.FindAll()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), typeAsset)
}

func (suite *TypeAssetUsecaseTestSuite) TestGetByName_Success() {
	mockData := []model.TypeAsset{{
		Id:   "1",
		Name: "Bergerak",
	},
	}

	suite.repo.On("FindByName", "Bergerak").Return(mockData, nil)
	typeAsset, err := suite.usecase.FindByName("Bergerak")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockData, typeAsset)
}

func (suite *TypeAssetUsecaseTestSuite) TestGetByName_Failed() {
	suite.repo.On("FindByName", "").Return(nil, errors.New("failed"))
	typeAsset, err := suite.usecase.FindByName("")
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), typeAsset)
}

func (suite *TypeAssetUsecaseTestSuite) TestFindById_Success() {
	mockData := model.TypeAsset{
		Id:   "1",
		Name: "Bergerak",
	}

	suite.repo.On("FindById", "1").Return(mockData, nil)
	typeAsset, err := suite.usecase.FindById("1")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockData, typeAsset)
}

func (suite *TypeAssetUsecaseTestSuite) TestFindById_Failed() {

	suite.repo.On("FindById", "1").Return(model.TypeAsset{}, errors.New("failed"))
	typeAsset, err := suite.usecase.FindById("1")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.TypeAsset{}, typeAsset)
}

func (suite *TypeAssetUsecaseTestSuite) TestUpdate_Success() {
	mockData := model.TypeAsset{
		Id:   "1",
		Name: "Bergerak",
	}
	suite.repo.On("FindById", "1").Return(mockData, nil)
	suite.repo.On("Update", mockData).Return(nil)
	gotErr := suite.usecase.Update(mockData)
	assert.NoError(suite.T(), gotErr)
}

func (suite *TypeAssetUsecaseTestSuite) TestUpdate_EmptyField() {

	gotErr := suite.usecase.Update(model.TypeAsset{
		Id:   "1",
		Name: "",
	})
	assert.Error(suite.T(), gotErr)
}

func (suite *TypeAssetUsecaseTestSuite) TestUpdate_InvalidId() {
	mockData := model.TypeAsset{
		Id:   "1",
		Name: "Bergerak",
	}
	suite.repo.On("FindById", "1").Return(model.TypeAsset{}, errors.New("failed get typeAsset"))
	gotErr := suite.usecase.Update(mockData)
	assert.Error(suite.T(), gotErr)
}

func (suite *TypeAssetUsecaseTestSuite) TestUpdate_Failed() {
	mockData := model.TypeAsset{
		Id:   "1",
		Name: "Bergerak",
	}

	suite.repo.On("FindById", "1").Return(mockData, nil)
	suite.repo.On("Update", mockData).Return(errors.New("failed update typeAsset"))
	gotErr := suite.usecase.Update(mockData)
	assert.Error(suite.T(), gotErr)
}

func (suite *TypeAssetUsecaseTestSuite) TestDelete_Success() {
	mockData := model.TypeAsset{
		Id:   "1",
		Name: "Bergerak",
	}

	suite.repo.On("FindById", "1").Return(mockData, nil)
	suite.repo.On("Delete", "1").Return(nil)
	gotErr := suite.usecase.Delete("1")
	assert.NoError(suite.T(), gotErr)
}

func (suite *TypeAssetUsecaseTestSuite) TestDelete_InvalidId() {
	suite.repo.On("FindById", "1").Return(model.TypeAsset{}, errors.New("failed get typeAsset"))
	gotErr := suite.usecase.Delete("1")
	assert.Error(suite.T(), gotErr)
}

func (suite *TypeAssetUsecaseTestSuite) TestDelete_Failed() {
	mockData := model.TypeAsset{
		Id:   "1",
		Name: "Bergerak",
	}

	suite.repo.On("FindById", "1").Return(mockData, nil)
	suite.repo.On("Delete", "1").Return(errors.New("failed delete"))
	gotErr := suite.usecase.Delete("1")
	assert.Error(suite.T(), gotErr)
}

func (suite *TypeAssetUsecaseTestSuite) TestPaging_Success() {
	mockData := []model.TypeAsset{
		{
			Id:   "1",
			Name: "Bergerak",
		},
	}
	mockPaging := dto.Paging{
		Page:       1,
		Size:       5,
		TotalRows:  1,
		TotalPages: 1,
	}
	mockPageRequest := dto.PageRequest{
		Page: 1,
		Size: 5,
	}
	suite.repo.On("Paging", mockPageRequest).Return(mockData, mockPaging, nil)
	gotUom, gotPaging, gotErr := suite.usecase.Paging(mockPageRequest)
	assert.Nil(suite.T(), gotErr)
	assert.NoError(suite.T(), gotErr)
	assert.Equal(suite.T(), mockData, gotUom)
	assert.Equal(suite.T(), len(gotUom), 1)
	assert.Equal(suite.T(), mockPaging, gotPaging)
	assert.Equal(suite.T(), mockPaging.Size, gotPaging.Size)
}
