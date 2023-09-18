package usecasemock

import (
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"

	"github.com/stretchr/testify/mock"
)

type StaffUsecaseMock struct {
	mock.Mock
}

// FindByName implements usecase.StaffUseCase.
func (s *StaffUsecaseMock) FindByName(name string) ([]model.Staff, error) {
	args := s.Called(name)
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]model.Staff), nil
}

// Paging implements usecase.StaffUseCase.
func (s *StaffUsecaseMock) Paging(payload dto.PageRequest) ([]model.Staff, dto.Paging, error) {
	args := s.Called(payload)

	// Extract the arguments and return values from the recorded call
	staffs := args.Get(0).([]model.Staff)
	paging := args.Get(1).(dto.Paging)
	err := args.Error(2)

	return staffs, paging, err
}

func (s *StaffUsecaseMock) FindById(id string) (model.Staff, error) {
	args := s.Called(id)
	if args.Get(1) != nil {
		return model.Staff{}, args.Error(1)
	}

	return args.Get(0).(model.Staff), nil

}

// CreateNew implements StaffUseCase.
func (s *StaffUsecaseMock) CreateNew(payload model.Staff) error {
	return s.Called(payload).Error(0)
}

// Delete implements StaffUseCase.
func (s *StaffUsecaseMock) Delete(id string) error {
	// panic("implement me")
	return s.Called(id).Error(0)
}

// FindAll implements StaffUseCase.
func (s *StaffUsecaseMock) FindByAll() ([]model.Staff, error) {
	args := s.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]model.Staff), nil
}

// Update implements StaffUseCase.
func (s *StaffUsecaseMock) Update(payload model.Staff) error {
	// panic("implement me")
	return s.Called(payload).Error(0)
}
