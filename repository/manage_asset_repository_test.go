package repository

import (
	"database/sql"
	"errors"
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ManageAssetRepoTestSuite struct {
	suite.Suite
	mockDB *sql.DB
	mockSQL sqlmock.Sqlmock
	repo ManageAssetRepository
}

func (suite *ManageAssetRepoTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDB = db
	suite.mockSQL = mock
	suite.repo = NewManageAssetRepository(suite.mockDB)
}

func TestManageAssetRepoTestSuite(t *testing.T)  {
	suite.Run(t, new(ManageAssetRepoTestSuite))
}

func (suite *ManageAssetRepoTestSuite) TestCreate_Success() {

	payload := dto.ManageAssetRequest{
		Id:                   "1",
		IdUser:               "1",
		NikStaff:             "111",
		SubmisstionDate:      time.Now(),
		ReturnDate:           time.Now(),
		Duration:             2,
		ManageAssetDetailReq: []dto.ManageAssetDetailRequest{{
			Id:            "1",
			IdManageAsset: "1",
			IdAsset:       "1",
			TotalItem:     3,
			Status:        "ok",
		}},
	}

	suite.mockSQL.ExpectBegin()
	suite.mockSQL.ExpectExec("insert into manage_asset").
	WithArgs(payload.Id, payload.IdUser, payload.NikStaff, payload.SubmisstionDate, payload.ReturnDate).WillReturnResult(sqlmock.NewResult(1, 1))

	for _, data := range payload.ManageAssetDetailReq {
		suite.mockSQL.ExpectExec("insert into detail_manage_asset").
		WithArgs(data.Id, payload.Id, data.IdAsset, data.TotalItem, data.Status).WillReturnResult(sqlmock.NewResult(1, 1))
	}

	suite.mockSQL.ExpectCommit()
	err := suite.repo.CreateTransaction(payload)
	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), err)
}

func (suite *ManageAssetRepoTestSuite) TestCreate_BeginError() {
	payload := dto.ManageAssetRequest{
		Id:                   "1",
		IdUser:               "1",
		NikStaff:             "111",
		SubmisstionDate:      time.Now(),
		ReturnDate:           time.Now(),
		Duration:             2,
		ManageAssetDetailReq: []dto.ManageAssetDetailRequest{{
			Id:            "1",
			IdManageAsset: "1",
			IdAsset:       "1",
			TotalItem:     3,
			Status:        "ok",
		}},
	}

	suite.mockSQL.ExpectBegin().WillReturnError(errors.New("failed begin tansaction"))
	err := suite.repo.CreateTransaction(payload)
	assert.Error(suite.T(), err)
	assert.NotNil(suite.T(), err)
}

func (suite *ManageAssetRepoTestSuite) TestCreate_ErrorManageQuery() {
	payload := dto.ManageAssetRequest{
		Id:                   "1",
		IdUser:               "1",
		NikStaff:             "111",
		SubmisstionDate:      time.Now(),
		ReturnDate:           time.Now(),
		Duration:             2,
		ManageAssetDetailReq: []dto.ManageAssetDetailRequest{{
			Id:            "1",
			IdManageAsset: "1",
			IdAsset:       "1",
			TotalItem:     3,
			Status:        "ok",
		}},
	}

	suite.mockSQL.ExpectBegin()
	suite.mockSQL.ExpectExec("insert into manage_asset").
	WithArgs(payload.Id, payload.IdUser, payload.NikStaff, payload.SubmisstionDate, payload.ReturnDate).WillReturnError(errors.New("failed"))
	suite.mockSQL.ExpectRollback()
	err := suite.repo.CreateTransaction(payload)
	assert.Error(suite.T(), err)
	assert.NotNil(suite.T(), err)
}

