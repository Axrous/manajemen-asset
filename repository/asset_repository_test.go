package repository

import (
	"database/sql"
	"errors"
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"
	"final-project-enigma-clean/util/helper"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AssetRepositoryTestSuite struct {
	suite.Suite
	mockDB *sql.DB
	mockSQL sqlmock.Sqlmock
	repository AssetRepository
}

func (suite *AssetRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDB = db
	suite.mockSQL = mock
	suite.repository = NewAssetRepository(suite.mockDB)
}

func TestAssetRepositoryTestSuite(t *testing.T)  {
	suite.Run(t, new(AssetRepositoryTestSuite))
}

func (suite *AssetRepositoryTestSuite) TestCreate_Success() {
	
	asset := model.AssetRequest{
		Id:         helper.GenerateUUID(),
		CategoryId: "TEST1",
		AssetTypeId:  "TEST2",
		Name:       "Laptop",
		Available: 5,
		EntryDate:  time.Now(),
		ImgUrl:     "nothing",
		Total:     5,
	}


	suite.mockSQL.ExpectExec("insert into asset").WithArgs(
		asset.Id, 
		asset.CategoryId, 
		asset.AssetTypeId, 
		asset.Name,
		asset.Available,
		asset.Status, 
		asset.EntryDate,  
		asset.ImgUrl,
		asset.Total,).WillReturnResult(sqlmock.NewResult(1, 1))

	got := suite.repository.Save(asset)
	assert.NoError(suite.T(), got)
	assert.Nil(suite.T(), got)
}

func (suite *AssetRepositoryTestSuite) TestCreate_Failed() {
	
	asset := model.AssetRequest{
		Id:         helper.GenerateUUID(),
		CategoryId: "TEST1",
		AssetTypeId:"TEST2",
		Name:       "Laptop",
		Total:     5,
		EntryDate:  time.Now(),
		ImgUrl:     "nothing",
	}


	suite.mockSQL.ExpectExec("insert into asset").WithArgs(
		asset.Id, 
		asset.CategoryId, 
		asset.AssetTypeId, 
		asset.Name, 
		asset.Total, 
		asset.Status, 
		asset.EntryDate, 
		asset.ImgUrl,).WillReturnError(errors.New("failed save asset"))

	got := suite.repository.Save(asset)
	assert.Error(suite.T(), got)
	assert.NotNil(suite.T(), got)
}

func (suite *AssetRepositoryTestSuite) TestFindAll_Success() {
	assets := []model.Asset{{
		Id:        "1",
		Category:  model.Category{
			Id:   "1",
			Name: "bergerak",
		},
		AssetType: model.TypeAsset{
			Id:   "2",
			Name: "ringan",
		},
		Name:      "Mobil",
		Available: 50,
		Status:    "ready",
		EntryDate: time.Now(),
		ImgUrl:    "qwerty",
		Total:    50,
	},
	{
		Id:        "2",
		Category:  model.Category{
			Id:   "1",
			Name: "tIdak bergerak",
		},
		AssetType: model.TypeAsset{
			Id:   "2",
			Name: "berat",
		},
		Name:      "papan tulis",
		Available: 2,
		Status:    "ready",
		EntryDate: time.Now(),
		ImgUrl:    "qwerty",
		Total:    2,
	},
	}
	rows := sqlmock.NewRows([]string{"id", "name", "available", "status", "entry_date", "img_url", "total", "id_category", "category_name", "id_asset_type", "asset_type_name"})

	for _, v := range assets {
		rows.AddRow(v.Id, v.Name, v.Available, v.Status, v.EntryDate, v.ImgUrl, v.Total, v.Category.Id, v.Category.Name, v.AssetType.Id, v.AssetType.Name)
	}

	suite.mockSQL.ExpectQuery("select a.id, a.name, a.available, a.status, a.entry_date, a.img_url, a.total, c.id, c.name, at.id, at.name from asset").WillReturnRows(rows)

	got, err := suite.repository.FindAll()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), got, assets)
}

func (suite *AssetRepositoryTestSuite) TestFindAll_Failed() {
	suite.mockSQL.ExpectQuery("select a.id, a.name, a.Total, a.status, a.entry_date, a.img_url, c.id, c.name, at.id, at.name from asset").WillReturnError(errors.New("failed to get asset"))

	got, err := suite.repository.FindAll()
	assert.Nil(suite.T(), got)
	assert.Error(suite.T(), err)
}

