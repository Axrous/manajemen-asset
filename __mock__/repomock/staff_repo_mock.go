package repomock

import (
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"

	"github.com/stretchr/testify/mock"
)

type StaffRepoMock struct {
	mock.Mock
}

// FindByName implements repository.TypeAssetRepository.
func (t *StaffRepoMock) FindByName(name string) ([]model.Staff, error) {
	args := t.Called(name)
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Staff), nil
}

// Paging implements repository.TypeAssetRepository.
func (t *StaffRepoMock) Paging(payload dto.PageRequest) ([]model.Staff, dto.Paging, error) {
	args := t.Called(payload)

	// Extract the arguments and return values from the recorded call
	staffs := args.Get(0).([]model.Staff)
	paging := args.Get(1).(dto.Paging)
	err := args.Error(2)

	return staffs, paging, err
}

// FindById implements categoryRepository.
func (t *StaffRepoMock) FindById(id string) (model.Staff, error) {
	args := t.Called(id)
	if args.Get(1) != nil {
		return model.Staff{}, args.Error(1)
	}
	return args.Get(0).(model.Staff), nil
}

// Delete implements categoryRepository.
func (t *StaffRepoMock) Delete(id string) error {
	return t.Called(id).Error(0)
}

// FindAll implements categoryRepository.
func (t *StaffRepoMock) FindAll() ([]model.Staff, error) {
	args := t.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Staff), nil
}

// Save implements categoryRepository.
func (t *StaffRepoMock) Save(payload model.Staff) error {
	return t.Called(payload).Error(0)
}

// Update implements categoryRepository.
func (t *StaffRepoMock) Update(payload model.Staff) error {
	return t.Called(payload).Error(0)
}