func (suite *ManageAssetRepoTestSuite) TestCreate_ErrorManageDetailQuery() {
	payload := dto.ManageAssetRequest{
		Id:                   "1",
		IdUser:               "1",
		NikStaff:             "111",
		SubmisstionDate:      time.Now(),
		ReturnDate:           time.Now(),
		Duration:             2,
		ManageAssetDetailReq: []dto.ManageAssetDetailRequest{{
			Id:            "1",
			IdManageAsset: "1",
			IdAsset:       "1",
			TotalItem:     3,
			Status:        "ok",
		}},
	}

	suite.mockSQL.ExpectBegin()
	suite.mockSQL.ExpectExec("insert into manage_asset").
	WithArgs(payload.Id, payload.IdUser, payload.NikStaff, payload.SubmisstionDate, payload.ReturnDate).WillReturnResult(sqlmock.NewResult(1, 1))
	for _, data := range payload.ManageAssetDetailReq {
		suite.mockSQL.ExpectExec("insert into detail_manage_asset").
		WithArgs(data.Id, payload.Id, data.IdAsset, data.TotalItem, data.Status).WillReturnError(errors.New("failed save manage detail"))
	}
	suite.mockSQL.ExpectRollback()
	err := suite.repo.CreateTransaction(payload)
	assert.Error(suite.T(), err)
	assert.NotNil(suite.T(), err)
}

func (suite *ManageAssetRepoTestSuite) TestFindAll_Success() {
	
	datas := []model.ManageAsset{{
		Id:             "1",
		User:           model.UserCredentials{
			ID:       "1",
			Email:    "",
			Password: "",
			Name:     "Jhon",
			IsActive: false,
		},
		Staff:          model.Staff{
			Nik_Staff:    "1",
			Name:         "Sigit",
			Phone_number: "",
			Address:      "",
			Birth_date:   time.Time{},
			Img_url:      "",
			Divisi:       "",
		},
		SubmissionDate: time.Time{},
		ReturnDate:     time.Time{},
		Detail:         []model.ManageDetailAsset{},
	}}

	rows := sqlmock.NewRows([]string{"id", "user_id", "user_name", "staff_nik", "staff_name", "submission_date", "return_date"})
	for _, data := range datas {
		rows.AddRow(data.Id, data.User.ID, data.User.Name, data.Staff.Nik_Staff, data.Staff.Name, data.SubmissionDate, data.ReturnDate)
	}
	suite.mockSQL.ExpectQuery("SELECT m.id, u.id, u.name, s.nik_staff, s.name, m.submission_date, m.return_date FROM manage_asset AS m").WillReturnRows(rows)
	result, err := suite.repo.FindAllTransaction()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), datas[0].Id, result[0].Id)
}

func (suite *ManageAssetRepoTestSuite) TestFindAll_Failed() {

	suite.mockSQL.ExpectQuery("SELECT m.id, u.id, u.name, s.nik_staff, s.name, m.submission_date, m.return_date FROM manage_asset AS m").WillReturnError(errors.New("failed get transaction"))
	result, err := suite.repo.FindAllTransaction()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
}

func (suite *ManageAssetRepoTestSuite) TestFindAll_RowsFailed() {
	
	datas := []model.ManageAsset{{
		Id:             "1",
		User:           model.UserCredentials{
			ID:       "1",
			Email:    "",
			Password: "",
			Name:     "Jhon",
			IsActive: false,
		},
		Staff:          model.Staff{
			Nik_Staff:    "1",
			Name:         "Sigit",
			Phone_number: "",
			Address:      "",
			Birth_date:   time.Time{},
			Img_url:      "",
			Divisi:       "",
		},
		SubmissionDate: time.Time{},
		ReturnDate:     time.Time{},
		Detail:         []model.ManageDetailAsset{},
	}}

	rows := sqlmock.NewRows([]string{"id", "user_id", "user_name", "staff_nik", "staff_name", "submission_date", "return_date"})
	for _, data := range datas {
		rows.AddRow(data.Id, data.User.ID, data.User.Name, data.Staff.Nik_Staff, data.Staff.Name, data.SubmissionDate, data.ReturnDate).RowError(0, errors.New("errors row"))
	}
	suite.mockSQL.ExpectQuery("SELECT m.id, u.id, u.name, s.nik_staff, s.name, m.submission_date, m.return_date FROM manage_asset AS m").WillReturnRows(rows)
	result, err := suite.repo.FindAllTransaction()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
}