func (suite *AssetRepositoryTestSuite) TestFindAll_FailedRows() {
	asset := model.Asset{
		Id:        "1",
		Category:  model.Category{
			Id:   "1",
			Name: "bergerak",
		},
		AssetType: model.TypeAsset{
			Id:   "2",
			Name: "ringan",
		},
		Name:      "Mobil",
		Available: 50,
		Status:    "ready",
		EntryDate: time.Now(),
		ImgUrl:    "qwerty",
		Total:    50,
	}
	rows := sqlmock.NewRows([]string{"id", "name", "available", "status", "entry_date", "img_url", "total", "id_category", "category_name", "id_asset_type", "asset_type_name"}).
	AddRow(asset.Id, asset.Name, asset.Available, asset.Status, asset.EntryDate, asset.ImgUrl, asset.Total, asset.Category.Id, asset.Category.Name, asset.AssetType.Id, asset.AssetType.Name).
	RowError(0, fmt.Errorf("error scan"))

	suite.mockSQL.ExpectQuery("select a.id, a.name, a.available, a.status, a.entry_date, a.img_url, a.total, c.id, c.name, at.id, at.name from asset").WillReturnRows(rows)

	//erros rows.Scan
	got, err := suite.repository.FindAll()
	fmt.Println(got)
	fmt.Println(err)
	assert.Error(suite.T(), err)
	// assert.Nil(suite.T(), got)
	assert.Len(suite.T(), got, 0)
}

func (suite *AssetRepositoryTestSuite) TestFindById_Success() {
	asset := model.Asset{
		Id:        "1",
		Category:  model.Category{
			Id:   "1",
			Name: "bergerak",
		},
		AssetType: model.TypeAsset{
			Id:   "2",
			Name: "ringan",
		},
		Name:      "Mobil",
		Available: 50,
		Status:    "ready",
		EntryDate: time.Now(),
		ImgUrl:    "qwerty",
		Total:    50,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "Available", "status", "entry_date", "img_url", "total", "id_category", "category_name", "id_asset_type", "asset_type_name"}).
	AddRow(asset.Id, asset.Name, asset.Available, asset.Status, asset.EntryDate, asset.ImgUrl, asset.Total, asset.Category.Id, asset.Category.Name, asset.AssetType.Id, asset.AssetType.Name)

	suite.mockSQL.ExpectQuery("select a.id, a.name, a.available, a.status, a.entry_date, a.img_url, a.total, c.id, c.name, at.id, at.name from asset").WillReturnRows(rows)

	got, err := suite.repository.FindById("1")
	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), asset, got)
}

func (suite *AssetRepositoryTestSuite) TestFindById_Failed() {

	suite.mockSQL.ExpectQuery("select a.id, a.name, a.Total, a.status, a.entry_date, a.img_url, c.id, c.name, at.id, at.name from asset").WillReturnError(errors.New("failed get asset with id"))

	got, err := suite.repository.FindById("x")
	assert.Error(suite.T(), err)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), got.Id, "")
}

func (suite *AssetRepositoryTestSuite) TestFindById_FailedRowScan() {
	asset := model.Asset{
		Id:        "1",
		Category:  model.Category{Id: "1", Name: "bergerak"},
		AssetType: model.TypeAsset{Id: "2", Name: "ringan"},
		Name:      "Mobil",
		Total: 50,
		Available:     0,
		Status:    "ready",
		EntryDate: time.Now(),
		ImgUrl:    "qwerty",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "available", "status", "entry_date", "img_url", "total", "id_category", "category_name", "id_asset_type", "asset_type_name"}).
	AddRow(asset.Id, asset.Name, asset.Total, asset.Status, asset.EntryDate, asset.ImgUrl, asset.Total , asset.Category.Id, asset.Category.Name, asset.AssetType.Id, asset.AssetType.Name).RowError(0, errors.New("row scan error"))

	suite.mockSQL.ExpectQuery("select a.id, a.name, a.Total, a.status, a.entry_date, a.img_url, a.total, c.id, c.name, at.id, at.name from asset").WillReturnRows(rows)

	got, err := suite.repository.FindById("x")
	assert.Error(suite.T(), err)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), got.Id, "")
}

