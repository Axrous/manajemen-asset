package manager

import "final-project-enigma-clean/repository"

type RepoManager interface {
	UserRepo() repository.UserCredentialsRepository
	TypeAssetRepo() repository.TypeAssetRepository
	StaffRepo() repository.StaffRepository
	AssetRepo() repository.AssetRepository
	CategoryRepo() repository.CategoryRepository
	ManageAssetRepo() repository.ManageAssetRepository
}

type repoManager struct {
	im InfraManager
}

// ManageAssetRepo implements RepoManager.
func (r *repoManager) ManageAssetRepo() repository.ManageAssetRepository {
	return repository.NewManageAssetRepository(r.im.Connect())
}

// StaffRepo implements RepoManager.
func (r *repoManager) StaffRepo() repository.StaffRepository {
	return repository.NewStaffRepository(r.im.Connect())
}

// CategoryRepo implements RepoManager.
func (r *repoManager) CategoryRepo() repository.CategoryRepository {
	return repository.NewCategoryRepository(r.im.Connect())
}

// AssetRepo implements RepoManager.
func (r *repoManager) AssetRepo() repository.AssetRepository {
	return repository.NewAssetRepository(r.im.Connect())

}

// TypeAssetRepo implements RepoManager.
func (r *repoManager) TypeAssetRepo() repository.TypeAssetRepository {
	return repository.NewTypeAssetRepository(r.im.Connect())
}

func (r *repoManager) UserRepo() repository.UserCredentialsRepository {
	//TODO implement me
	return repository.NewUserDetailsRepository(r.im.Connect())
}

func NewRepoManager(im InfraManager) RepoManager {
	return &repoManager{
		im: im,
	}
}
