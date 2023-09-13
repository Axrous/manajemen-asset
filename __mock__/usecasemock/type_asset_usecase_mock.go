package usecasemock

import (
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"

	"github.com/stretchr/testify/mock"
)

type TypeAssetUsecaseMock struct {
	mock.Mock
}

// CreateNew implements usecase.TypeAssetUseCase.
func (t *TypeAssetUsecaseMock) CreateNew(payload model.TypeAsset) error {
	panic("unimplemented")
}

// Delete implements usecase.TypeAssetUseCase.
func (t *TypeAssetUsecaseMock) Delete(id string) error {
	panic("unimplemented")
}

// FindAll implements usecase.TypeAssetUseCase.
func (t *TypeAssetUsecaseMock) FindAll() ([]model.TypeAsset, error) {
	panic("unimplemented")
}

// FindByName implements usecase.TypeAssetUseCase.
func (t *TypeAssetUsecaseMock) FindByName(name string) ([]model.TypeAsset, error) {
	panic("unimplemented")
}

// Paging implements usecase.TypeAssetUseCase.
func (t *TypeAssetUsecaseMock) Paging(payload dto.PageRequest) ([]model.TypeAsset, dto.Paging, error) {
	panic("unimplemented")
}

// Update implements usecase.TypeAssetUseCase.
func (t *TypeAssetUsecaseMock) Update(payload model.TypeAsset) error {
	panic("unimplemented")
}

func (t *TypeAssetUsecaseMock) FindById(id string) (model.TypeAsset, error) {

	args := t.Called(id)
	if args.Get(1) != nil {
		return model.TypeAsset{}, args.Error(1)
	}

	return args.Get(0).(model.TypeAsset), nil
}