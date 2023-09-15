package repository

import (
	"database/sql"
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"
	"math"
)

type AssetRepository interface {
	Save(asset model.AssetRequest) error
	FindAll() ([]model.Asset, error)
	FindById(id string) (model.Asset, error)
	FindByName(name string) ([]model.Asset, error)
	Update(asset model.AssetRequest) error
	UpdateAvailable(id string, amount int) error
	Delete(id string) error
	Paging(payload dto.PageRequest) ([]model.Asset, dto.Paging, error)
}

type assetRepository struct {
	db *sql.DB
}

// Paging implements AssetRepository.
func (a *assetRepository) Paging(payload dto.PageRequest) ([]model.Asset, dto.Paging, error) {
	q := `select a.id, a.name, a.available, a.status, a.entry_date, a.img_url, a.total, c.id, c.name, at.id, at.name
	from asset as a 
	left join category as c on c.id = a.id_category
	left join asset_type as at on at.id = a.id_asset_type
	limit $2 offset $1`

	rows, err := a.db.Query(q, (payload.Page-1)*payload.Size, payload.Size)
	if err != nil {
		return nil, dto.Paging{}, err
	}

	var assets []model.Asset
	for rows.Next() {
		var asset model.Asset
		rows.Scan(&asset.Id, &asset.Name, &asset.Available, &asset.Status, &asset.EntryDate, &asset.ImgUrl, &asset.Total, &asset.Category.Id, &asset.Category.Name, &asset.AssetType.Id, &asset.AssetType.Name)
		assets = append(assets, asset)
	}
	if rows.Err() != nil {
		return nil, dto.Paging{}, rows.Err()
	}

	var count int
	row := a.db.QueryRow("select count(id) from asset")
	if err := row.Scan(&count); err != nil {
		return nil, dto.Paging{}, err
	}

	paging := dto.Paging{
		Page:       payload.Page,
		Size:       payload.Size,
		TotalRows:  count,
		TotalPages: int(math.Ceil(float64(count) / float64(payload.Size))), // (totalrow / size)
	}

	return assets, paging, nil
}

// UpdateAmount implements AssetRepository.
func (a *assetRepository) UpdateAvailable(id string, amount int) error {
	query := "update asset set available = $2 where id = $1"

	_, err := a.db.Exec(query, id, amount)
	if err != nil {
		return err
	}

	return nil
}

// FindByName implements AssetRepository.
func (a *assetRepository) FindByName(name string) ([]model.Asset, error) {
	query := `select a.id, a.name, a.available, a.status, a.entry_date, a.img_url, a.total, c.id, c.name, at.id, at.name
			from asset as a 
			left join category as c on c.id = a.id_category
			left join asset_type as at on at.id = a.id_asset_type
			where a.name ilike $1`

	rows, err := a.db.Query(query, "%"+name+"%")
	if err != nil {
		return nil, err
	}

	var assets []model.Asset
	for rows.Next() {
		var asset model.Asset
		rows.Scan(&asset.Id, &asset.Name, &asset.Available, &asset.Status, &asset.EntryDate, &asset.ImgUrl, &asset.Total, &asset.Category.Id, &asset.Category.Name, &asset.AssetType.Id, &asset.AssetType.Name)
		assets = append(assets, asset)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return assets, nil

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

	query := `select a.id, a.name, a.available, a.status, a.entry_date, a.img_url, a.total, c.id, c.name, at.id, at.name
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
		rows.Scan(&asset.Id, &asset.Name, &asset.Available, &asset.Status, &asset.EntryDate, &asset.ImgUrl, &asset.Total, &asset.Category.Id, &asset.Category.Name, &asset.AssetType.Id, &asset.AssetType.Name)
		assets = append(assets, asset)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return assets, nil
}

// FindById implements AssetRepository.
func (a *assetRepository) FindById(id string) (model.Asset, error) {
	query := `select a.id, a.name, a.available, a.status, a.entry_date, a.img_url, a.total, c.id, c.name, at.id, at.name
			from asset as a 
			left join category as c on c.id = a.id_category
			left join asset_type as at on at.id = a.id_asset_type
			where a.id = $1`

	row := a.db.QueryRow(query, id)
	var asset model.Asset
	err := row.Scan(&asset.Id, &asset.Name, &asset.Available, &asset.Status, &asset.EntryDate, &asset.ImgUrl, &asset.Total, &asset.Category.Id, &asset.Category.Name, &asset.AssetType.Id, &asset.AssetType.Name)
	if err != nil {
		return model.Asset{}, err
	}

	return asset, nil
}

// Save implements AssetRepository.
func (a *assetRepository) Save(asset model.AssetRequest) error {
	query := "insert into asset(id, id_category, id_asset_type, name, available, status, entry_date, img_url, total) values($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	_, err := a.db.Exec(query, asset.Id, asset.CategoryId, asset.AssetTypeId, asset.Name, asset.Available, asset.Status, asset.EntryDate, asset.ImgUrl, asset.Total)
	if err != nil {
		return err
	}

	return nil
}

// Update implements AssetRepository.
func (a *assetRepository) Update(asset model.AssetRequest) error {
	query := `update asset set id_category = $2, id_asset_type = $3, name = $4, available = $5, status = $6, img_url = $7, total = $8 where id = $1`

	_, err := a.db.Exec(query, asset.Id, asset.CategoryId, asset.AssetTypeId, asset.Name, asset.Available, asset.Status, asset.ImgUrl, asset.Total)
	if err != nil {
		return err
	}

	return nil
}

func NewAssetRepository(db *sql.DB) AssetRepository {
	return &assetRepository{
		db: db,
	}
}
