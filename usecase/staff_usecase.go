package usecase

import (
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"
	"final-project-enigma-clean/repository"
	"final-project-enigma-clean/util/helper"
	"fmt"
)

type StaffUseCase interface {
	CreateNew(payload model.Staff) error
	FindByName(name string) ([]model.Staff, error)
	FindById(nik_staff string) (model.Staff, error)
	FindByAll() ([]model.Staff, error)
	Update(payload model.Staff) error
	Delete(nik_staff string) error
	Paging(payload dto.PageRequest) ([]model.Staff, dto.Paging, error)
	DownloadAllStaff() ([]byte, error)
}

type staffUseCase struct {
	repo repository.StaffRepository
}

// FindById implements StaffUseCase.
func (s *staffUseCase) FindById(nik_staff string) (model.Staff, error) {
	staff, err := s.repo.FindById(nik_staff)
	if err != nil {
		return model.Staff{}, fmt.Errorf("staff not found")
	}
	return staff, nil

}

// CreateNew implements StaffUseCase.
func (s *staffUseCase) CreateNew(payload model.Staff) error {
	if payload.Nik_Staff == "" {
		return fmt.Errorf("nik staff is required")
	}
	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}
	if len(payload.Phone_number) < 10 || len(payload.Phone_number) > 15 {
		return fmt.Errorf("phone number must be between 10 and 15 characters")
	}
	if payload.Address == "" {
		return fmt.Errorf("address is required")
	}
	if payload.Divisi == "" {
		return fmt.Errorf("divisi is required")
	}
	err := s.repo.Save(payload)
	if err != nil {
		return fmt.Errorf("failed to create new staff: %v", err)
	}
	return nil
}

// Delete implements StaffUseCase.
func (s *staffUseCase) Delete(nik_staff string) error {
	staff, err := s.FindById(nik_staff)
	if err != nil {
		return err
	}
	err = s.repo.Delete(staff.Nik_Staff)
	if err != nil {
		return fmt.Errorf("failed to delete staff: %v", err)
	}
	return nil
}

// FindAll implements StaffUseCase.
func (s *staffUseCase) FindByAll() ([]model.Staff, error) {
	staff, err := s.repo.FindByAll()
	if err != nil {
		return nil, fmt.Errorf("failed to find by all staff: %v", err)
	}
	return staff, nil
}

// FindByName implements StaffUseCase.
func (s *staffUseCase) FindByName(name string) ([]model.Staff, error) {
	staff, err := s.repo.FindByName(name)
	if err != nil {
		return nil, fmt.Errorf("name staff not found: %v", err)
	}
	return staff, nil

}

// Paging implements StaffUseCase.
func (s *staffUseCase) Paging(payload dto.PageRequest) ([]model.Staff, dto.Paging, error) {
	return s.repo.Paging(payload)
}

// Update implements StaffUseCase.
func (s *staffUseCase) Update(payload model.Staff) error {
	if payload.Nik_Staff == "" {
		return fmt.Errorf("nik staff is required")
	}
	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}
	if len(payload.Phone_number) < 10 || len(payload.Phone_number) > 15 {
		return fmt.Errorf("phone number must be between 10 and 15 characters")
	}
	if payload.Address == "" {
		return fmt.Errorf("address is required")
	}
	if payload.Divisi == "" {
		return fmt.Errorf("divisi is required")
	}
	_, err := s.FindById(payload.Nik_Staff)
	if err != nil {
		return err
	}
	err = s.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update staff: %v", err)
	}
	return nil
}

func (s *staffUseCase) DownloadAllStaff() ([]byte, error) {
	// Mengambil data staff dari repository
	staffs, err := s.repo.FindByAll()
	if err != nil {
		return nil, fmt.Errorf("failed to find staff: %v", err)
	}

	// Mengonversi data staff ke format CSV
	csvData, err := helper.ConvertToCSVForStaff(staffs)
	if err != nil {
		return nil, fmt.Errorf("failed to convert staff data to CSV: %v", err)
	}

	return csvData, nil
}

func NewStaffUseCase(repo repository.StaffRepository) StaffUseCase {
	return &staffUseCase{
		repo: repo,
	}
}
