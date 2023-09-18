package usecasemock

import (
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"

	"github.com/stretchr/testify/mock"
)

type ManageAssetsMock struct {
	mock.Mock
}

func (m *ManageAssetsMock) Create(payload model.AssetRequest) error {
	//TODO implement me
	return m.Called(payload).Error(0)
}

func (m *ManageAssetsMock) FindAll() ([]model.Asset, error) {
	//TODO implement me
	panic("implement me")
}

func (m *ManageAssetsMock) FindById(id string) (model.Asset, error) {
	//TODO implement me
	args := m.Called(id)
	if args.Get(1) != nil {
		return model.Asset{}, nil
	}
	return args.Get(0).(model.Asset), nil
}

func (m *ManageAssetsMock) Update(payload model.AssetRequest) error {
	//TODO implement me
	panic("implement me")
}

func (m *ManageAssetsMock) UpdateAvailable(id string, amount int) error {
	//TODO implement me
	panic("implement me")
}

func (m *ManageAssetsMock) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func (m *ManageAssetsMock) FindByName(name string) ([]model.Asset, error) {
	//TODO implement me
	panic("implement me")
}

func (m *ManageAssetsMock) Paging(payload dto.PageRequest) ([]model.Asset, dto.Paging, error) {
	//TODO implement me
	panic("implement me")
}

func (m *ManageAssetsMock) FindByTransactionID(id string) ([]model.ManageAsset, error) {
	//TODO implement me
	//TODO implement me
	args := m.Called(id)
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.ManageAsset), nil
}

func (m *ManageAssetsMock) FindTransactionByName(name string) ([]model.ManageAsset, error) {
	//TODO implement me
	args := m.Called(name)
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]model.ManageAsset), nil
}

func (m *ManageAssetsMock) CreateTransaction(payload dto.ManageAssetRequest) error {
	//TODO implement me
	return m.Called(payload).Error(0)
}

func (m *ManageAssetsMock) ShowAllAsset() ([]model.ManageAsset, error) {
	//TODO implement me
	args := m.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]model.ManageAsset), nil
}

func (m *ManageAssetsMock) DownloadAssets() ([]byte, error) {
	//TODO implement me
	panic("implement me")
}
