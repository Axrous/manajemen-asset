package repository

import (
	"database/sql"
	"errors"
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StaffRepositoryTestSuite struct {
	suite.Suite
	repo    StaffRepository
	mockDB  *sql.DB
	mockSQL sqlmock.Sqlmock
}

func (suite *StaffRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDB = db
	suite.mockSQL = mock
	suite.repo = NewStaffRepository(suite.mockDB)
}

func TestStaffRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(StaffRepositoryTestSuite))
}

func (suite *StaffRepositoryTestSuite) TestCreate_Success() {
	mockData := model.Staff{
		Nik_Staff:    "1",
		Name:         "Bergerak",
		Phone_number: "212211",
		Address:      "pku",
		Birth_date:   time.Time{},
		Img_url:      "jjj.png",
		Divisi:       "IT",
	}
	suite.mockSQL.ExpectExec("INSERT INTO staff").WithArgs(mockData.Nik_Staff, mockData.Name, mockData.Phone_number, mockData.Address, mockData.Birth_date, mockData.Img_url, mockData.Divisi).WillReturnResult(sqlmock.NewResult(1, 1))
	err := suite.repo.Save(mockData)
	assert.NoError(suite.T(), err)
}

func (suite *StaffRepositoryTestSuite) TestCreate_Failed() {
	mockData := model.Staff{
		Nik_Staff:    "1",
		Name:         "Bergerak",
		Phone_number: "0822",
		Address:      "pku",
		Birth_date:   time.Time{},
		Img_url:      "sss.png",
		Divisi:       "IT",
	}
	suite.mockSQL.ExpectExec("INSERT INTO staff").WithArgs(mockData.Nik_Staff, mockData.Name, mockData.Phone_number, mockData.Address, mockData.Birth_date, mockData.Img_url, mockData.Divisi).WillReturnError(errors.New("failed save Staff"))
	err := suite.repo.Save(mockData)
	assert.Error(suite.T(), err)
}

func (suite *StaffRepositoryTestSuite) TestFindAll_Success() {
	expectedAssets := []model.Staff{
		{
			Nik_Staff:    "1",
			Name:         "Bergerak",
			Phone_number: "0822",
			Address:      "pku",
			Birth_date:   time.Time{},
			Img_url:      "sss.png",
			Divisi:       "IT",
		},
		{
			Nik_Staff:    "2",
			Name:         "Bergerak",
			Phone_number: "0822222",
			Address:      "pku",
			Birth_date:   time.Time{},
			Img_url:      "sss.png",
			Divisi:       "IT",
		},
	}

	// Membuat rows mock dengan kolom yang sesuai
	rows := sqlmock.NewRows([]string{"nik_staff", "name", "phone_number", "address", "birth_date", "img_url", "divisi"})
	for _, asset := range expectedAssets {
		rows.AddRow(asset.Nik_Staff, asset.Name, asset.Phone_number, asset.Address, asset.Birth_date, asset.Img_url, asset.Divisi)
	}

	// Mengharapkan query SELECT * FROM asset_type dan mengembalikan rows mock
	suite.mockSQL.ExpectQuery("SELECT nik_staff, name, phone_number, address, birth_date, img_url, divisi FROM staff").WillReturnRows(rows)

	// Menjalankan fungsi yang diuji
	got, err := suite.repo.FindByAll()
	assert.NoError(suite.T(), err)

	// Membandingkan hasil yang diharapkan dengan hasil yang diperoleh
	assert.Equal(suite.T(), expectedAssets, got)
}

func (suite *StaffRepositoryTestSuite) TestFindAll_FailedRows() {
	assets := []model.Staff{
		{
			Nik_Staff:    "1",
			Name:         "Bergerak",
			Phone_number: "08222",
			Address:      "pku",
			Birth_date:   time.Time{},
			Img_url:      "hhh.jpg",
			Divisi:       "IT",
		},
		{
			Nik_Staff:    "2",
			Name:         "dsds",
			Phone_number: "08766",
			Address:      "jkt",
			Birth_date:   time.Time{},
			Img_url:      "ss.jpg",
			Divisi:       "IT",
		},
	}

	// Membuat rows mock dengan kolom yang sesuai
	rows := sqlmock.NewRows([]string{"id", "name", "phone_number", "address", "birth_date", "img_url", "divisi"})
	for _, asset := range assets {
		rows.AddRow(asset.Nik_Staff, asset.Name, asset.Phone_number, asset.Address, asset.Birth_date, asset.Img_url, asset.Divisi)
	}

	// Menambahkan row yang akan menghasilkan error
	rows.RowError(0, errors.New("error new scan"))

	// Mengharapkan query SELECT id, name FROM asset_type dan mengembalikan rows mock
	suite.mockSQL.ExpectQuery("SELECT nik_staff, name, phone_number, address, birth_date, img_url, divisi FROM staff").WillReturnRows(rows)

	// Menjalankan fungsi yang diuji
	got, err := suite.repo.FindByAll()

	// Memeriksa bahwa error dihasilkan
	assert.Error(suite.T(), err)
	assert.NotEqual(suite.T(), assets, got)
}