func (suite *AssetRepositoryTestSuite) TestUpdate_Success()  {

	asset := model.AssetRequest	{
		Id:         helper.GenerateUUID(),
		CategoryId: "TEST1",
		AssetTypeId:"TEST2",
		Name:       "Laptop",
		Available: 5,
		EntryDate:  time.Now(),
		ImgUrl:     "nothing",
		Total:     5,
	}

	suite.mockSQL.ExpectExec("update asset").WithArgs(
		asset.Id,
		asset.CategoryId, 
		asset.AssetTypeId, 
		asset.Name, 
		asset.Available, 
		asset.Status,
		asset.ImgUrl,
		asset.Total).WillReturnResult(sqlmock.NewResult(1, 1))

	gotError := suite.repository.Update(asset)

	assert.NoError(suite.T(), gotError)
	assert.Nil(suite.T(), gotError)
}

func (suite *AssetRepositoryTestSuite) TestUpdate_Failed()  {
	asset := model.AssetRequest	{
		Id:         helper.GenerateUUID(),
		CategoryId: "TEST1",
		AssetTypeId:"TEST2",
		Name:       "Laptop",
		Total:     5,
		EntryDate:  time.Now(),
		ImgUrl:     "nothing",
	}

	suite.mockSQL.ExpectExec("update asset").WithArgs(
		asset.Id,
		asset.CategoryId, 
		asset.AssetTypeId, 
		asset.Name, 
		asset.Total, 
		asset.Status,
		asset.ImgUrl,).WillReturnError(errors.New("failed to update"))

	gotError := suite.repository.Update(asset)

	assert.Error(suite.T(), gotError)
	assert.NotNil(suite.T(), gotError)
}

func (suite *AssetRepositoryTestSuite) TestDelete_Success()  {
	
	suite.mockSQL.ExpectExec("delete from asset").WithArgs("1").WillReturnResult(sqlmock.NewResult(1, 1))
	gotError := suite.repository.Delete("1")
	assert.NoError(suite.T(), gotError)
	assert.Nil(suite.T(), gotError)
}

func (suite *AssetRepositoryTestSuite) TestDelete_Failed()  {
	
	suite.mockSQL.ExpectExec("delete from asset").WithArgs("1").WillReturnError(errors.New("failed to delete"))
	gotError := suite.repository.Delete("1")
	assert.Error(suite.T(), gotError)
	assert.NotNil(suite.T(), gotError)
}

func (suite *AssetRepositoryTestSuite) TestPaging_Success() {
	mockPaging := dto.PageRequest{
		Page: 1,
		Size: 5,
	}

	asset := model.Asset{
		Id:        "1",
		Category:  model.Category{Id: "1", Name: "bergerak"},
		AssetType: model.TypeAsset{Id: "2", Name: "ringan"},
		Name:      "Mobil",
		Total: 50,
		Available:     0,
		Status:    "ready",
		EntryDate: time.Now(),
		ImgUrl:    "qwerty",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "available", "status", "entry_date", "img_url", "total", "id_category", "category_name", "id_asset_type", "asset_type_name"}).
	AddRow(asset.Id, asset.Name, asset.Total, asset.Status, asset.EntryDate, asset.ImgUrl, asset.Total , asset.Category.Id, asset.Category.Name, asset.AssetType.Id, asset.AssetType.Name)

	suite.mockSQL.ExpectQuery("select a.id, a.name, a.available, a.status, a.entry_date, a.img_url, a.total, c.id, c.name, at.id, at.name from asset as a").
	WithArgs((mockPaging.Page-1)*mockPaging.Size, mockPaging.Size).WillReturnRows(rows)

	rowCount := sqlmock.NewRows([]string{"count"})
	rowCount.AddRow(1)
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta("select count(id) from asset")).WillReturnRows(rowCount)

	actualAsset, actualPaging, actualErr := suite.repository.Paging(mockPaging)
	assert.Nil(suite.T(), actualErr)
	assert.NotNil(suite.T(), actualAsset)
	assert.Equal(suite.T(), 1, actualPaging.TotalRows)
}

func (suite *AssetRepositoryTestSuite) TestPaging_Failed() {
	mockPaging := dto.PageRequest{
		Page: 1,
		Size: 5,
	}

	suite.mockSQL.ExpectQuery("select a.id, a.name, a.available, a.status, a.entry_date, a.img_url, a.total, c.id, c.name, at.id, at.name from asset as a").
	WithArgs((mockPaging.Page-1)*mockPaging.Size, mockPaging.Size).WillReturnError(errors.New("failed get assets"))

	actualAsset, actualPaging, actualErr := suite.repository.Paging(mockPaging)
	assert.Error(suite.T(), actualErr)
	assert.Nil(suite.T(), actualAsset)
	assert.Equal(suite.T(), 0, actualPaging.TotalRows)
}

