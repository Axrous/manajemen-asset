package manager

import "final-project-enigma-clean/usecase"

type UsecaseManager interface {
	UserUsecase() usecase.UserCredentialUsecase
	TypeAssetUseCase() usecase.TypeAssetUseCase
	StaffUseCase() usecase.StaffUseCase
}

type usecaseManager struct {
	rm RepoManager
}

// StaffUseCase implements UsecaseManager.
func (u *usecaseManager) StaffUseCase() usecase.StaffUseCase {
	return usecase.NewStaffUseCase(u.rm.StaffRepo())
}

// TypeAssetUseCase implements UsecaseManager.
func (u *usecaseManager) TypeAssetUseCase() usecase.TypeAssetUseCase {
	return usecase.NewTypeAssetUseCase(u.rm.TypeAssetRepo())
}

func (u usecaseManager) UserUsecase() usecase.UserCredentialUsecase {
	//TODO implement me
	return usecase.NewUserCredentialUsecase(u.rm.UserRepo())
}

func NewUsecaseManager(rm RepoManager) UsecaseManager {
	return &usecaseManager{
		rm: rm,
	}
}
