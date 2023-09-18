package repository

import (
	"database/sql"
	"errors"
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TypeAssetRepositoryTestSuite struct {
	suite.Suite
	repo    TypeAssetRepository
	mockDB  *sql.DB
	mockSQL sqlmock.Sqlmock
}

func (suite *TypeAssetRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDB = db
	suite.mockSQL = mock
	suite.repo = NewTypeAssetRepository(suite.mockDB)
}

func TestTypeAssetRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TypeAssetRepositoryTestSuite))
}

func (suite *TypeAssetRepositoryTestSuite) TestCreate_Success() {
	mockData := model.TypeAsset{
		Id:   "1",
		Name: "Bergerak",
	}
	suite.mockSQL.ExpectExec("INSERT INTO asset_type").WithArgs(mockData.Id, mockData.Name).WillReturnResult(sqlmock.NewResult(1, 1))
	err := suite.repo.Save(mockData)
	assert.NoError(suite.T(), err)
}

func (suite *TypeAssetRepositoryTestSuite) TestCreate_Failed() {
	mockData := model.TypeAsset{
		Id:   "1",
		Name: "Bergerak",
	}
	suite.mockSQL.ExpectExec("INSERT INTO asset_type").WithArgs(mockData.Id, mockData.Name).WillReturnError(errors.New("failed save category"))
	err := suite.repo.Save(mockData)
	assert.Error(suite.T(), err)
}

func (suite *TypeAssetRepositoryTestSuite) TestFindAll_Success() {
	expectedAssets := []model.TypeAsset{
		{
			Id:   "1",
			Name: "Bergerak",
		},
		{
			Id:   "2",
			Name: "Tidak Bergerak",
		},
	}

	// Membuat rows mock dengan kolom yang sesuai
	rows := sqlmock.NewRows([]string{"id", "name"})
	for _, asset := range expectedAssets {
		rows.AddRow(asset.Id, asset.Name)
	}

	// Mengharapkan query SELECT * FROM asset_type dan mengembalikan rows mock
	suite.mockSQL.ExpectQuery("SELECT id, name FROM asset_type").WillReturnRows(rows)

	// Menjalankan fungsi yang diuji
	got, err := suite.repo.FindAll()
	assert.NoError(suite.T(), err)

	// Membandingkan hasil yang diharapkan dengan hasil yang diperoleh
	assert.Equal(suite.T(), expectedAssets, got)
}

func (suite *TypeAssetRepositoryTestSuite) TestFindAll_FailedRows() {
	assets := []model.TypeAsset{
		{
			Id:   "1",
			Name: "Bergerak",
		},
		{
			Id:   "2",
			Name: "dsds",
		},
	}

	// Membuat rows mock dengan kolom yang sesuai
	rows := sqlmock.NewRows([]string{"id", "name"})
	for _, asset := range assets {
		rows.AddRow(asset.Id, asset.Name)
	}

	// Menambahkan row yang akan menghasilkan error
	rows.RowError(0, errors.New("error new scan"))

	// Mengharapkan query SELECT id, name FROM asset_type dan mengembalikan rows mock
	suite.mockSQL.ExpectQuery("SELECT id, name FROM asset_type").WillReturnRows(rows)

	// Menjalankan fungsi yang diuji
	got, err := suite.repo.FindAll()

	// Memeriksa bahwa error dihasilkan
	assert.Error(suite.T(), err)
	assert.NotEqual(suite.T(), assets, got)
}

func (suite *TypeAssetRepositoryTestSuite) TestFindAll_Failed() {

	suite.mockSQL.ExpectQuery("SELECT id, name FROM asset_type").WillReturnError(errors.New("failed get categories"))
	got, err := suite.repo.FindAll()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), got)
}

func (suite *TypeAssetRepositoryTestSuite) TestFindById_Success() {
	assets := model.TypeAsset{
		Id:   "1",
		Name: "Bergerak",
	}
	row := sqlmock.NewRows([]string{"id", "name"}).AddRow(assets.Id, assets.Name)
	suite.mockSQL.ExpectQuery("SELECT id,name FROM asset_type WHERE id").WithArgs("1").WillReturnRows(row)
	got, err := suite.repo.FindById("1")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), assets, got)
}

