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
	return c.Called(payload).Error(0)
}

// Delete implements CategoryUseCase.
func (c *CategoryUsecaseMock) Delete(id string) error {
	// panic("implement me")
	return c.Called(id).Error(0)
}

// FindAll implements CategoryUseCase.
func (c *CategoryUsecaseMock) FindAll() ([]model.Category, error) {
	args := c.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]model.Category), nil
}

// Update implements CategoryUseCase.
func (c *CategoryUsecaseMock) Update(payload model.Category) error {
	// panic("implement me")
	return c.Called(payload).Error(0)
}
