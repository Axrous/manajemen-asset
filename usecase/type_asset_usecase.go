package usecase

import (
	"final-project-enigma-clean/exception"
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"
	"final-project-enigma-clean/repository"
	"final-project-enigma-clean/util/helper"
	"fmt"
)

type TypeAssetUseCase interface {
	CreateNew(payload model.TypeAsset) error
	FindByName(name string) ([]model.TypeAsset, error)
	FindById(id string) (model.TypeAsset, error)
	FindAll() ([]model.TypeAsset, error)
	Update(payload model.TypeAsset) error
	Delete(id string) error
	Paging(payload dto.PageRequest) ([]model.TypeAsset, dto.Paging, error)
}

type typeAssetUseCase struct {
	repo repository.TypeAssetRepository
}

// FindById implements TypeAssetUseCase.
func (t *typeAssetUseCase) FindById(id string) (model.TypeAsset, error) {
	typeAsset, err := t.repo.FindById(id)
	if err != nil {
		return model.TypeAsset{}, exception.BadRequestErr("type asset not found")
	}
	return typeAsset, nil

}

// CreateNew implements TypeAssetUseCase.
func (t *typeAssetUseCase) CreateNew(payload model.TypeAsset) error {
	if payload.Name == "" {
		return exception.BadRequestErr("name cannot Empty")
	}
	payload.Id = helper.GenerateUUID()
	err := t.repo.Save(payload)
	if err != nil {
		return fmt.Errorf("failed to create new type asset: %v", err)
	}
	return nil
}

// Delete implements TypeAssetUseCase.
func (t *typeAssetUseCase) Delete(id string) error {
	typeAsset, err := t.FindById(id)
	if err != nil {
		return err
	}
	err = t.repo.Delete(typeAsset.Id)
	if err != nil {
		return fmt.Errorf("failed to delete type asset: %v", err)
	}
	return nil
}

// FindAll implements TypeAssetUseCase.
func (t *typeAssetUseCase) FindAll() ([]model.TypeAsset, error) {
	typeAsset, err := t.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to find all type asset: %v", err)
	}
	return typeAsset, nil
}

// FindByName implements TypeAssetUseCase.
func (t *typeAssetUseCase) FindByName(name string) ([]model.TypeAsset, error) {
	typeAsset, err := t.repo.FindByName(name)
	if err != nil {
		return nil, exception.BadRequestErr("name type asset not found")
	}
	return typeAsset, nil

}

// Paging implements TypeAssetUseCase.
func (t *typeAssetUseCase) Paging(payload dto.PageRequest) ([]model.TypeAsset, dto.Paging, error) {
	return t.repo.Paging(payload)
}

// Update implements TypeAssetUseCase.
func (t *typeAssetUseCase) Update(payload model.TypeAsset) error {
	if payload.Name == "" {
		return exception.BadRequestErr("name cannot Empty")
	}
	_, err := t.FindById(payload.Id)
	if err != nil {
		return err
	}
	err = t.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update type asset: %v", err)
	}
	return nil
}

func NewTypeAssetUseCase(repo repository.TypeAssetRepository) TypeAssetUseCase {
	return &typeAssetUseCase{
		repo: repo,
	}
}
