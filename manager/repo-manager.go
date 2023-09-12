package manager

import "final-project-enigma-clean/repository"

type RepoManager interface {
	UserRepo() repository.UserCredentialsRepository
	TypeAssetRepo() repository.TypeAssetRepository
	StaffRepo() repository.StaffRepository
}

type repoManager struct {
	im InfraManager
}

// StaffRepo implements RepoManager.
func (r *repoManager) StaffRepo() repository.StaffRepository {
	return repository.NewStaffRepository(r.im.Connect())
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
