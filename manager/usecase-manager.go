package manager

import "final-project-enigma-clean/usecase"

type UsecaseManager interface {
	UserUsecase() usecase.UserCredentialUsecase
	TypeAssetUseCase() usecase.TypeAssetUseCase
<<<<<<< HEAD
	StaffUseCase() usecase.StaffUseCase
=======
	CategoryUsecase() usecase.CategoryUsecase
>>>>>>> ceca9a7
}

type usecaseManager struct {
	rm RepoManager
}

<<<<<<< HEAD
// StaffUseCase implements UsecaseManager.
func (u *usecaseManager) StaffUseCase() usecase.StaffUseCase {
	return usecase.NewStaffUseCase(u.rm.StaffRepo())
=======
// CategoryUsecase implements UsecaseManager.
func (u *usecaseManager) CategoryUsecase() usecase.CategoryUsecase {
	return usecase.NewCategoryUseCase(u.rm.CategoryRepo())
>>>>>>> ceca9a7
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
