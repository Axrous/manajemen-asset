package manager

import "final-project-enigma-clean/repository"

type RepoManager interface {
	UserRepo() repository.UserCredentialsRepository
	TypeAssetRepo() repository.TypeAssetRepository
<<<<<<< HEAD
	StaffRepo() repository.StaffRepository
	AssetRepo() repository.AssetRepository
=======
	CategoryRepo() repository.CategoryRepository
>>>>>>> ceca9a7
}

type repoManager struct {
	im InfraManager
}

<<<<<<< HEAD
// StaffRepo implements RepoManager.
func (r *repoManager) StaffRepo() repository.StaffRepository {
	return repository.NewStaffRepository(r.im.Connect())
=======
// CategoryRepo implements RepoManager.
func (r *repoManager) CategoryRepo() repository.CategoryRepository {
	return repository.NewCategoryRepository(r.im.Connect())
>>>>>>> ceca9a7
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