func (suite *ManageAssetRepoTestSuite) TestFindAllById_Success() {
	datas := []model.ManageAsset{{
		Id:             "1",
		User:           model.UserCredentials{
			ID:       "1",
			Email:    "",
			Password: "",
			Name:     "Jhon",
			IsActive: false,
		},
		Staff:          model.Staff{
			Nik_Staff:    "1",
			Name:         "Sigit",
			Phone_number: "",
			Address:      "",
			Birth_date:   time.Time{},
			Img_url:      "",
			Divisi:       "",
		},
		SubmissionDate: time.Time{},
		ReturnDate:     time.Time{},
		Detail:         []model.ManageDetailAsset{},
	}}

	dataDetails := []model.ManageDetailAsset{{
		Id:            "1",
		ManageAssetId: "1",
		Asset:         model.Asset{
			Id:        "1",
			Name:      "Laptop",
		},
		TotalItem:     2,
		Status:        "Ready",
	}}
	rows := sqlmock.NewRows([]string{"id", "user_id", "user_name", "staff_nik", "staff_name", "submission_date", "return_date"})
	for _, data := range datas {
		rows.AddRow(data.Id, data.User.ID, data.User.Name, data.Staff.Nik_Staff, data.Staff.Name, data.SubmissionDate, data.ReturnDate)
	}
	rowDetail := sqlmock.NewRows([]string{"id", "id_manage", "id_asset", "asset_name", "total_item", "status"})
	for _, data := range dataDetails {
		rowDetail.AddRow(data.Id, data.ManageAssetId, data.Asset.Id, data.Asset.Name, data.TotalItem, data.Status)
	}

	suite.mockSQL.ExpectQuery("SELECT m.id, u.id, u.name, s.nik_staff, s.name, m.submission_date, m.return_date FROM manage_asset AS m").WillReturnRows(rows)
	suite.mockSQL.ExpectQuery("SELECT d.id, d.id_manage_asset, a.id, a.name, d.total_item, d.status FROM detail_manage_asset AS d").WillReturnRows(rowDetail)

	result, resulDetail, err := suite.repo.FindAllByTransId("1")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), datas[0].Id, result[0].Id)
	assert.Equal(suite.T(), datas[0].Id, resulDetail[0].Id)

}
func (suite *ManageAssetRepoTestSuite) TestFindAllById_Failed() {
	datas := []model.ManageAsset{{
		Id:             "1",
		User:           model.UserCredentials{
			ID:       "1",
			Email:    "",
			Password: "",
			Name:     "Jhon",
			IsActive: false,
		},
		Staff:          model.Staff{
			Nik_Staff:    "1",
			Name:         "Sigit",
			Phone_number: "",
			Address:      "",
			Birth_date:   time.Time{},
			Img_url:      "",
			Divisi:       "",
		},
		SubmissionDate: time.Time{},
		ReturnDate:     time.Time{},
		Detail:         []model.ManageDetailAsset{},
	}}

	// dataDetails := []model.ManageDetailAsset{{
	// 	Id:            "1",
	// 	ManageAssetId: "1",
	// 	Asset:         model.Asset{
	// 		Id:        "1",
	// 		Name:      "Laptop",
	// 	},
	// 	TotalItem:     2,
	// 	Status:        "Ready",
	// }}
	rows := sqlmock.NewRows([]string{"id", "user_id", "user_name", "staff_nik", "staff_name", "submission_date", "return_date"})
	for _, data := range datas {
		rows.AddRow(data.Id, data.User.ID, data.User.Name, data.Staff.Nik_Staff, data.Staff.Name, data.SubmissionDate, data.ReturnDate)
	}
	// rowDetail := sqlmock.NewRows([]string{"id", "id_manage", "id_asset", "asset_name", "total_item", "status"})
	// for _, data := range dataDetails {
	// 	rowDetail.AddRow(data.Id, data.ManageAssetId, data.Asset.Id, data.Asset.Name, data.TotalItem, data.Status)
	// }

	suite.mockSQL.ExpectQuery("SELECT m.id, u.id, u.name, s.nik_staff, s.name, m.submission_date, m.return_date FROM manage_asset AS m").WillReturnError(errors.New("error get manage asset"))
	
	result, resultDetail, err := suite.repo.FindAllByTransId("1")
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Nil(suite.T(), resultDetail)
	
	suite.mockSQL.ExpectQuery("SELECT m.id, u.id, u.name, s.nik_staff, s.name, m.submission_date, m.return_date FROM manage_asset AS m").WillReturnRows(rows)
	suite.mockSQL.ExpectQuery("SELECT d.id, d.id_manage_asset, a.id, a.name, d.total_item, d.status FROM detail_manage_asset AS d").WillReturnError(errors.New("errors get manage detail"))
	result, resultDetail, err = suite.repo.FindAllByTransId("1")
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Nil(suite.T(), resultDetail)
}

