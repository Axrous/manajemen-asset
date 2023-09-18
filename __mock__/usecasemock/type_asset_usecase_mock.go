package usecasemock

import (
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"

	"github.com/stretchr/testify/mock"
)

type TypeAssetUsecaseMock struct {
	mock.Mock
}

// FindByName implements usecase.TypeAssetUseCase.
func (t *TypeAssetUsecaseMock) FindByName(name string) ([]model.TypeAsset, error) {
	args := t.Called(name)
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]model.TypeAsset), nil
}

// Paging implements usecase.TypeAssetUseCase.
func (t *TypeAssetUsecaseMock) Paging(payload dto.PageRequest) ([]model.TypeAsset, dto.Paging, error) {
	args := t.Called(payload)

	// Extract the arguments and return values from the recorded call
	typeAssets := args.Get(0).([]model.TypeAsset)
	paging := args.Get(1).(dto.Paging)
	err := args.Error(2)

	return typeAssets, paging, err
}

func (t *TypeAssetUsecaseMock) FindById(id string) (model.TypeAsset, error) {
	args := t.Called(id)
	if args.Get(1) != nil {
		return model.TypeAsset{}, args.Error(1)
	}

	return args.Get(0).(model.TypeAsset), nil

}

// CreateNew implements CategoryUseCase.
func (t *TypeAssetUsecaseMock) CreateNew(payload model.TypeAsset) error {
	return t.Called(payload).Error(0)
}

// Delete implements CategoryUseCase.
func (t *TypeAssetUsecaseMock) Delete(id string) error {
	// panic("implement me")
	return t.Called(id).Error(0)
}

// FindAll implements CategoryUseCase.
func (t *TypeAssetUsecaseMock) FindAll() ([]model.TypeAsset, error) {
	args := t.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]model.TypeAsset), nil
}

// Update implements CategoryUseCase.
func (t *TypeAssetUsecaseMock) Update(payload model.TypeAsset) error {
	// panic("implement me")
	return t.Called(payload).Error(0)
}
