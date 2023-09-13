package manager

import "final-project-enigma-clean/usecase"

type UsecaseManager interface {
	UserUsecase() usecase.UserCredentialUsecase
	AssetUsecase() usecase.AssetUsecase
	TypeAssetUseCase() usecase.TypeAssetUseCase
	CategoryUsecase() usecase.CategoryUsecase
}

type usecaseManager struct {
	rm RepoManager
}

// CategoryUsecase implements UsecaseManager.
func (u *usecaseManager) CategoryUsecase() usecase.CategoryUsecase {
	return usecase.NewTypeCategoryUseCase(u.rm.CategoryRepo())
}

// AssetUsecase implements UsecaseManager.
func (u *usecaseManager) AssetUsecase() usecase.AssetUsecase {
	return usecase.NewAssetUsecase(u.rm.AssetRepo(), u.TypeAssetUseCase())
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
