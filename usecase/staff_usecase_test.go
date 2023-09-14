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

type StaffUsecaseTestSuite struct {
	suite.Suite
	repo    *repomock.StaffRepoMock
	usecase StaffUseCase
}

func (suite *StaffUsecaseTestSuite) SetupTest() {
	suite.repo = new(repomock.StaffRepoMock)
	suite.usecase = NewStaffUseCase(suite.repo)
}

func TestStafftUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(StaffUsecaseTestSuite))
}

func (suite *StaffUsecaseTestSuite) TestCreate_Success() {
	mockData := model.Staff{
		Nik_Staff:    "11651103422",
		Name:         "Product A",
		Phone_number: "082284163929",
		Address:      "pku",
		Birth_date:   time.Time{},
		Img_url:      "jhj.jpg",
		Divisi:       "IT",
	}
	suite.repo.On("Save", mockData).Return(nil)
	err := suite.usecase.CreateNew(mockData)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)

}

func (suite *StaffUsecaseTestSuite) TestCreate_EmptyField() {

	gotErr := suite.usecase.CreateNew(model.Staff{
		Nik_Staff:    "",
		Name:         "Product A",
		Phone_number: "082284163929",
		Address:      "pku",
		Birth_date:   time.Time{},
		Img_url:      "jhj.jpg",
		Divisi:       "IT",
	})
	gotErr = suite.usecase.CreateNew(model.Staff{
		Nik_Staff:    "qqqq",
		Name:         "",
		Phone_number: "082284163929",
		Address:      "pku",
		Birth_date:   time.Time{},
		Img_url:      "jhj.jpg",
		Divisi:       "IT",
	})
	assert.Error(suite.T(), gotErr)

}

func (suite *StaffUsecaseTestSuite) TestCreate_Failed() {
	mockData := model.Staff{
		Nik_Staff:    "11651103422",
		Name:         "Bergerak",
		Phone_number: "082284163929",
		Address:      "pku",
		Birth_date:   time.Time{},
		Img_url:      "jhj.jpg",
		Divisi:       "IT",
	}

	suite.repo.On("Save", mockData).Return(errors.New("failed save staff"))
	gotErr := suite.usecase.CreateNew(mockData)
	assert.Error(suite.T(), gotErr)
}

func (suite *StaffUsecaseTestSuite) TestFindAll_Success() {
	mockData := []model.Staff{{
		Nik_Staff:    "11651103422",
		Name:         "Bergerak",
		Phone_number: "082284163929",
		Address:      "pku",
		Birth_date:   time.Time{},
		Img_url:      "imh.jpg",
		Divisi:       "IT",
	},
	}

	suite.repo.On("FindByAll").Return(mockData, nil)
	staff, err := suite.usecase.FindByAll()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockData, staff)
}

func (suite *StaffUsecaseTestSuite) TestFindAll_Failed() {
	suite.repo.On("FindByAll").Return(nil, errors.New("failed"))
	staff, err := suite.usecase.FindByAll()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), staff)
}

func (suite *StaffUsecaseTestSuite) TestGetByName_Success() {
	mockData := []model.Staff{{
		Nik_Staff:    "1",
		Name:         "Bergerak",
		Phone_number: "0822",
		Address:      "ddd",
		Birth_date:   time.Time{},
		Img_url:      "ddd.jpg",
		Divisi:       "IT",
	},
	}

	suite.repo.On("FindByName", "Bergerak").Return(mockData, nil)
	staff, err := suite.usecase.FindByName("Bergerak")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockData, staff)
}

func (suite *StaffUsecaseTestSuite) TestGetByName_Failed() {
	suite.repo.On("FindByName", "").Return(nil, errors.New("failed"))
	staff, err := suite.usecase.FindByName("")
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), staff)
}

// func (suite *TypeAssetUsecaseTestSuite) TestFindById_Success() {
// 	mockData := model.TypeAsset{
// 		Id:   "1",
// 		Name: "Bergerak",
// 	}

// 	suite.repo.On("FindById", "1").Return(mockData, nil)
// 	typeAsset, err := suite.usecase.FindById("1")
// 	assert.NoError(suite.T(), err)
// 	assert.Equal(suite.T(), mockData, typeAsset)
// }

// func (suite *TypeAssetUsecaseTestSuite) TestFindById_Failed() {

