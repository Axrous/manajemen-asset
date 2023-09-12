package usecasemock

import (
	"final-project-enigma-clean/model"

	"github.com/stretchr/testify/mock"
)

type TypeAssetUsecaseMock struct {
	mock.Mock
}

func (t *TypeAssetUsecaseMock) FindById(id string) (model.TypeAsset, error) {
	
	args := t.Called(id)
	if args.Get(1) != nil {
		return model.TypeAsset{}, args.Error(1)
	}

	return args.Get(0).(model.TypeAsset), nil

}