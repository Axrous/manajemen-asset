package repository

import (
	"database/sql"
	"errors"
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/util/helper"
	"fmt"
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
		ID:         helper.GenerateUUID(),
		CategoryId: "TEST1",
		AssetTypeId:  "TEST2",
		Name:       "Laptop",
		Amount:     "5",
		EntryDate:  time.Now(),
		ImgUrl:     "nothing",
	}


	suite.mockSQL.ExpectExec("insert into asset").WithArgs(
		asset.ID, 
		asset.CategoryId, 
		asset.AssetTypeId, 
		asset.Name, 
		asset.Amount, 
		asset.Status, 
		asset.EntryDate, 
		asset.ImgUrl,).WillReturnResult(sqlmock.NewResult(1, 1))

	got := suite.repository.Save(asset)
	assert.NoError(suite.T(), got)
	assert.Nil(suite.T(), got)
}

func (suite *AssetRepositoryTestSuite) TestCreate_Failed() {
	
	asset := model.AssetRequest{
		ID:         helper.GenerateUUID(),
		CategoryId: "TEST1",
		AssetTypeId:"TEST2",
		Name:       "Laptop",
		Amount:     "5",
		EntryDate:  time.Now(),
		ImgUrl:     "nothing",
	}


	suite.mockSQL.ExpectExec("insert into asset").WithArgs(
		asset.ID, 
		asset.CategoryId, 
		asset.AssetTypeId, 
		asset.Name, 
		asset.Amount, 
		asset.Status, 
		asset.EntryDate, 
		asset.ImgUrl,).WillReturnError(errors.New("failed save asset"))

	got := suite.repository.Save(asset)
	assert.Error(suite.T(), got)
	assert.NotNil(suite.T(), got)
}

func (suite *AssetRepositoryTestSuite) TestFindAll_Success() {
	assets := []model.Asset{{
		ID:        "1",
		Category:  model.Category{
			ID:   "1",
			Name: "bergerak",
		},
		AssetType: model.AssetType{
			ID:   "2",
			Name: "ringan",
		},
		Name:      "Mobil",
		Amount:    50,
		Status:    "ready",
		EntryDate: time.Now(),
		ImgUrl:    "qwerty",
	},
	{
		ID:        "2",
		Category:  model.Category{
			ID:   "1",
			Name: "tidak bergerak",
		},
		AssetType: model.AssetType{
			ID:   "2",
			Name: "berat",
		},
		Name:      "papan tulis",
		Amount:    2,
		Status:    "ready",
		EntryDate: time.Now(),
		ImgUrl:    "qwerty",
	},
	}
	rows := sqlmock.NewRows([]string{"id", "name", "amount", "status", "entry_date", "img_url", "id_category", "category_name", "id_asset_type", "asset_type_name"})

	for _, v := range assets {
		rows.AddRow(v.ID, v.Name, v.Amount, v.Status, v.EntryDate, v.ImgUrl, v.Category.ID, v.Category.Name, v.AssetType.ID, v.AssetType.Name)
	}

	suite.mockSQL.ExpectQuery("select a.id, a.name, a.amount, a.status, a.entry_date, a.img_url, c.id, c.name, at.id, at.name from asset").WillReturnRows(rows)

	got, err := suite.repository.FindAll()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), got, assets)
}

func (suite *AssetRepositoryTestSuite) TestFindAll_Failed() {
	suite.mockSQL.ExpectQuery("select a.id, a.name, a.amount, a.status, a.entry_date, a.img_url, c.id, c.name, at.id, at.name from asset").WillReturnError(errors.New("failed to get asset"))

	got, err := suite.repository.FindAll()
	assert.Nil(suite.T(), got)
	assert.Error(suite.T(), err)
}