func (suite *TypeAssetRepositoryTestSuite) TestFindById_Failed() {
	suite.mockSQL.ExpectQuery("SELECT id, name FROM asset_type").WithArgs("1").WillReturnError(errors.New("failed get type asset"))
	got, err := suite.repo.FindById("1")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.TypeAsset{}, got)
}

func (suite *TypeAssetRepositoryTestSuite) TestUpdate_Success() {
	mockData := model.TypeAsset{
		Id:   "1",
		Name: "Bergerak",
	}
	suite.mockSQL.ExpectExec("UPDATE asset_type SET").WithArgs(mockData.Id, mockData.Name).WillReturnResult(sqlmock.NewResult(1, 1))
	err := suite.repo.Update(mockData)
	assert.NoError(suite.T(), err)
}

func (suite *TypeAssetRepositoryTestSuite) TestUpdate_Failed() {
	mockData := model.TypeAsset{
		Id:   "1",
		Name: "Bergerak",
	}
	suite.mockSQL.ExpectExec("UPDATE asset_type SET").WithArgs(mockData.Id, mockData.Name).WillReturnError(errors.New("failed update type asset"))
	err := suite.repo.Update(mockData)
	assert.Error(suite.T(), err)
}

func (suite *TypeAssetRepositoryTestSuite) TestDelete_Success() {
	suite.mockSQL.ExpectExec("DELETE FROM asset_type").WithArgs("1").WillReturnResult(sqlmock.NewResult(1, 1))
	gotErr := suite.repo.Delete("1")
	assert.NoError(suite.T(), gotErr)
}

func (suite *TypeAssetRepositoryTestSuite) TestDelete_Failed() {
	suite.mockSQL.ExpectExec("DELETE FROM asset_type").WithArgs("1").WillReturnError(errors.New("failed delete type asset"))
	gotErr := suite.repo.Delete("1")
	assert.Error(suite.T(), gotErr)
}

func (suite *TypeAssetRepositoryTestSuite) TestFindByName_Success() {
	expectedAssets := []model.TypeAsset{
		{
			Id:   "1",
			Name: "Bergerak",
		},
		{
			Id:   "2",
			Name: "Tidak Bergerak",
		},
	}

	// Membuat rows mock dengan kolom yang sesuai
	rows := sqlmock.NewRows([]string{"id", "name"})
	for _, asset := range expectedAssets {
		rows.AddRow(asset.Id, asset.Name)
	}

	// Mengharapkan query SELECT * FROM asset_type dan mengembalikan rows mock
	suite.mockSQL.ExpectQuery("SELECT id, name FROM asset_type WHERE name ILIKE").WillReturnRows(rows)

	// Menjalankan fungsi yang diuji
	got, err := suite.repo.FindByName("Bergerak")
	assert.NoError(suite.T(), err)

	// Membandingkan hasil yang diharapkan dengan hasil yang diperoleh
	assert.Equal(suite.T(), expectedAssets, got)
}

func (suite *TypeAssetRepositoryTestSuite) TestFindByName_FailedRows() {
	assets := []model.TypeAsset{
		{
			Id:   "1",
			Name: "Bergerak",
		},
		{
			Id:   "2",
			Name: "dsds",
		},
	}

	// Membuat rows mock dengan kolom yang sesuai
	rows := sqlmock.NewRows([]string{"id", "name"})
	for _, asset := range assets {
		rows.AddRow(asset.Id, asset.Name)
	}

	// Menambahkan row yang akan menghasilkan error
	rows.RowError(0, errors.New("error new scan"))

	// Mengharapkan query SELECT id, name FROM asset_type dan mengembalikan rows mock
	suite.mockSQL.ExpectQuery("SELECT id, name FROM asset_type WHERE name ILIKE").WillReturnRows(rows)

	// Menjalankan fungsi yang diuji
	got, err := suite.repo.FindByName("Bergerak")

	// Memeriksa bahwa error dihasilkan
	assert.Error(suite.T(), err)
	assert.NotEqual(suite.T(), assets, got)
}