func (suite *StaffRepositoryTestSuite) TestFindAll_Failed() {

	suite.mockSQL.ExpectQuery("SELECT nik_staff, name, phone_number, address, birth_date, img_url, divisi FROM staff").WillReturnError(errors.New("failed get staff"))
	got, err := suite.repo.FindByAll()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), got)
}

func (suite *StaffRepositoryTestSuite) TestFindById_Success() {
	assets := model.Staff{
		Nik_Staff:    "1",
		Name:         "Bergerak",
		Phone_number: "08222",
		Address:      "pku",
		Birth_date:   time.Time{},
		Img_url:      "ssd.jpg",
		Divisi:       "IT",
	}
	row := sqlmock.NewRows([]string{"nik_staff", "name", "phone_number", "address", "birth_date", "img_url", "divisi"}).AddRow(assets.Nik_Staff, assets.Name, assets.Phone_number, assets.Address, assets.Birth_date, assets.Img_url, assets.Divisi)
	suite.mockSQL.ExpectQuery("SELECT nik_staff, name, phone_number, address, birth_date, img_url, divisi FROM staff WHERE nik_staff").WithArgs("1").WillReturnRows(row)
	got, err := suite.repo.FindById("1")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), assets, got)
}

func (suite *StaffRepositoryTestSuite) TestFindById_Failed() {
	suite.mockSQL.ExpectQuery("SELECT * FROM staff").WithArgs("1").WillReturnError(errors.New("failed get staff"))
	got, err := suite.repo.FindById("1")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), model.Staff{}, got)
}

func (suite *StaffRepositoryTestSuite) TestUpdate_Success() {
	mockData := model.Staff{
		Nik_Staff:    "1",
		Name:         "Bergerak",
		Phone_number: "0822",
		Address:      "pku",
		Birth_date:   time.Time{},
		Img_url:      "jhj.jpg",
		Divisi:       "IT",
	}
	suite.mockSQL.ExpectExec("UPDATE staff SET").WithArgs(mockData.Nik_Staff, mockData.Name, mockData.Phone_number, mockData.Address, mockData.Birth_date, mockData.Img_url, mockData.Divisi).WillReturnResult(sqlmock.NewResult(1, 1))
	err := suite.repo.Update(mockData)
	assert.NoError(suite.T(), err)
}

func (suite *StaffRepositoryTestSuite) TestUpdate_Failed() {
	mockData := model.Staff{
		Nik_Staff:    "1",
		Name:         "Bergerak",
		Phone_number: "0822",
		Address:      "pku",
		Birth_date:   time.Time{},
		Img_url:      "jhhg.jpg",
		Divisi:       "IT",
	}
	suite.mockSQL.ExpectExec("UPDATE staff SET").WithArgs(mockData.Nik_Staff, mockData.Name, mockData.Phone_number, mockData.Address, mockData.Birth_date, mockData.Img_url, mockData.Divisi).WillReturnError(errors.New("failed update staff"))
	err := suite.repo.Update(mockData)
	assert.Error(suite.T(), err)
}

func (suite *StaffRepositoryTestSuite) TestDelete_Success() {
	suite.mockSQL.ExpectExec("DELETE FROM staff").WithArgs("1").WillReturnResult(sqlmock.NewResult(1, 1))
	gotErr := suite.repo.Delete("1")
	assert.NoError(suite.T(), gotErr)
}

func (suite *StaffRepositoryTestSuite) TestDelete_Failed() {
	suite.mockSQL.ExpectExec("DELETE FROM staff").WithArgs("1").WillReturnError(errors.New("failed delete staff"))
	gotErr := suite.repo.Delete("1")
	assert.Error(suite.T(), gotErr)
}

