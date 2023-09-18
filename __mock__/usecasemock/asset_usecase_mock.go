package usecasemock

import (
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"

	"github.com/stretchr/testify/mock"
)

type AssetUsecaseMock struct {
	mock.Mock
}

// Paging implements usecase.AssetUsecase.
func (a *AssetUsecaseMock) Paging(payload dto.PageRequest) ([]model.Asset, dto.Paging, error) {
	args := a.Called(payload)
	if args.Get(2) != nil {
		return nil, dto.Paging{}, args.Error(2)
	}

	return args.Get(0).([]model.Asset), args.Get(1).(dto.Paging), nil
}

// UpdateAvailable implements usecase.AssetUsecase.
func (a *AssetUsecaseMock) UpdateAvailable(id string, amount int) error {
	return a.Called(id, amount).Error(0)
}

// FindByName implements usecase.AssetUsecase.
func (a *AssetUsecaseMock) FindByName(name string) ([]model.Asset, error) {
	args := a.Called(name)
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]model.Asset), nil
}

// Create implements AssetUsecase.
func (a *AssetUsecaseMock) Create(payload model.AssetRequest) error {
	return a.Called(payload).Error(0)
}

// Delete implements AssetUsecase.
func (a *AssetUsecaseMock) Delete(id string) error {
	return a.Called(id).Error(0)
}

// FindAll implements AssetUsecase.
func (a *AssetUsecaseMock) FindAll() ([]model.Asset, error) {
	args := a.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]model.Asset), nil
}

func (a *AssetUsecaseMock) FindById(id string) (model.Asset, error) {
	args := a.Called(id)
	if args.Get(1) != nil {
		return model.Asset{}, args.Error(1)
	}

	return args.Get(0).(model.Asset), nil
}

// Update implements AssetUsecase.
func (a *AssetUsecaseMock) Update(payload model.AssetRequest) error {
	return a.Called(payload).Error(0)
}