func (suite *ManageAssetRepoTestSuite) TestFindAllById_ManageRowsError() {
	datas := []model.ManageAsset{{
		Id:             "1",
		User:           model.UserCredentials{
			ID:       "1",
			Email:    "",
			Password: "",
			Name:     "Jhon",
			IsActive: false,
		},
		Staff:          model.Staff{
			Nik_Staff:    "1",
			Name:         "Sigit",
			Phone_number: "",
			Address:      "",
			Birth_date:   time.Time{},
			Img_url:      "",
			Divisi:       "",
		},
		SubmissionDate: time.Time{},
		ReturnDate:     time.Time{},
		Detail:         []model.ManageDetailAsset{},
	}}

	dataDetails := []model.ManageDetailAsset{{
		Id:            "1",
		ManageAssetId: "1",
		Asset:         model.Asset{
			Id:        "1",
			Name:      "Laptop",
		},
		TotalItem:     2,
		Status:        "Ready",
	}}

	//rows manage
	rows := sqlmock.NewRows([]string{"id", "user_id", "user_name", "staff_nik", "staff_name", "submission_date", "return_date"})
	for _, data := range datas {
		rows.AddRow(data.Id, data.User.ID, data.User.Name, data.Staff.Nik_Staff, data.Staff.Name, data.SubmissionDate, data.ReturnDate).RowError(0, errors.New("erros row"))
	}

	rowDetail := sqlmock.NewRows([]string{"id", "id_manage", "id_asset", "asset_name", "total_item", "status"})
	for _, data := range dataDetails {
		rowDetail.AddRow(data.Id, data.ManageAssetId, data.Asset.Id, data.Asset.Name, data.TotalItem, data.Status)
	}
	suite.mockSQL.ExpectQuery("SELECT m.id, u.id, u.name, s.nik_staff, s.name, m.submission_date, m.return_date FROM manage_asset AS m").WillReturnRows(rows)
	suite.mockSQL.ExpectQuery("SELECT d.id, d.id_manage_asset, a.id, a.name, d.total_item, d.status FROM detail_manage_asset AS d").WillReturnRows(rowDetail)

	result, resultDetail, err := suite.repo.FindAllByTransId("1")
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Nil(suite.T(), resultDetail)

	newRows := sqlmock.NewRows([]string{"id", "user_id", "user_name", "staff_nik", "staff_name", "submission_date", "return_date"})
	for _, data := range datas {
		newRows.AddRow(data.Id, data.User.ID, data.User.Name, data.Staff.Nik_Staff, data.Staff.Name, data.SubmissionDate, data.ReturnDate)
	}

	newRowDetail := sqlmock.NewRows([]string{"id", "id_manage", "id_asset", "asset_name", "total_item", "status"})
	for _, data := range dataDetails {
		newRowDetail.AddRow(data.Id, data.ManageAssetId, data.Asset.Id, data.Asset.Name, data.TotalItem, data.Status).RowError(0, errors.New("errors row"))
	}

	suite.mockSQL.ExpectQuery("SELECT m.id, u.id, u.name, s.nik_staff, s.name, m.submission_date, m.return_date FROM manage_asset AS m").WillReturnRows(newRows)
	suite.mockSQL.ExpectQuery("SELECT d.id, d.id_manage_asset, a.id, a.name, d.total_item, d.status FROM detail_manage_asset AS d").WillReturnRows(newRowDetail)
	result, resultDetail, err = suite.repo.FindAllByTransId("1")
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Nil(suite.T(), resultDetail)

}

