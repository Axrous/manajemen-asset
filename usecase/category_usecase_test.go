package usecase

import (
	"errors"
	"final-project-enigma-clean/model"
	"final_project-enigma-clean/__mock__/repomock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CategoryUsecaseTest struct {
	suite.Suite
	repo    *repomock.CategoryRepoMock
	usecase CategoryUsecase
}

func (suite *CategoryUsecaseTest) SetupTest() {
	suite.repo = new(repomock.CategoryRepoMock)
	suite.usecase = NewCategoryUseCase(suite.repo)
}

func TestCategoryUsecaseSuiteTest(t *testing.T) {
	suite.Run(t, new(CategoryUsecaseTest))
}

func (suite *CategoryUsecaseTest) TestCreate_Success() {

	mockData := model.Category{
		Id:   "",
		Name: "Bergerak",
	}

	suite.repo.On("Save", mockData).Return(nil)
	gotErr := suite.usecase.CreateNew(mockData)
	assert.NoError(suite.T(), gotErr)
}

func (suite *CategoryUsecaseTest) TestCreate_EmptyField() {

	gotErr := suite.usecase.CreateNew(model.Category{
		Id:   "1",
		Name: "",
	})
	assert.Error(suite.T(), gotErr)
}

func (suite *CategoryUsecaseTest) TestCreate_Failed() {
	mockData := model.Category{
		Id:   "",
		Name: "Bergerak",
	}

	suite.repo.On("Save", mockData).Return(errors.New("failed save category"))
	gotErr := suite.usecase.CreateNew(mockData)
	assert.Error(suite.T(), gotErr)
}

func (suite *CategoryUsecaseTest) TestFindAll_Success() {
	mockData := []model.Category{{
		Id:   "",
		Name: "Bergerak",
	},
	}

	suite.repo.On("FindAll").Return(mockData, nil)
	assets, err := suite.usecase.FindAll()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockData, assets)
}

func (suite *CategoryUsecaseTest) TestFindAll_Failed() {
	suite.repo.On("FindAll").Return(nil, errors.New("failed"))
	assets, err := suite.usecase.FindAll()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), assets)
}

func (suite *CategoryUsecaseTest) TestFindById_Success() {
	mockData := model.Category{
		Id:   "1",
		Name: "Bergerak",
	}

	suite.repo.On("FindById", "1").Return(mockData, nil)
	category, err := suite.usecase.FindById("1")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockData, category)
}

func (suite *CategoryUsecaseTest) TestFindById_Failed() {

	suite.repo.On("FindById", "1").Return(model.Category{}, errors.New("failed"))
	category, err := suite.usecase.FindById("1")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Category{}, category)
}

func (suite *CategoryUsecaseTest) TestUpdate_Success() {
	mockData := model.Category{
		Id:   "1",
		Name: "Bergerak",
	}
	suite.repo.On("FindById", "1").Return(mockData, nil)
	suite.repo.On("Update", mockData).Return(nil)
	gotErr := suite.usecase.Update(mockData)
	assert.NoError(suite.T(), gotErr)
}

func (suite *CategoryUsecaseTest) TestUpdate_EmptyField() {

	gotErr := suite.usecase.Update(model.Category{
		Id:   "1",
		Name: "",
	})
	assert.Error(suite.T(), gotErr)
}

func (suite *CategoryUsecaseTest) TestUpdate_InvalidId() {
	mockData := model.Category{
		Id:   "1",
		Name: "Bergerak",
	}
	suite.repo.On("FindById", "1").Return(model.Category{}, errors.New("failed get category"))
	gotErr := suite.usecase.Update(mockData)
	assert.Error(suite.T(), gotErr)
}

func (suite *CategoryUsecaseTest) TestUpdate_Failed() {
	mockData := model.Category{
		Id:   "1",
		Name: "Bergerak",
	}

	suite.repo.On("FindById", "1").Return(mockData, nil)
	suite.repo.On("Update", mockData).Return(errors.New("failed update category"))
	gotErr := suite.usecase.Update(mockData)
	assert.Error(suite.T(), gotErr)
}

func (suite *CategoryUsecaseTest) TestDelete_Success() {
	mockData := model.Category{
		Id:   "1",
		Name: "Bergerak",
	}

	suite.repo.On("FindById", "1").Return(mockData, nil)
	suite.repo.On("Delete", "1").Return(nil)
	gotErr := suite.usecase.Delete("1")
	assert.NoError(suite.T(), gotErr)
}

func (suite *CategoryUsecaseTest) TestDelete_InvalidId() {
	suite.repo.On("FindById", "1").Return(model.Category{}, errors.New("failed get category"))
	gotErr := suite.usecase.Delete("1")
	assert.Error(suite.T(), gotErr)
}

func (suite *CategoryUsecaseTest) TestDelete_Failed() {
	mockData := model.Category{
		Id:   "1",
		Name: "Bergerak",
	}

	suite.repo.On("FindById", "1").Return(mockData, nil)
	suite.repo.On("Delete", "1").Return(errors.New("failed delete"))
	gotErr := suite.usecase.Delete("1")
	assert.Error(suite.T(), gotErr)
}