func (suite *AssetRepositoryTestSuite) TestPaging_RowsError() {
	mockPaging := dto.PageRequest{
		Page: 1,
		Size: 5,
	}

	asset := model.Asset{
		Id:        "1",
		Category:  model.Category{Id: "1", Name: "bergerak"},
		AssetType: model.TypeAsset{Id: "2", Name: "ringan"},
		Name:      "Mobil",
		Total: 50,
		Available:     0,
		Status:    "ready",
		EntryDate: time.Now(),
		ImgUrl:    "qwerty",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "available", "status", "entry_date", "img_url", "total", "id_category", "category_name", "id_asset_type", "asset_type_name"}).
	AddRow(asset.Id, asset.Name, asset.Total, asset.Status, asset.EntryDate, asset.ImgUrl, asset.Total , asset.Category.Id, asset.Category.Name, asset.AssetType.Id, asset.AssetType.Name).
	RowError(0, errors.New("error scan"))

	suite.mockSQL.ExpectQuery("select a.id, a.name, a.available, a.status, a.entry_date, a.img_url, a.total, c.id, c.name, at.id, at.name from asset as a").
	WithArgs((mockPaging.Page-1)*mockPaging.Size, mockPaging.Size).WillReturnRows(rows)

	actualAsset, actualPaging, actualErr := suite.repository.Paging(mockPaging)
	assert.Error(suite.T(), actualErr)
	assert.Nil(suite.T(), actualAsset)
	assert.Equal(suite.T(), 0, actualPaging.TotalRows)
}

func (suite *AssetRepositoryTestSuite) TestPaging_RowCountErr() {
	mockPaging := dto.PageRequest{
		Page: 1,
		Size: 5,
	}

	asset := model.Asset{
		Id:        "1",
		Category:  model.Category{Id: "1", Name: "bergerak"},
		AssetType: model.TypeAsset{Id: "2", Name: "ringan"},
		Name:      "Mobil",
		Total: 50,
		Available:     0,
		Status:    "ready",
		EntryDate: time.Now(),
		ImgUrl:    "qwerty",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "available", "status", "entry_date", "img_url", "total", "id_category", "category_name", "id_asset_type", "asset_type_name"}).
	AddRow(asset.Id, asset.Name, asset.Total, asset.Status, asset.EntryDate, asset.ImgUrl, asset.Total , asset.Category.Id, asset.Category.Name, asset.AssetType.Id, asset.AssetType.Name).
	RowError(0, errors.New("error scan"))

	suite.mockSQL.ExpectQuery("select a.id, a.name, a.available, a.status, a.entry_date, a.img_url, a.total, c.id, c.name, at.id, at.name from asset as a").
	WithArgs((mockPaging.Page-1)*mockPaging.Size, mockPaging.Size).WillReturnRows(rows)

	suite.mockSQL.ExpectQuery(regexp.QuoteMeta("select count(id) from asset")).WillReturnError(errors.New("failed get count"))

	actualAsset, actualPaging, actualErr := suite.repository.Paging(mockPaging)
	assert.Error(suite.T(), actualErr)
	assert.Nil(suite.T(), actualAsset)
	assert.Equal(suite.T(), 0, actualPaging.TotalRows)
}

func (suite *AssetRepositoryTestSuite) TestPaging_RowCountError() {
	mockPaging := dto.PageRequest{
		Page: 1,
		Size: 5,
	}

	asset := model.Asset{
		Id:        "1",
		Category:  model.Category{Id: "1", Name: "bergerak"},
		AssetType: model.TypeAsset{Id: "2", Name: "ringan"},
		Name:      "Mobil",
		Total: 50,
		Available:     0,
		Status:    "ready",
		EntryDate: time.Now(),
		ImgUrl:    "qwerty",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "available", "status", "entry_date", "img_url", "total", "id_category", "category_name", "id_asset_type", "asset_type_name"}).
	AddRow(asset.Id, asset.Name, asset.Total, asset.Status, asset.EntryDate, asset.ImgUrl, asset.Total , asset.Category.Id, asset.Category.Name, asset.AssetType.Id, asset.AssetType.Name)

	suite.mockSQL.ExpectQuery("select a.id, a.name, a.available, a.status, a.entry_date, a.img_url, a.total, c.id, c.name, at.id, at.name from asset as a").
	WithArgs((mockPaging.Page-1)*mockPaging.Size, mockPaging.Size).WillReturnRows(rows)

	rowCount := sqlmock.NewRows([]string{"count"})
	rowCount.AddRow(1).RowError(0, errors.New("failed row count"))
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta("select count(id) from asset")).WillReturnRows(rowCount)

	actualAsset, actualPaging, actualErr := suite.repository.Paging(mockPaging)
	assert.Error(suite.T(), actualErr)
	assert.Nil(suite.T(), actualAsset)
	assert.Equal(suite.T(), 0, actualPaging.TotalRows)
}

