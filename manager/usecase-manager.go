package manager

import "final-project-enigma-clean/usecase"

type UsecaseManager interface {
	UserUsecase() usecase.UserCredentialUsecase
	TypeAssetUseCase() usecase.TypeAssetUseCase
	StaffUseCase() usecase.StaffUseCase
	AssetUsecase() usecase.AssetUsecase
	CategoryUsecase() usecase.CategoryUsecase
	ManageAssetUsecase() usecase.ManageAssetUsecase
}

type usecaseManager struct {
	rm RepoManager
}

// ManageAssetUsecase implements UsecaseManager.
func (u *usecaseManager) ManageAssetUsecase() usecase.ManageAssetUsecase {
	return usecase.NewManageAssetUsecase(u.rm.ManageAssetRepo(), u.StaffUseCase(), u.AssetUsecase())
}

// StaffUseCase implements UsecaseManager.
func (u *usecaseManager) StaffUseCase() usecase.StaffUseCase {
	return usecase.NewStaffUseCase(u.rm.StaffRepo())
}

// CategoryUsecase implements UsecaseManager.
func (u *usecaseManager) CategoryUsecase() usecase.CategoryUsecase {
	return usecase.NewCategoryUseCase(u.rm.CategoryRepo())
}

// AssetUsecase implements UsecaseManager.
func (u *usecaseManager) AssetUsecase() usecase.AssetUsecase {
	return usecase.NewAssetUsecase(u.rm.AssetRepo(), u.TypeAssetUseCase(), u.CategoryUsecase())
}

// TypeAssetUseCase implements UsecaseManager.
func (u *usecaseManager) TypeAssetUseCase() usecase.TypeAssetUseCase {
	return usecase.NewTypeAssetUseCase(u.rm.TypeAssetRepo())
}

func (u *usecaseManager) UserUsecase() usecase.UserCredentialUsecase {
	//TODO implement me
	return usecase.NewUserCredentialUsecase(u.rm.UserRepo())
}

func NewUsecaseManager(rm RepoManager) UsecaseManager {
	return &usecaseManager{
		rm: rm,
	}
}
