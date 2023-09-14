package repository

import (
	"database/sql"
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"
)

type ManageAssetRepository interface {
	CreateTransaction(payload dto.ManageAssetRequest) error
	FindAllTransaction() ([]model.ManageAsset, error)
	FindAllByTransId(id string) ([]model.ManageAsset, []model.ManageDetailAsset, error)
	FindByNameTransaction(name string) ([]model.ManageAsset, []model.ManageDetailAsset, error)
}

type manageAssetRepository struct {
	db *sql.DB
}

// FindByNameTransaction implements ManageAssetRepository.
func (m *manageAssetRepository) FindByNameTransaction(name string) ([]model.ManageAsset, []model.ManageDetailAsset, error) {
	query := `
    SELECT m.id, u.id, u.name, s.nik_staff, s.name, m.submission_date, m.return_date
    FROM manage_asset AS m
    JOIN user_credential AS u ON u.id = m.id_user
    JOIN staff AS s ON s.nik_staff = m.nik_staff
	where s.name ilike $1`

	queryDetail := `SELECT d.id, d.id_manage_asset, a.id, a.name, d.total_item, d.status FROM detail_manage_asset AS d
	JOIN asset AS a ON a.id = d.id_asset`

	rows, err := m.db.Query(query, "%"+name+"%")
	if err != nil {
		return nil, nil, err
	}
	rowsDetail, err := m.db.Query(queryDetail)
	if err != nil {
		return nil, nil, err
	}

	var transactions []model.ManageAsset
	for rows.Next() {
		var t model.ManageAsset
		rows.Scan(&t.Id, &t.User.ID, &t.User.Name, &t.Staff.Nik_Staff, &t.Staff.Name, &t.SubmissionDate, &t.ReturnDate)
		transactions = append(transactions, t)
	}
	if rows.Err() != nil {
		return nil,nil, rows.Err()
	}

	var transactionDetail []model.ManageDetailAsset
	for rowsDetail.Next() {
		var td model.ManageDetailAsset
		rowsDetail.Scan(&td.Id, &td.ManageAssetId, &td.Asset.Id, &td.Asset.Name, &td.TotalItem, &td.Status)
		transactionDetail = append(transactionDetail, td)
	}
	if rows.Err() != nil {
		return nil, nil, rows.Err()
	}

	return transactions, transactionDetail, nil
}

// FindAllByTransId implements ManageAssetRepository.
func (m *manageAssetRepository) FindAllByTransId(id string) ([]model.ManageAsset, []model.ManageDetailAsset, error) {
	query := `
    SELECT m.id, u.id, u.name, s.nik_staff, s.name, m.submission_date, m.return_date
    FROM manage_asset AS m
    JOIN user_credential AS u ON u.id = m.id_user
    JOIN staff AS s ON s.nik_staff = m.nik_staff
	where m.id = $1`

	queryDetail := `SELECT d.id, d.id_manage_asset, a.id, a.name, d.total_item, d.status FROM detail_manage_asset AS d
	JOIN asset AS a ON a.id = d.id_asset`

	rows, err := m.db.Query(query, id)
	if err != nil {
		return nil, nil, err
	}
	rowsDetail, err := m.db.Query(queryDetail)
	if err != nil {
		return nil, nil, err
	}

	var transactions []model.ManageAsset
	for rows.Next() {
		var t model.ManageAsset
		rows.Scan(&t.Id, &t.User.ID, &t.User.Name, &t.Staff.Nik_Staff, &t.Staff.Name, &t.SubmissionDate, &t.ReturnDate)
		transactions = append(transactions, t)
	}
	if rows.Err() != nil {
		return nil,nil, rows.Err()
	}

	var transactionDetail []model.ManageDetailAsset
	for rowsDetail.Next() {
		var td model.ManageDetailAsset
		rowsDetail.Scan(&td.Id, &td.ManageAssetId, &td.Asset.Id, &td.Asset.Name, &td.TotalItem, &td.Status)
		transactionDetail = append(transactionDetail, td)
	}
	if rows.Err() != nil {
		return nil, nil, rows.Err()
	}

	return transactions, transactionDetail, nil
}

// FindAll implements ManageAssetRepository.
func (m *manageAssetRepository) FindAllTransaction() ([]model.ManageAsset, error) {

	query := `SELECT m.id, u.id, u.name, s.nik_staff, s.name, m.submission_date, m.return_date
    FROM manage_asset AS m
    JOIN user_credential AS u ON u.id = m.id_user
    JOIN staff AS s ON s.nik_staff = m.nik_staff`

	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}

	var transactions []model.ManageAsset
	for rows.Next() {
		var tr model.ManageAsset
		rows.Scan(&tr.Id, &tr.User.ID, &tr.User.Name, &tr.Staff.Nik_Staff, &tr.Staff.Name, &tr.SubmissionDate, &tr.ReturnDate)
		transactions = append(transactions, tr)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return transactions, nil
}

// CreateTransaksi implements ManageAssetRepository.
func (m *manageAssetRepository) CreateTransaction(payload dto.ManageAssetRequest) error {

	query := "insert into manage_asset(id, id_user, nik_staff, submission_date, return_date) values($1, $2, $3, $4, $5)"

	tx, err := m.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(query, payload.Id, payload.IdUser, payload.NikStaff, payload.SubmisstionDate, payload.ReturnDate)
	if err != nil {
		tx.Rollback()
		return err
	}

	queryDetail := "insert into detail_manage_asset(id, id_asset, id_manage_asset, total_item, status) values ($1, $2, $3, $4, $5)"
	for _, v := range payload.ManageAssetDetailReq {
		_, err = tx.Exec(queryDetail, v.Id, v.IdAsset, v.IdManageAsset, v.TotalItem, v.Status)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func NewManageAssetRepository(db *sql.DB) ManageAssetRepository {
	return &manageAssetRepository{
		db: db,
	}
}
