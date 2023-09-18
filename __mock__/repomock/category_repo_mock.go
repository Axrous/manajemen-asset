package repomock

import (
	"final-project-enigma-clean/model"

	"github.com/stretchr/testify/mock"
)

type CategoryRepoMock struct {
	mock.Mock
}

// FindById implements categoryRepository.
func (c *CategoryRepoMock) FindById(id string) (model.Category, error) {
	args := c.Called(id)
	if args.Get(1) != nil {
		return model.Category{}, args.Error(1)
	}
	return args.Get(0).(model.Category), nil
}

// Delete implements categoryRepository.
func (c *CategoryRepoMock) Delete(id string) error {
	return c.Called(id).Error(0)
}

// FindAll implements categoryRepository.
func (c *CategoryRepoMock) FindAll() ([]model.Category, error) {
	args := c.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Category), nil
}

// Save implements categoryRepository.
func (c *CategoryRepoMock) Save(category model.Category) error {
	return c.Called(category).Error(0)
}

// Update implements categoryRepository.
func (c *CategoryRepoMock) Update(category model.Category) error {
	return c.Called(category).Error(0)
}
