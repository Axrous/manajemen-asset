package repomock

import (
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"

	"github.com/stretchr/testify/mock"
)

type StaffRepoMock struct {
	mock.Mock
}

// Save implements repository.StaffRepository.
func (s *StaffRepoMock) Save(payload model.Staff) error {
	return s.Called(payload).Error(0)
}

// FindByName implements repository.StaffRepository.
func (s *StaffRepoMock) FindByName(name string) ([]model.Staff, error) {
	args := s.Called(name)
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Staff), nil
}

// Paging implements repository.StaffRepository.
func (s *StaffRepoMock) Paging(payload dto.PageRequest) ([]model.Staff, dto.Paging, error) {
	args := s.Called(payload)

	// Extract the arguments and return values from the recorded call
	staffs := args.Get(0).([]model.Staff)
	paging := args.Get(1).(dto.Paging)
	err := args.Error(2)

	return staffs, paging, err
}

// FindById implements StaffRepository.
func (s *StaffRepoMock) FindById(id string) (model.Staff, error) {
	args := s.Called(id)
	if args.Get(1) != nil {
		return model.Staff{}, args.Error(1)
	}
	return args.Get(0).(model.Staff), nil
}

// Delete implements StaffRepository.
func (s *StaffRepoMock) Delete(id string) error {
	return s.Called(id).Error(0)
}

// FindAll implements StaffRepository.
func (s *StaffRepoMock) FindByAll() ([]model.Staff, error) {
	args := s.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Staff), nil
}

// Save implements StaffRepository.

// Update implements StaffRepository.
func (s *StaffRepoMock) Update(payload model.Staff) error {
	return s.Called(payload).Error(0)
}
