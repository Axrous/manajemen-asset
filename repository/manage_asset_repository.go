package repository

import (
	"database/sql"
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"
)

type ManageAssetRepository interface {
	CreateTransaction(payload dto.ManageAssetRequest) error
	FindAllTransaction() ([]model.ManageAsset, error)
	FindAllByTransId(id string) ([]model.ManageDetailAsset, error)
}

type manageAssetRepository struct {
	db *sql.DB
}

// FindAllByTransId implements ManageAssetRepository.
func (m *manageAssetRepository) FindAllByTransId(id string) ([]model.ManageDetailAsset, error) {
	query := "select d.id, d.id_manage_asset, a.id, a.name, total_item, status from manage_detail_asset as d left join asset as a on a.id = d.id_asset where d.id = $1"

	rows, err := m.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {

	}
	panic("")
}

// FindAll implements ManageAssetRepository.
func (m *manageAssetRepository) FindAllTransaction() ([]model.ManageAsset, error) {

	// query := `select m.id, u.id, u.name, s.nik_staff, s.name, d.id, a.id, a.name, d.total_item, d.status from manage_asset
	// join user_credential as u on u.id = m.id_user
	// join staff as s on s.nik_staff = m.nik_staff
	// join detail_manage_asset as d on d.id_manage_asset = m.id
	// join asset as a on a.id = d.id_asset`
	query := `select m.id, u.id, u.name, s.nik_staff, s.name, submission_date, return_date from manage_asset as m
			join user_credential as u on u.id = m.id_user
			join staff as s on s.nik_staff = m.nik_staff`

	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}

	var transactions []model.ManageAsset
	for rows.Next() {
		var t model.ManageAsset
		rows.Scan(&t.Id,
			&t.User.ID,
			&t.User.Name,
			&t.Staff.Nik_Staff,
			&t.Staff.Name,
			&t.SubmissionDate,
			&t.ReturnDate,
		)
		transactions = append(transactions, t)
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
