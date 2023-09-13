package usecasemock

import (
	"final-project-enigma-clean/model"

	"github.com/stretchr/testify/mock"
)

type CategoryUsecaseMock struct {
	mock.Mock
}

func (c *CategoryUsecaseMock) FindById(id string) (model.Category, error) {
	args := c.Called(id)
	if args.Get(1) != nil {
		return model.Category{}, args.Error(1)
	}

	return args.Get(0).(model.Category), nil

}

// CreateNew implements CategoryUseCase.
func (c *CategoryUsecaseMock) CreateNew(payload model.Category) error {
	panic("implement me")
}

// Delete implements CategoryUseCase.
func (c *CategoryUsecaseMock) Delete(id string) error {
	panic("implement me")
}

// FindAll implements CategoryUseCase.
func (c *CategoryUsecaseMock) FindAll() ([]model.Category, error) {
	panic("implement me")
}

// Update implements CategoryUseCase.
func (c *CategoryUsecaseMock) Update(payload model.Category) error {
	panic("implement me")
}