func (suite *AssetRepositoryTestSuite) TestUpdateAvailable_Success() {
	
	suite.mockSQL.ExpectExec("update asset set available").WithArgs("1", 5).WillReturnResult(sqlmock.NewResult(1, 1))

	gotErr := suite.repository.UpdateAvailable("1", 5)
	assert.NoError(suite.T(), gotErr)
}

func (suite *AssetRepositoryTestSuite) TestUpdateAvailable_Failed() {
	
	suite.mockSQL.ExpectExec("update asset set available").WithArgs("1", 5).WillReturnError(errors.New("failed update"))

	gotErr := suite.repository.UpdateAvailable("1", 5)
	assert.Error(suite.T(), gotErr)
}

func (suite *AssetRepositoryTestSuite) TestFindByName_Success()  {
	
	assetMock := []model.Asset{{
		Id:        "1",
		Category:  model.Category{Id: "1", Name: "bergerak"},
		AssetType: model.TypeAsset{Id: "2", Name: "ringan"},
		Name:      "Mobil",
		Available:     0,
		Status:    "ready",
		EntryDate: time.Now(),
		ImgUrl:    "qwerty",
		Total: 50,
	}}

	rows := sqlmock.NewRows([]string{"id", "name", "available", "status", "entry_date", "img_url", "total", "id_category", "category_name", "id_asset_type", "asset_type_name"})
	for _, asset := range assetMock {
		rows.AddRow(asset.Id, asset.Name, asset.Available, asset.Status, asset.EntryDate, asset.ImgUrl, asset.Total , asset.Category.Id, asset.Category.Name, asset.AssetType.Id, asset.AssetType.Name)
	}

	suite.mockSQL.ExpectQuery("select a.id, a.name, a.available, a.status, a.entry_date, a.img_url, a.total, c.id, c.name, at.id, at.name from asset as a ").WithArgs("%"+"mobil"+"%").WillReturnRows(rows)

	assets, err := suite.repository.FindByName("mobil")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), assetMock, assets)
}

func (suite *AssetRepositoryTestSuite) TestFindByName_Failed()  {


	suite.mockSQL.ExpectQuery("select a.id, a.name, a.available, a.status, a.entry_date, a.img_url, a.total, c.id, c.name, at.id, at.name from asset as a ").WithArgs("%"+"mobil"+"%").
	WillReturnError(errors.New("failed get asset"))

	assets, err := suite.repository.FindByName("mobil")
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), assets)
}

func (suite *AssetRepositoryTestSuite) TestFindByName_FailedRowScan()  {
	
	assetMock := []model.Asset{{
		Id:        "1",
		Category:  model.Category{Id: "1", Name: "bergerak"},
		AssetType: model.TypeAsset{Id: "2", Name: "ringan"},
		Name:      "Mobil",
		Available:     0,
		Status:    "ready",
		EntryDate: time.Now(),
		ImgUrl:    "qwerty",
		Total: 50,
	}}

	rows := sqlmock.NewRows([]string{"id", "name", "available", "status", "entry_date", "img_url", "total", "id_category", "category_name", "id_asset_type", "asset_type_name"})
	for _, asset := range assetMock {
		rows.AddRow(asset.Id, asset.Name, asset.Available, asset.Status, asset.EntryDate, asset.ImgUrl, asset.Total , asset.Category.Id, asset.Category.Name, asset.AssetType.Id, asset.AssetType.Name).
		RowError(0, errors.New("failed rows"))
	}

	suite.mockSQL.ExpectQuery("select a.id, a.name, a.available, a.status, a.entry_date, a.img_url, a.total, c.id, c.name, at.id, at.name from asset as a ").WithArgs("%"+"mobil"+"%").WillReturnRows(rows)

	assets, err := suite.repository.FindByName("mobil")
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), assets)
}