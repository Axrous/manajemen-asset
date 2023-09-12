package repository

import (
	"database/sql"
	"final-project-enigma-clean/model"
)

type AssetRepository interface {
	Save(asset model.AssetRequest) error
	FindAll() ([]model.Asset, error)
	FindById(id string) (model.Asset, error)
	Update(asset model.AssetRequest) error
	Delete(id string) error
}

type assetRepository struct {
	db *sql.DB
}

// Delete implements AssetRepository.
func (a *assetRepository) Delete(id string) error {

	query := "delete from asset where id = $1"

	_, err := a.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

// FindAll implements AssetRepository.
func (a *assetRepository) FindAll() ([]model.Asset, error) {
	// panic("unimplemented")

	query := `select a.id, a.name, a.amount, a.status, a.entry_date, a.img_url, c.id, c.name, at.id, at.name
			from asset as a 
			left join category as c on c.id = a.id_category
			left join asset_type as at on at.id = a.id_asset_type`
	
	rows, err := a.db.Query(query)
	if err != nil {
		return nil, err
	}

	var assets []model.Asset
	for rows.Next() {
		var asset model.Asset
		rows.Scan(&asset.ID, &asset.Name, &asset.Amount, &asset.Status, &asset.EntryDate, &asset.ImgUrl, &asset.Category.ID, &asset.Category.Name, &asset.AssetType.ID, &asset.AssetType.Name)
		assets = append(assets, asset)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}


	return assets, nil
}

// FindById implements AssetRepository.
func (a *assetRepository) FindById(id string) (model.Asset, error) {
	query := `select a.id, a.name, a.amount, a.status, a.entry_date, a.img_url, c.id, c.name, at.id, at.name
			from asset as a 
			left join category as c on c.id = a.id_category
			left join asset_type as at on at.id = a.id_asset_type
			where a.id = $1`
	
	row := a.db.QueryRow(query, id)
	var asset model.Asset
	err := row.Scan(&asset.ID, &asset.Name, &asset.Amount, &asset.Status, &asset.EntryDate, &asset.ImgUrl, &asset.Category.ID, &asset.Category.Name, &asset.AssetType.ID, &asset.AssetType.Name)
	if err != nil {
		return model.Asset{}, err
	}

	return asset, nil
}

// Save implements AssetRepository.
func (a *assetRepository) Save(asset model.AssetRequest) error {
	query := "insert into asset(id, id_category, id_asset_type, name, amount, status, entry_date, img_url) values($1, $2, $3, $4, $5, $6, $7, $8)"

	_, err := a.db.Exec(query, asset.ID, asset.CategoryId, asset.AssetTypeId, asset.Name, asset.Amount, asset.Status, asset.EntryDate, asset.ImgUrl)
	if err != nil {
		return err
	}

	return nil
}

// Update implements AssetRepository.
func (a *assetRepository) Update(asset model.AssetRequest) error {
	query := `update asset set id_category = $2, id_asset_type = $3, name = $4, amount = $5, status = $6, img_url = $7 where id = $1`

	_, err := a.db.Exec(query, asset.ID, asset.CategoryId, asset.AssetTypeId, asset.Name, asset.Amount, asset.Status, asset.ImgUrl)
	if err !=nil {
		return err
	}

	return nil
}

func NewAssetRepository(db *sql.DB) AssetRepository {
	return &assetRepository{
		db: db,
	}
}
