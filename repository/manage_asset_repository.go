package repository

import (
	"database/sql"
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"
	"fmt"
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
	var results []model.ManageDetailAsset

	//query := `select m.id, u.id, u.name, s.nik_staff, s.name, submission_date, return_date
	//      from manage_asset as m
	//      join user_credential as u on u.id = m.id_user
	//      join staff as s on s.nik_staff = m.nik_staff
	//      where m.id = $1`

	query := `SELECT
    m.id,
    m.id_user,
    u.id AS user_id,
    u.name AS user_name,
    s.nik_staff AS staff_nik,
    s.name AS staff_name,
    m.submission_date,
    m.return_date,
    a.id AS asset_id,
    a.name AS asset_name,
    a.status AS asset_status,
    a.entry_date AS asset_entry_date,
    a.img_url AS asset_img_url,
    d.total_item,
    d.status AS detail_status
	FROM
    manage_asset AS m
	JOIN
    user_credential AS u ON u.id = m.id_user
	JOIN
    staff AS s ON s.nik_staff = m.nik_staff
	JOIN
    detail_manage_asset AS d ON d.id_manage_asset = m.id
	JOIN
    asset AS a ON a.id = d.id_asset
	WHERE
    m.id = $1
`

	rows, err := m.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var result model.ManageDetailAsset
		if err = rows.Scan(
			&result.Id,
			&result.ManageAsset.Id,
			&result.ManageAsset.User.ID,
			&result.ManageAsset.User.Name,
			&result.ManageAsset.Staff.Nik_Staff,
			&result.ManageAsset.Staff.Name,
			&result.ManageAsset.SubmissionDate,
			&result.ManageAsset.ReturnDate,
			&result.Asset.Id,
			&result.Asset.Name,
			//&result.Asset.Amount,
			&result.Asset.Status,
			&result.Asset.EntryDate,
			&result.Asset.ImgUrl,
			&result.TotalItem,
			&result.Status,
		); err != nil {
			return nil, fmt.Errorf("Failed to scan row: %v", err)
		}
		results = append(results, result)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// FindAll implements ManageAssetRepository.
//func (m *manageAssetRepository) FindAllTransaction() ([]model.ManageAsset, error) {
//
//	//query := `select m.id, u.id, u.name, s.nik_staff, s.name, d.id, a.id, a.name, d.total_item, d.status from manage_asset
//	//join user_credential as u on u.id = m.id_user
//	//join staff as s on s.nik_staff = m.nik_staff
//	//join detail_manage_asset as d on d.id_manage_asset = m.id
//	//join asset as a on a.id = d.id_asset`
//
//	query := `
//    SELECT m.id, u.id, u.name, s.nik_staff, s.name, d.id, a.id, a.name, d.total_item, d.status
//    FROM manage_asset AS m
//    JOIN user_credential AS u ON u.id = m.id_user
//    JOIN staff AS s ON s.nik_staff = m.nik_staff
//    JOIN detail_manage_asset AS d ON d.id_manage_asset = m.id
//    JOIN asset AS a ON a.id = d.id_asset
//`
//
//	//query := `select m.id, u.id, u.name, s.nik_staff, s.name, submission_date, return_date from manage_asset as m
//	//		join user_credential as u on u.id = m.id_user
//	//		join staff as s on s.nik_staff = m.nik_staff`
//
//	rows, err := m.db.Query(query)
//	if err != nil {
//		return nil, err
//	}
//
//	var transactions []model.ManageAsset
//	for rows.Next() {
//		var t model.ManageAsset
//		var staff model.Staff
//		var detail model.ManageDetailAsset
//
//		err = rows.Scan(
//			&t.Id,
//			&t.User.ID,
//			&t.User.Name,
//			&staff.Nik_Staff,
//			&staff.Name,
//			&detail.Id,
//			&detail.Asset.Id,
//			&detail.Asset.Name,
//			&detail.TotalItem,
//			&detail.Status,
//		)
//		if err != nil {
//			return nil, err
//		}
//
//		// Setel staf ke dalam manajemen detail aset
//		detail.ManageAsset = t
//		// Setel staf ke dalam manajemen aset
//		t.Staff = staff
//		// Tambahkan detail manajemen aset ke dalam manajemen aset
//		t.Detail = append(t.Detail, detail)
//
//		transactions = append(transactions, t)
//	}
//	if rows.Err() != nil {
//		return nil, rows.Err()
//	}
//
//	return transactions, nil
//}

func (m *manageAssetRepository) FindAllTransaction() ([]model.ManageAsset, error) {
	query := `
    SELECT
        m.id,
        u.id,
        u.name,
        s.nik_staff,
        s.name,
        d.id,
        a.id,
        a.name,
        a.id_category AS category_id,
        c.name AS name,
        d.total_item,
        d.status
    FROM
        manage_asset AS m
    JOIN
        user_credential AS u ON u.id = m.id_user
    JOIN
        staff AS s ON s.nik_staff = m.nik_staff
    JOIN
        detail_manage_asset AS d ON d.id_manage_asset = m.id
    JOIN
        asset AS a ON a.id = d.id_asset
    JOIN
        category AS c ON a.id_category = c.id
`

	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []model.ManageAsset
	for rows.Next() {
		var t model.ManageAsset
		var staff model.Staff
		var detail model.ManageDetailAsset
		var categoryID, categoryName string

		err = rows.Scan(
			&t.Id,
			&t.User.ID,
			&t.User.Name,
			&staff.Nik_Staff,
			&staff.Name,
			&detail.Id,
			&detail.Asset.Id,
			&detail.Asset.Name,
			&categoryID,
			&categoryName,
			&detail.TotalItem,
			&detail.Status,
		)
		if err != nil {
			return nil, err
		}

		// Setel staf ke dalam manajemen detail aset
		detail.ManageAsset = t
		// Setel staf ke dalam manajemen aset
		t.Staff = staff
		// Tambahkan detail manajemen aset ke dalam manajemen aset
		t.Detail = append(t.Detail, detail)

		// Setel informasi kategori
		t.Detail[len(t.Detail)-1].Asset.Category.Id = categoryID
		t.Detail[len(t.Detail)-1].Asset.Category.Name = categoryName

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
		return err
	}
	_, err = tx.Exec(query, payload.Id, payload.IdUser, payload.NikStaff, payload.SubmisstionDate, payload.ReturnDate)
	if err != nil {
		return err
	}

	queryDetail := "insert into detail_manage_asset(id, id_asset, id_manage_asset, total_item, status) values ($1, $2, $3, $4, $5)"
	for _, v := range payload.ManageAssetDetailReq {
		_, err = tx.Exec(queryDetail, v.Id, v.IdAsset, v.IdManageAsset, v.TotalItem, v.Status)
		if err != nil {
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