func (suite *AssetRepositoryTestSuite) TestFindAll_FailedRows() {
	asset := model.Asset{
		ID:        "1",
		Category:  model.Category{
			ID:   "1",
			Name: "bergerak",
		},
		AssetType: model.AssetType{
			ID:   "2",
			Name: "ringan",
		},
		Name:      "Mobil",
		Amount:    50,
		Status:    "ready",
		EntryDate: time.Now(),
		ImgUrl:    "qwerty",
	}
	rows := sqlmock.NewRows([]string{"id", "name", "amount", "status", "entry_date", "img_url", "id_category", "category_name", "id_asset_type", "asset_type_name"}).
	AddRow(asset.ID, asset.Name, asset.Amount, asset.Status, asset.EntryDate, asset.ImgUrl, asset.Category.ID, asset.Category.Name, asset.AssetType.ID, asset.AssetType.Name).
	RowError(0, fmt.Errorf("error scan"))

	suite.mockSQL.ExpectQuery("select a.id, a.name, a.amount, a.status, a.entry_date, a.img_url, c.id, c.name, at.id, at.name from asset").WillReturnRows(rows)

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
		ID:        "1",
		Category:  model.Category{
			ID:   "1",
			Name: "bergerak",
		},
		AssetType: model.AssetType{
			ID:   "2",
			Name: "ringan",
		},
		Name:      "Mobil",
		Amount:    50,
		Status:    "ready",
		EntryDate: time.Now(),
		ImgUrl:    "qwerty",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "amount", "status", "entry_date", "img_url", "id_category", "category_name", "id_asset_type", "asset_type_name"}).
	AddRow(asset.ID, asset.Name, asset.Amount, asset.Status, asset.EntryDate, asset.ImgUrl, asset.Category.ID, asset.Category.Name, asset.AssetType.ID, asset.AssetType.Name)

	suite.mockSQL.ExpectQuery("select a.id, a.name, a.amount, a.status, a.entry_date, a.img_url, c.id, c.name, at.id, at.name from asset").WillReturnRows(rows)

	got, err := suite.repository.FindById("1")
	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), asset, got)
}

func (suite *AssetRepositoryTestSuite) TestFindById_Failed() {

	suite.mockSQL.ExpectQuery("select a.id, a.name, a.amount, a.status, a.entry_date, a.img_url, c.id, c.name, at.id, at.name from asset").WillReturnError(errors.New("failed get asset with id"))

	got, err := suite.repository.FindById("x")
	assert.Error(suite.T(), err)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), got.ID, "")
}

func (suite *AssetRepositoryTestSuite) TestFindById_FailedRowScan() {
	asset := model.Asset{
		ID:        "1",
		Category:  model.Category{
			ID:   "1",
			Name: "bergerak",
		},
		AssetType: model.AssetType{
			ID:   "2",
			Name: "ringan",
		},
		Name:      "Mobil",
		Amount:    50,
		Status:    "ready",
		EntryDate: time.Now(),
		ImgUrl:    "qwerty",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "amount", "status", "entry_date", "img_url", "id_category", "category_name", "id_asset_type", "asset_type_name"}).
	AddRow(asset.ID, asset.Name, asset.Amount, asset.Status, asset.EntryDate, asset.ImgUrl, asset.Category.ID, asset.Category.Name, asset.AssetType.ID, asset.AssetType.Name).RowError(0, errors.New("row scan error"))

	suite.mockSQL.ExpectQuery("select a.id, a.name, a.amount, a.status, a.entry_date, a.img_url, c.id, c.name, at.id, at.name from asset").WillReturnRows(rows)

	got, err := suite.repository.FindById("x")
	assert.Error(suite.T(), err)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), got.ID, "")
}

func (suite *AssetRepositoryTestSuite) TestUpdate_Success()  {

	asset := model.AssetRequest	{
		ID:         helper.GenerateUUID(),
		CategoryId: "TEST1",
		AssetTypeId:"TEST2",
		Name:       "Laptop",
		Amount:     "5",
		EntryDate:  time.Now(),
		ImgUrl:     "nothing",
	}

	suite.mockSQL.ExpectExec("update asset").WithArgs(
		asset.ID,
		asset.CategoryId, 
		asset.AssetTypeId, 
		asset.Name, 
		asset.Amount, 
		asset.Status,
		asset.ImgUrl,).WillReturnResult(sqlmock.NewResult(1, 1))

	gotError := suite.repository.Update(asset)

	assert.NoError(suite.T(), gotError)
	assert.Nil(suite.T(), gotError)
}

func (suite *AssetRepositoryTestSuite) TestUpdate_Failed()  {
	asset := model.AssetRequest	{
		ID:         helper.GenerateUUID(),
		CategoryId: "TEST1",
		AssetTypeId:"TEST2",
		Name:       "Laptop",
		Amount:     "5",
		EntryDate:  time.Now(),
		ImgUrl:     "nothing",
	}

	suite.mockSQL.ExpectExec("update asset").WithArgs(
		asset.ID,
		asset.CategoryId, 
		asset.AssetTypeId, 
		asset.Name, 
		asset.Amount, 
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

