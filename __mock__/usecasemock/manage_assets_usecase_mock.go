package usecasemock

import (
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"

	"github.com/stretchr/testify/mock"
)

type ManageAssetsMock struct {
	mock.Mock
}

// FindByTransactionID implements usecase.ManageAssetUsecase.
func (*ManageAssetsMock) FindByTransactionID(id string) ([]model.ManageAsset, error) {
	panic("unimplemented")
}

// FindTransactionByName implements usecase.ManageAssetUsecase.
func (*ManageAssetsMock) FindTransactionByName(name string) ([]model.ManageAsset, error) {
	panic("unimplemented")
}

func (m ManageAssetsMock) CreateTransaction(payload dto.ManageAssetRequest) error {
	//TODO implement me
	panic("implement me")
}

func (m ManageAssetsMock) ShowAllAsset() ([]model.ManageAsset, error) {
	//TODO implement me
	panic("implement me")
}

func (m ManageAssetsMock) DownloadAssets() ([]byte, error) {
	//TODO implement me
	panic("implement me")
}
