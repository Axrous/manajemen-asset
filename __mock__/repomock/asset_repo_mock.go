package repomock

import (
	"final-project-enigma-clean/model"

	"github.com/stretchr/testify/mock"
)

type AssetRepoMock struct {
	mock.Mock
}

// Delete implements AssetRepoMock.
func (a *AssetRepoMock) Delete(id string) error {
	return a.Called(id).Error(0)
}

// FindAll implements AssetRepoMock.
func (a *AssetRepoMock) FindAll() ([]model.Asset, error) {
	args := a.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Asset), nil
}

// FindById implements AssetRepoMock.
func (a *AssetRepoMock) FindById(id string) (model.Asset, error) {
	args := a.Called(id)
	if args.Get(1) != nil {
		return model.Asset{}, args.Error(1)
	}
	return args.Get(0).(model.Asset), nil
}

// Save implements AssetRepoMock.
func (a *AssetRepoMock) Save(asset model.AssetRequest) error {
	return a.Called(asset).Error(0)
}

// Update implements AssetRepoMock.
func (a *AssetRepoMock) Update(asset model.AssetRequest) error {
	return a.Called(asset).Error(0)
}