func (suite *TypeAssetRepositoryTestSuite) TestFindByName_Failed() {

	suite.mockSQL.ExpectQuery("SELECT id, name FROM asset_type WHERE name ILIKE").WillReturnError(errors.New("failed get type asset"))
	got, err := suite.repo.FindByName("Bergerak")
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), got)
}

func (suite *TypeAssetRepositoryTestSuite) TestPaging_Success() {
	mockPageRequest := dto.PageRequest{
		Page: 1,
		Size: 5,
	}
	mockData := []model.TypeAsset{
		{
			Id:   "1",
			Name: "bergerak",
		},
		{
			Id:   "2",
			Name: "tidak bergerak",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name"})
	for _, v := range mockData {
		rows.AddRow(v.Id, v.Name)
	}
	expectedQuery := `SELECT id, name FROM asset_type LIMIT $2 OFFSET $1`
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs(
		(mockPageRequest.Page-1)*mockPageRequest.Size,
		mockPageRequest.Size,
	).WillReturnRows(rows)

	rowCount := sqlmock.NewRows([]string{"count"})
	rowCount.AddRow(2)
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(id) FROM asset_type`)).
		WillReturnRows(rowCount)

	actualTypeAsset, actualPaging, actualErr := suite.repo.Paging(mockPageRequest)
	assert.Nil(suite.T(), actualErr)
	assert.NotNil(suite.T(), actualTypeAsset)
	assert.Equal(suite.T(), 2, actualPaging.TotalRows)
}

func (suite *TypeAssetRepositoryTestSuite) TestPaging_Fail() {
	mockPageRequest := dto.PageRequest{
		Page: 1,
		Size: 5,
	}
	mockData := []model.TypeAsset{
		{
			Id:   "1",
			Name: "berg",
		},
		{
			Id:   "2",
			Name: "tidak b",
		},
	}

	//err select paging
	expectedQuery := `SELECT id, name FROM asset_type LIMIT $2 OFFSET $1`
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WillReturnError(errors.New("failed"))
	actualTypeAsset, actualPaging, actualErr := suite.repo.Paging(dto.PageRequest{})
	assert.Error(suite.T(), actualErr)
	assert.Nil(suite.T(), actualTypeAsset)
	assert.Equal(suite.T(), 0, actualPaging.TotalRows)

	// Konfigurasi untuk mengharapkan panggilan ke rows.Scan dengan kesalahan
	expectedQuery = `SELECT id, name FROM asset_type LIMIT $2 OFFSET $1`
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name"}).AddRow("invalid", "data"),
	).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name"}).AddRow(nil, "Pcs"),
	)
	actualTypeAsset, actualPaging, actualErr = suite.repo.Paging(mockPageRequest)
	assert.Error(suite.T(), actualErr)
	assert.Nil(suite.T(), actualTypeAsset)
	assert.Equal(suite.T(), 0, actualPaging.TotalRows)

	//err select count
	rows := sqlmock.NewRows([]string{"id", "name"})
	for _, v := range mockData {
		rows.AddRow(v.Id, v.Name)
	}
	expectedQuery = `SELECT id, name FROM asset_type LIMIT $2 OFFSET $1`
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs(
		(mockPageRequest.Page-1)*mockPageRequest.Size,
		mockPageRequest.Size).WillReturnRows(rows)
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(id) FROM asset_type`)).WillReturnError(errors.New("failed"))

	actualTypeAsset, actualPaging, actualErr = suite.repo.Paging(mockPageRequest)
	assert.Error(suite.T(), actualErr)
	assert.Nil(suite.T(), actualTypeAsset)
	assert.Equal(suite.T(), 0, actualPaging.TotalRows)
}
