package usecase

import (
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/repository"
	"final-project-enigma-clean/util/helper"
	"fmt"
	"time"
)

type AssetUsecase interface {
	Create(payload model.AssetRequest) error
	FindAll() ([]model.Asset, error)
	FindById(id string) (model.Asset, error)
	Update(payload model.AssetRequest) error
	Delete(id string) error
}

type assetUsecase struct {
	repo repository.AssetRepository
	//get category usecase
	//get asset type usecase
}

// Create implements AssetUsecase.
func (a *assetUsecase) Create(payload model.AssetRequest) error {
	if payload.Name == "" {
		return fmt.Errorf("name cannot empty")
	}
	if payload.AssetTypeId == "" || payload.CategoryId == "" {
		return fmt.Errorf("asset type id or category id cannot empty")
	}
	if payload.Amount < 0 {
		return fmt.Errorf("amoun cannot negative number")
	}
	if payload.Status == "" {
		return fmt.Errorf("status cannot empty")
	}

	//implement category find by id
	//implement asset type find by id
	payload.ID = helper.GenerateUUID()
	payload.EntryDate = time.Now()
	err := a.repo.Save(payload)
	if err != nil {
		return fmt.Errorf("failed save asset %s", err)
	}

	return nil
}

// Delete implements AssetUsecase.
func (a *assetUsecase) Delete(id string) error {
	//find assert first
	_, err := a.FindById(id)
	if err != nil {
		return err
	}

	err = a.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete asset, %s", err)
	}

	return nil
}

// FindAll implements AssetUsecase.
func (a *assetUsecase) FindAll() ([]model.Asset, error) {
	assets, err := a.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed get assets, %s", err)
	}
	return assets, nil
}

// FindById implements AssetUsecase.
func (a *assetUsecase) FindById(id string) (model.Asset, error) {
	asset, err := a.repo.FindById(id)
	if err != nil {
		return model.Asset{}, fmt.Errorf("asset by id:%s cannot found, err:%s", id, err)
	}

	return asset, nil
}

// Update implements AssetUsecase.
func (a *assetUsecase) Update(payload model.AssetRequest) error {
	if payload.Name == "" {
		return fmt.Errorf("name cannot empty")
	}
	if payload.AssetTypeId == "" || payload.CategoryId == "" {
		return fmt.Errorf("asset type id or category id cannot empty")
	}
	if payload.Amount < 0 {
		return fmt.Errorf("amoun cannot negative number")
	}
	if payload.Status == "" {
		return fmt.Errorf("status cannot empty")
	}

	//implement category find by id
	//implement asset type find by id

	_, err := a.FindById(payload.ID)
	if err != nil {
		return err
	}

	err = a.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed update asset %s", err)
	}

	return nil
}

func NewAssetUsecase(assetRepo repository.AssetRepository) AssetUsecase {
	return &assetUsecase{
		repo: assetRepo,
	}
}
