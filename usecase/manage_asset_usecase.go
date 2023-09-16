package usecase

import (
	"database/sql"
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"
	"final-project-enigma-clean/repository"
	"final-project-enigma-clean/util/helper"
	"fmt"
	"time"
)

type ManageAssetUsecase interface {
	CreateTransaction(payload dto.ManageAssetRequest) error
	ShowAllAsset() ([]model.ManageAsset, error)
	FindByTransactionID(id string) ([]model.ManageDetailAsset, error)
	DownloadAssets() ([]byte, error)
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

	err = m.repo.CreateTransaction(payload)
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

func (m *manageAssetUsecase) FindByTransactionID(id string) ([]model.ManageDetailAsset, error) {
	//TODO implement me
	if id == "" {
		return nil, fmt.Errorf("ID is required")
	}

	detailAssets, err := m.repo.FindAllByTransId(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Transaction not found")
		}
		return nil, fmt.Errorf("Failed to fetch transaction details: %v", err)
	}
	return detailAssets, nil
}

func (m *manageAssetUsecase) DownloadAssets() ([]byte, error) {
	//TODO implement me
	assets, err := m.repo.FindAllTransaction()
	if err != nil {
		return nil, fmt.Errorf("Failed to find assets %v", err.Error())
	}

	//convert data to csv
	csvData, err := helper.ConvertToCSVForAssets(assets)
	return csvData, nil
}

func NewManageAssetUsecase(repo repository.ManageAssetRepository, staffUC StaffUseCase, assetUC AssetUsecase) ManageAssetUsecase {
	return &manageAssetUsecase{
		repo:    repo,
		staffUC: staffUC,
		assetUC: assetUC,
	}
}
