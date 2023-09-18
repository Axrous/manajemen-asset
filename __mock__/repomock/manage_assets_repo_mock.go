package repomock


type ManageAssetsRepoMock struct {
	mock.Mock
}
import (
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"

	"github.com/stretchr/testify/mock"
)

type ManageAssetRepoMock struct {
	mock.Mock
}

func (m *ManageAssetRepoMock) FindByNameTransaction(name string) ([]model.ManageAsset, []model.ManageDetailAsset, error) {
	args := m.Called(name)
	if args.Get(2) != nil {
		return nil, nil, args.Error(2)
	}
	return args.Get(0).([]model.ManageAsset), args.Get(1).([]model.ManageDetailAsset), nil
}

// FindAllByTransId implements ManageAssetRepository.
func (m *ManageAssetRepoMock) FindAllByTransId(id string) ([]model.ManageAsset, []model.ManageDetailAsset, error) {
	args := m.Called(id)
	if args.Get(2) != nil {
		return nil, nil, args.Error(2)
	}
	return args.Get(0).([]model.ManageAsset), args.Get(1).([]model.ManageDetailAsset), nil
}

// FindAll implements ManageAssetRepository.
func (m *ManageAssetRepoMock) FindAllTransaction() ([]model.ManageAsset, error) {
	args := m.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.ManageAsset), nil
}

// CreateTransaksi implements ManageAssetRepository.
func (m *ManageAssetRepoMock) CreateTransaction(payload dto.ManageAssetRequest) error {
	return m.Called(payload).Error(0)
}