func (suite *StaffRepositoryTestSuite) TestFindByName_Success() {
	expectedAssets := []model.Staff{
		{
			Nik_Staff:    "1",
			Name:         "Bergerak",
			Phone_number: "08223",
			Address:      "pku",
			Birth_date:   time.Time{},
			Img_url:      "pki.jpg",
			Divisi:       "IT",
		},
		{
			Nik_Staff:    "2",
			Name:         "Tidak Bergerak",
			Phone_number: "0812",
			Address:      "pku",
			Birth_date:   time.Time{},
			Img_url:      "pik.jpg",
			Divisi:       "IT",
		},
	}

	// Membuat rows mock dengan kolom yang sesuai
	rows := sqlmock.NewRows([]string{"nik_staff", "name", "phone_number", "address", "birth_date", "img_url", "divisi"})
	for _, asset := range expectedAssets {
		rows.AddRow(asset.Nik_Staff, asset.Name, asset.Phone_number, asset.Address, asset.Birth_date, asset.Img_url, asset.Divisi)
	}

	// Mengharapkan query SELECT * FROM asset_type dan mengembalikan rows mock
	suite.mockSQL.ExpectQuery("SELECT nik_staff, name, phone_number, address, birth_date, img_url, divisi FROM staff WHERE name ILIKE").WillReturnRows(rows)

	// Menjalankan fungsi yang diuji
	got, err := suite.repo.FindByName("Bergerak")
	assert.NoError(suite.T(), err)

	// Membandingkan hasil yang diharapkan dengan hasil yang diperoleh
	assert.Equal(suite.T(), expectedAssets, got)
}

func (suite *StaffRepositoryTestSuite) TestFindByName_FailedRows() {
	assets := []model.Staff{
		{
			Nik_Staff:    "1",
			Name:         "Bergerak",
			Phone_number: "0822",
			Address:      "pku",
			Birth_date:   time.Time{},
			Img_url:      "ijj.jpg",
			Divisi:       "IT",
		},
		{
			Nik_Staff:    "2",
			Name:         "dsds",
			Phone_number: "0812",
			Address:      "pku",
			Birth_date:   time.Time{},
			Img_url:      "oku.jpg",
			Divisi:       "IT",
		},
	}

	// Membuat rows mock dengan kolom yang sesuai
	rows := sqlmock.NewRows([]string{"nik_staff", "name", "phone_number", "address", "birth_date", "img_url", "divisi"})
	for _, asset := range assets {
		rows.AddRow(asset.Nik_Staff, asset.Name, asset.Phone_number, asset.Address, asset.Birth_date, asset.Img_url, asset.Divisi)
	}

	// Menambahkan row yang akan menghasilkan error
	rows.RowError(0, errors.New("error new scan"))

	// Mengharapkan query SELECT id, name FROM asset_type dan mengembalikan rows mock
	suite.mockSQL.ExpectQuery("SELECT nik_staff, name, phone_number, address, birth_date, img_url, divisi FROM staff WHERE name ILIKE").WillReturnRows(rows)

	// Menjalankan fungsi yang diuji
	got, err := suite.repo.FindByName("Bergerak")

	// Memeriksa bahwa error dihasilkan
	assert.Error(suite.T(), err)
	assert.NotEqual(suite.T(), assets, got)
}

func (suite *StaffRepositoryTestSuite) TestFindByName_Failed() {

	suite.mockSQL.ExpectQuery("SELECT nik_staff, name, phone_number, address, birth_date, img_url, divisi FROM staff WHERE name ILIKE").WillReturnError(errors.New("failed get staff"))
	got, err := suite.repo.FindByName("Bergerak")
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), got)
}