func (suite *ManageAssetRepoTestSuite) TestFindAllByName_Success() {
	datas := []model.ManageAsset{{
		Id:             "1",
		User:           model.UserCredentials{
			ID:       "1",
			Email:    "",
			Password: "",
			Name:     "Jhon",
			IsActive: false,
		},
		Staff:          model.Staff{
			Nik_Staff:    "1",
			Name:         "Sigit",
			Phone_number: "",
			Address:      "",
			Birth_date:   time.Time{},
			Img_url:      "",
			Divisi:       "",
		},
		SubmissionDate: time.Time{},
		ReturnDate:     time.Time{},
		Detail:         []model.ManageDetailAsset{},
	}}

	dataDetails := []model.ManageDetailAsset{{
		Id:            "1",
		ManageAssetId: "1",
		Asset:         model.Asset{
			Id:        "1",
			Name:      "Laptop",
		},
		TotalItem:     2,
		Status:        "Ready",
	}}
	rows := sqlmock.NewRows([]string{"id", "user_id", "user_name", "staff_nik", "staff_name", "submission_date", "return_date"})
	for _, data := range datas {
		rows.AddRow(data.Id, data.User.ID, data.User.Name, data.Staff.Nik_Staff, data.Staff.Name, data.SubmissionDate, data.ReturnDate)
	}
	rowDetail := sqlmock.NewRows([]string{"id", "id_manage", "id_asset", "asset_name", "total_item", "status"})
	for _, data := range dataDetails {
		rowDetail.AddRow(data.Id, data.ManageAssetId, data.Asset.Id, data.Asset.Name, data.TotalItem, data.Status)
	}

	suite.mockSQL.ExpectQuery("SELECT m.id, u.id, u.name, s.nik_staff, s.name, m.submission_date, m.return_date FROM manage_asset AS m").WithArgs("%"+"Jhon"+"%").WillReturnRows(rows)
	suite.mockSQL.ExpectQuery("SELECT d.id, d.id_manage_asset, a.id, a.name, d.total_item, d.status FROM detail_manage_asset AS d").WillReturnRows(rowDetail)

	result, resulDetail, err := suite.repo.FindByNameTransaction("Jhon")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), datas[0].Id, result[0].Id)
	assert.Equal(suite.T(), datas[0].Id, resulDetail[0].Id)

}

func (suite *ManageAssetRepoTestSuite) TestFindAllByName_Failed() {
	datas := []model.ManageAsset{{
		Id:             "1",
		User:           model.UserCredentials{
			ID:       "1",
			Email:    "",
			Password: "",
			Name:     "Jhon",
			IsActive: false,
		},
		Staff:          model.Staff{
			Nik_Staff:    "1",
			Name:         "Sigit",
			Phone_number: "",
			Address:      "",
			Birth_date:   time.Time{},
			Img_url:      "",
			Divisi:       "",
		},
		SubmissionDate: time.Time{},
		ReturnDate:     time.Time{},
		Detail:         []model.ManageDetailAsset{},
	}}

	rows := sqlmock.NewRows([]string{"id", "user_id", "user_name", "staff_nik", "staff_name", "submission_date", "return_date"})
	for _, data := range datas {
		rows.AddRow(data.Id, data.User.ID, data.User.Name, data.Staff.Nik_Staff, data.Staff.Name, data.SubmissionDate, data.ReturnDate)
	}

	suite.mockSQL.ExpectQuery("SELECT m.id, u.id, u.name, s.nik_staff, s.name, m.submission_date, m.return_date FROM manage_asset AS m").WillReturnError(errors.New("error get manage asset"))
	
	result, resultDetail, err := suite.repo.FindByNameTransaction("Jhon")
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Nil(suite.T(), resultDetail)
	
	suite.mockSQL.ExpectQuery("SELECT m.id, u.id, u.name, s.nik_staff, s.name, m.submission_date, m.return_date FROM manage_asset AS m").WillReturnRows(rows)
	suite.mockSQL.ExpectQuery("SELECT d.id, d.id_manage_asset, a.id, a.name, d.total_item, d.status FROM detail_manage_asset AS d").WillReturnError(errors.New("errors get manage detail"))
	result, resultDetail, err = suite.repo.FindByNameTransaction("John")
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Nil(suite.T(), resultDetail)
}

