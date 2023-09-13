package usecase

import (
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"
	"final-project-enigma-clean/repository"
	"fmt"
	"time"
)

type ManageAssetUsecase interface {
	CreateTransaction(payload dto.ManageAssetRequest) error
	ShowAllAsset() ([]model.ManageAsset, error)
}

type manageAssetUsecase struct {
	repo    repository.ManageAssetRepository
	staffUC StaffUseCase
	assetUC AssetUsecase
}

// CreateTransaction implements ManageAssetUsecase.
func (m *manageAssetUsecase) CreateTransaction(payload dto.ManageAssetRequest) error {
	if payload.NikStaff == "" {
		return fmt.Errorf("nik staff cannot empty")
	}

	var newManageDetail []dto.ManageAssetDetailRequest
	for _, detail := range payload.ManageAssetDetailReq {
		if detail.IdAsset == "" {
			return fmt.Errorf("id asset cannot empty")
		}

		if detail.Status == "" {
			return fmt.Errorf("status cannot empty")
		}

		if detail.TotalItem < 0 {
			return fmt.Errorf("total item must equal than 0")
		}

		_, err := m.assetUC.FindById(detail.IdAsset)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
		newManageDetail = append(newManageDetail, detail)
	}

	_, err := m.staffUC.FindById(payload.NikStaff)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	payload.ManageAssetDetailReq = newManageDetail
	payload.SubmisstionDate = time.Now()
	payload.ReturnDate = payload.SubmisstionDate.AddDate(0, 0, payload.Duration)
	err = m.repo.CreateTransaction(payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

func (m *manageAssetUsecase) ShowAllAsset() ([]model.ManageAsset, error) {
	//TODO implement me
	return m.repo.FindAllTransaction()
}

func NewManageAssetUsecase(repo repository.ManageAssetRepository, staffUC StaffUseCase, assetUC AssetUsecase) ManageAssetUsecase {
	return &manageAssetUsecase{
		repo:    repo,
		staffUC: staffUC,
		assetUC: assetUC,
	}
}