func (suite *StaffRepositoryTestSuite) TestPaging_Success() {
	mockPageRequest := dto.PageRequest{
		Page: 1,
		Size: 5,
	}
	mockData := []model.Staff{
		{
			Nik_Staff:    "1",
			Name:         "bergerak",
			Phone_number: "0822",
			Address:      "pku",
			Birth_date:   time.Time{},
			Img_url:      "pku.jpg",
			Divisi:       "IT",
		},
		{
			Nik_Staff:    "2",
			Name:         "tidak bergerak",
			Phone_number: "0812",
			Address:      "jkt",
			Birth_date:   time.Time{},
			Img_url:      "iku.jpg",
			Divisi:       "IT",
		},
	}

	rows := sqlmock.NewRows([]string{"nik_staff", "name", "phone_number", "address", "birth_date", "img_url", "divisi"})
	for _, v := range mockData {
		rows.AddRow(v.Nik_Staff, v.Name, v.Phone_number, v.Address, v.Birth_date, v.Img_url, v.Divisi)
	}
	expectedQuery := `SELECT nik_staff, name, phone_number, address, birth_date, img_url, divisi FROM staff LIMIT $2 OFFSET $1`
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs(
		(mockPageRequest.Page-1)*mockPageRequest.Size,
		mockPageRequest.Size,
	).WillReturnRows(rows)

	rowCount := sqlmock.NewRows([]string{"count"})
	rowCount.AddRow(2)
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(nik_staff) FROM staff`)).
		WillReturnRows(rowCount)

	actualTypeAsset, actualPaging, actualErr := suite.repo.Paging(mockPageRequest)
	assert.Nil(suite.T(), actualErr)
	assert.NotNil(suite.T(), actualTypeAsset)
	assert.Equal(suite.T(), 2, actualPaging.TotalRows)
}

func (suite *StaffRepositoryTestSuite) TestPaging_Fail() {
	mockPageRequest := dto.PageRequest{
		Page: 1,
		Size: 5,
	}
	mockData := []model.Staff{
		{
			Nik_Staff:    "1",
			Name:         "berg",
			Phone_number: "0822",
			Address:      "pku",
			Birth_date:   time.Time{},
			Img_url:      "pku.jpg",
			Divisi:       "IT",
		},
		{
			Nik_Staff:    "2",
			Name:         "tidak b",
			Phone_number: "0812",
			Address:      "jkt",
			Birth_date:   time.Time{},
			Img_url:      "jkt.png",
			Divisi:       "IT",
		},
	}

	//err select paging
	expectedQuery := `SELECT nik_staff, name, phone_number, address, birth_date, img_url, divisi FROM staff LIMIT $2 OFFSET $1`
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WillReturnError(errors.New("failed"))
	actualTypeAsset, actualPaging, actualErr := suite.repo.Paging(dto.PageRequest{})
	assert.Error(suite.T(), actualErr)
	assert.Nil(suite.T(), actualTypeAsset)
	assert.Equal(suite.T(), 0, actualPaging.TotalRows)

	// Konfigurasi untuk mengharapkan panggilan ke rows.Scan dengan kesalahan
	expectedQuery = `SELECT nik_staff, name, phone_number, address, birth_date, img_url, divisi FROM staff LIMIT $2 OFFSET $1`
	// data sql yg apa aja, jangan semuanya
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WillReturnRows(
		sqlmock.NewRows([]string{"nik_staff", "name"}).AddRow("invalid", "data"),
	).WillReturnRows(
		sqlmock.NewRows([]string{"nik_staff", "name"}).AddRow(nil, "Pcs"),
	)
	actualTypeAsset, actualPaging, actualErr = suite.repo.Paging(mockPageRequest)
	assert.Error(suite.T(), actualErr)
	assert.Nil(suite.T(), actualTypeAsset)
	assert.Equal(suite.T(), 0, actualPaging.TotalRows)

	//err select count
	rows := sqlmock.NewRows([]string{"nik_staff", "name", "phone_number", "address", "birth_date", "img_url", "divisi"})
	for _, v := range mockData {
		rows.AddRow(v.Nik_Staff, v.Name, v.Phone_number, v.Address, v.Birth_date, v.Img_url, v.Divisi)
	}
	expectedQuery = `SELECT nik_staff, name, phone_number, address, birth_date, img_url, divisi FROM staff LIMIT $2 OFFSET $1`
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).WithArgs(
		(mockPageRequest.Page-1)*mockPageRequest.Size,
		mockPageRequest.Size).WillReturnRows(rows)
	suite.mockSQL.ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(nik_staff) FROM staff`)).WillReturnError(errors.New("failed"))

	actualTypeAsset, actualPaging, actualErr = suite.repo.Paging(mockPageRequest)
	assert.Error(suite.T(), actualErr)
	assert.Nil(suite.T(), actualTypeAsset)
	assert.Equal(suite.T(), 0, actualPaging.TotalRows)
}