// 	suite.repo.On("FindById", "1").Return(model.TypeAsset{}, errors.New("failed"))
// 	typeAsset, err := suite.usecase.FindById("1")
// 	assert.Error(suite.T(), err)
// 	assert.Equal(suite.T(), model.TypeAsset{}, typeAsset)
// }

// func (suite *TypeAssetUsecaseTestSuite) TestUpdate_Success() {
// 	mockData := model.TypeAsset{
// 		Id:   "1",
// 		Name: "Bergerak",
// 	}
// 	suite.repo.On("FindById", "1").Return(mockData, nil)
// 	suite.repo.On("Update", mockData).Return(nil)
// 	gotErr := suite.usecase.Update(mockData)
// 	assert.NoError(suite.T(), gotErr)
// }

// func (suite *TypeAssetUsecaseTestSuite) TestUpdate_EmptyField() {

// 	gotErr := suite.usecase.Update(model.TypeAsset{
// 		Id:   "1",
// 		Name: "",
// 	})
// 	assert.Error(suite.T(), gotErr)
// }

// func (suite *TypeAssetUsecaseTestSuite) TestUpdate_InvalidId() {
// 	mockData := model.TypeAsset{
// 		Id:   "1",
// 		Name: "Bergerak",
// 	}
// 	suite.repo.On("FindById", "1").Return(model.TypeAsset{}, errors.New("failed get typeAsset"))
// 	gotErr := suite.usecase.Update(mockData)
// 	assert.Error(suite.T(), gotErr)
// }

// func (suite *TypeAssetUsecaseTestSuite) TestUpdate_Failed() {
// 	mockData := model.TypeAsset{
// 		Id:   "1",
// 		Name: "Bergerak",
// 	}

// 	suite.repo.On("FindById", "1").Return(mockData, nil)
// 	suite.repo.On("Update", mockData).Return(errors.New("failed update typeAsset"))
// 	gotErr := suite.usecase.Update(mockData)
// 	assert.Error(suite.T(), gotErr)
// }

// func (suite *TypeAssetUsecaseTestSuite) TestDelete_Success() {
// 	mockData := model.TypeAsset{
// 		Id:   "1",
// 		Name: "Bergerak",
// 	}

// 	suite.repo.On("FindById", "1").Return(mockData, nil)
// 	suite.repo.On("Delete", "1").Return(nil)
// 	gotErr := suite.usecase.Delete("1")
// 	assert.NoError(suite.T(), gotErr)
// }

// func (suite *TypeAssetUsecaseTestSuite) TestDelete_InvalidId() {
// 	suite.repo.On("FindById", "1").Return(model.TypeAsset{}, errors.New("failed get typeAsset"))
// 	gotErr := suite.usecase.Delete("1")
// 	assert.Error(suite.T(), gotErr)
// }

// func (suite *TypeAssetUsecaseTestSuite) TestDelete_Failed() {
// 	mockData := model.TypeAsset{
// 		Id:   "1",
// 		Name: "Bergerak",
// 	}

// 	suite.repo.On("FindById", "1").Return(mockData, nil)
// 	suite.repo.On("Delete", "1").Return(errors.New("failed delete"))
// 	gotErr := suite.usecase.Delete("1")
// 	assert.Error(suite.T(), gotErr)
// }

// func (suite *TypeAssetUsecaseTestSuite) TestPaging_Success() {
// 	mockData := []model.TypeAsset{
// 		{
// 			Id:   "1",
// 			Name: "Bergerak",
// 		},
// 	}
// 	mockPaging := dto.Paging{
// 		Page:       1,
// 		Size:       5,
// 		TotalRows:  1,
// 		TotalPages: 1,
// 	}
// 	mockPageRequest := dto.PageRequest{
// 		Page: 1,
// 		Size: 5,
// 	}
// 	suite.repo.On("Paging", mockPageRequest).Return(mockData, mockPaging, nil)
// 	gotUom, gotPaging, gotErr := suite.usecase.Paging(mockPageRequest)
// 	assert.Nil(suite.T(), gotErr)
// 	assert.NoError(suite.T(), gotErr)
// 	assert.Equal(suite.T(), mockData, gotUom)
// 	assert.Equal(suite.T(), len(gotUom), 1)
// 	assert.Equal(suite.T(), mockPaging, gotPaging)
// 	assert.Equal(suite.T(), mockPaging.Size, gotPaging.Size)
// }