func (suite *ManageAssetRepoTestSuite) TestFindAllByName_ManageRowsError() {
	datas := []model.ManageAsset{{
		Id:             "1",
		User:           model.UserCredentials{
			ID:       "1",
			Email:    "",
			Password: "",
			Name:     "Jhon",
			IsActive: false,
		},
		Staff:          model.Staff{
			Nik_Staff:    "1",
			Name:         "Sigit",
			Phone_number: "",
			Address:      "",
			Birth_date:   time.Time{},
			Img_url:      "",
			Divisi:       "",
		},
		SubmissionDate: time.Time{},
		ReturnDate:     time.Time{},
		Detail:         []model.ManageDetailAsset{},
	}}

	dataDetails := []model.ManageDetailAsset{{
		Id:            "1",
		ManageAssetId: "1",
		Asset:         model.Asset{
			Id:        "1",
			Name:      "Laptop",
		},
		TotalItem:     2,
		Status:        "Ready",
	}}

	//rows manage
	rows := sqlmock.NewRows([]string{"id", "user_id", "user_name", "staff_nik", "staff_name", "submission_date", "return_date"})
	for _, data := range datas {
		rows.AddRow(data.Id, data.User.ID, data.User.Name, data.Staff.Nik_Staff, data.Staff.Name, data.SubmissionDate, data.ReturnDate).RowError(0, errors.New("erros row"))
	}

	rowDetail := sqlmock.NewRows([]string{"id", "id_manage", "id_asset", "asset_name", "total_item", "status"})
	for _, data := range dataDetails {
		rowDetail.AddRow(data.Id, data.ManageAssetId, data.Asset.Id, data.Asset.Name, data.TotalItem, data.Status)
	}
	suite.mockSQL.ExpectQuery("SELECT m.id, u.id, u.name, s.nik_staff, s.name, m.submission_date, m.return_date FROM manage_asset AS m").WillReturnRows(rows)
	suite.mockSQL.ExpectQuery("SELECT d.id, d.id_manage_asset, a.id, a.name, d.total_item, d.status FROM detail_manage_asset AS d").WillReturnRows(rowDetail)

	result, resultDetail, err := suite.repo.FindByNameTransaction("Jhon")
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Nil(suite.T(), resultDetail)

	newRows := sqlmock.NewRows([]string{"id", "user_id", "user_name", "staff_nik", "staff_name", "submission_date", "return_date"})
	for _, data := range datas {
		newRows.AddRow(data.Id, data.User.ID, data.User.Name, data.Staff.Nik_Staff, data.Staff.Name, data.SubmissionDate, data.ReturnDate)
	}

	newRowDetail := sqlmock.NewRows([]string{"id", "id_manage", "id_asset", "asset_name", "total_item", "status"})
	for _, data := range dataDetails {
		newRowDetail.AddRow(data.Id, data.ManageAssetId, data.Asset.Id, data.Asset.Name, data.TotalItem, data.Status).RowError(0, errors.New("errors row"))
	}

	suite.mockSQL.ExpectQuery("SELECT m.id, u.id, u.name, s.nik_staff, s.name, m.submission_date, m.return_date FROM manage_asset AS m").WillReturnRows(newRows)
	suite.mockSQL.ExpectQuery("SELECT d.id, d.id_manage_asset, a.id, a.name, d.total_item, d.status FROM detail_manage_asset AS d").WillReturnRows(newRowDetail)
	result, resultDetail, err = suite.repo.FindByNameTransaction("Jhon")
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Nil(suite.T(), resultDetail)

}