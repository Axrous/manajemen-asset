package repomock

import (
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"

	"github.com/stretchr/testify/mock"
)

type TypeAssetRepoMock struct {
	mock.Mock
}

// FindByName implements repository.TypeAssetRepository.
func (t *TypeAssetRepoMock) FindByName(name string) ([]model.TypeAsset, error) {
	args := t.Called(name)
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.TypeAsset), nil
}

// Paging implements repository.TypeAssetRepository.
func (t *TypeAssetRepoMock) Paging(payload dto.PageRequest) ([]model.TypeAsset, dto.Paging, error) {
	args := t.Called(payload)

	// Extract the arguments and return values from the recorded call
	typeAssets := args.Get(0).([]model.TypeAsset)
	paging := args.Get(1).(dto.Paging)
	err := args.Error(2)

	return typeAssets, paging, err
}

// FindById implements categoryRepository.
func (t *TypeAssetRepoMock) FindById(id string) (model.TypeAsset, error) {
	args := t.Called(id)
	if args.Get(1) != nil {
		return model.TypeAsset{}, args.Error(1)
	}
	return args.Get(0).(model.TypeAsset), nil
}

// Delete implements categoryRepository.
func (t *TypeAssetRepoMock) Delete(id string) error {
	return t.Called(id).Error(0)
}

// FindAll implements categoryRepository.
func (t *TypeAssetRepoMock) FindAll() ([]model.TypeAsset, error) {
	args := t.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.TypeAsset), nil
}

// Save implements categoryRepository.
func (t *TypeAssetRepoMock) Save(typeAsset model.TypeAsset) error {
	return t.Called(typeAsset).Error(0)
}

// Update implements categoryRepository.
func (t *TypeAssetRepoMock) Update(payload model.TypeAsset) error {
	return t.Called(payload).Error(0)
}
