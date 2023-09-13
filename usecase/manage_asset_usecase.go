package usecase

import (
	"final-project-enigma-clean/model/dto"
	"final-project-enigma-clean/repository"
	"fmt"
)

type ManageAssetUsecase interface {
	CreateTransaction(payload dto.ManageAssetRequest) error
}

type manageAssetUsecase struct {
	repo repository.ManageAssetRepository
	staffUC StaffUseCase
	assetUC AssetUsecase
}

// CreateTransaction implements ManageAssetUsecase.
func (m *manageAssetUsecase) CreateTransaction(payload dto.ManageAssetRequest) error {
	if payload.NikStaff == "" {
		return fmt.Errorf("nik staff cannot empty")
	}

	for _, v := range payload.ManageAssetDetailReq {
		if v.IdAsset == "" {
			return fmt.Errorf("id asset cannot empty")
		}

		if v.Status == "" {
			return fmt.Errorf("status cannot empty")
		}
		
		if v.TotalItem < 0 {
			return fmt.Errorf("total item must equal than 0")
		}
		
		_, err := m.assetUC.FindById(v.IdAsset)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
	}

	_, err := m.staffUC.FindById(payload.NikStaff)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	err = m.repo.CreateTransaksi(payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

func NewManageAssetUsecase(repo repository.ManageAssetRepository, staffUC StaffUseCase, assetUC AssetUsecase) ManageAssetUsecase {
	return &manageAssetUsecase{
		repo:    repo,
		staffUC: staffUC,
	}
}
