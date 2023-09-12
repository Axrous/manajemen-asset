package repository

import (
	"database/sql"
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"
	"math"
)

type TypeAssetRepository interface {
	Save(payload model.TypeAsset) error
	FindById(id string) (model.TypeAsset, error)
	FindByName(name string) ([]model.TypeAsset, error)
	FindAll() ([]model.TypeAsset, error)
	Update(payload model.TypeAsset) error
	Delete(id string) error
	Paging(payload dto.PageRequest) ([]model.TypeAsset, dto.Paging, error)
}

type typeAssetRepository struct {
	db *sql.DB
}

// FindById implements TypeAssetRepository.
func (t *typeAssetRepository) FindById(id string) (model.TypeAsset, error) {
	row := t.db.QueryRow("SELECT id,name FROM asset_type WHERE id=$1", id)
	var typeAsset model.TypeAsset
	err := row.Scan(&typeAsset.Id, &typeAsset.Name)
	if err != nil {
		return model.TypeAsset{}, err
	}

	return typeAsset, nil

}

// Delete implements TypeAssetRepository.
func (t *typeAssetRepository) Delete(id string) error {
	_, err := t.db.Exec("DELETE FROM asset_type WHERE id= $1", id)
	if err != nil {
		return err
	}
	return nil
}

// FindAll implements TypeAssetRepository.
func (t *typeAssetRepository) FindAll() ([]model.TypeAsset, error) {
	rows, err := t.db.Query("SELECT * FROM asset_type")
	if err != nil {
		return nil, err
	}
	var typeAssets []model.TypeAsset
	for rows.Next() {
		var typeAsset model.TypeAsset
		err = rows.Scan(&typeAsset.Id, &typeAsset.Name)
		if err != nil {
			return nil, err
		}
		typeAssets = append(typeAssets, typeAsset)
	}
	return typeAssets, nil
}

// FindByName implements TypeAssetRepository.
func (t *typeAssetRepository) FindByName(name string) ([]model.TypeAsset, error) {
	rows, err := t.db.Query(`SELECT id, name FROM asset_type WHERE name ILIKE $1`, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	var typeAssets []model.TypeAsset
	for rows.Next() {
		typeAsset := model.TypeAsset{}
		err := rows.Scan(
			&typeAsset.Id,
			&typeAsset.Name,
		)
		if err != nil {
			return nil, err
		}
	}
	return typeAssets, nil

}

// Paging implements TypeAssetRepository.
func (t *typeAssetRepository) Paging(payload dto.PageRequest) ([]model.TypeAsset, dto.Paging, error) {
	if payload.Page <= 0 {
		payload.Page = 1
	}
	q := `SELECT id, name FROM asset_type LIMIT $2 OFFSET $1`
	rows, err := t.db.Query(q, (payload.Page-1)*payload.Size, payload.Size)
	if err != nil {
		return nil, dto.Paging{}, err
	}

	var typeAssets []model.TypeAsset
	for rows.Next() {
		var typeAsset model.TypeAsset
		err := rows.Scan(&typeAsset.Id, &typeAsset.Name)
		if err != nil {
			return nil, dto.Paging{}, err
		}
		typeAssets = append(typeAssets, typeAsset)
	}
	var count int
	row := t.db.QueryRow("SELECT COUNT(id) FROM asset_type")
	if err := row.Scan(&count); err != nil {
		return nil, dto.Paging{}, err
	}

	paging := dto.Paging{
		Page:       payload.Page,
		Size:       payload.Size,
		TotalRows:  count,
		TotalPages: int(math.Ceil(float64(count) / float64(payload.Size))), // (totalrow / size)
	}

	return typeAssets, paging, nil

}

// Save implements TypeAssetRepository.
func (t *typeAssetRepository) Save(payload model.TypeAsset) error {
	_, err := t.db.Exec("INSERT INTO asset_type VALUES ($1,$2)", payload.Id, payload.Name)
	if err != nil {
		return err
	}
	return nil
}

// Update implements TypeAssetRepository.
func (t *typeAssetRepository) Update(payload model.TypeAsset) error {
	_, err := t.db.Exec("UPDATE asset_type SET name=$2 WHERE id=$1", payload.Id, payload.Name)
	if err != nil {
		return err
	}
	return nil
}

func NewTypeAssetRepository(db *sql.DB) TypeAssetRepository {
	return &typeAssetRepository{
		db: db,
	}
}
