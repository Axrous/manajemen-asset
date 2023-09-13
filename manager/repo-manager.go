package manager

import "final-project-enigma-clean/repository"

type RepoManager interface {
	UserRepo() repository.UserCredentialsRepository
	AssetRepo() repository.AssetRepository
	TypeAssetRepo() repository.TypeAssetRepository
	CategoryRepo() repository.CategoryRepository
}

type repoManager struct {
	im InfraManager
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

func (r repoManager) UserRepo() repository.UserCredentialsRepository {
	//TODO implement me
	return repository.NewUserDetailsRepository(r.im.Connect())
}

func NewRepoManager(im InfraManager) RepoManager {
	return &repoManager{
		im: im,
	}
}
