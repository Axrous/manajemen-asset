package manager

import "final-project-enigma-clean/repository"

type RepoManager interface {
	UserRepo() repository.UserRepository
	UserDetailsRepo() repository.UserDetailsRepository
}

type repoManager struct {
	im InfraManager
}

func (r repoManager) UserRepo() repository.UserRepository {
	//TODO implement me
	return repository.NewUserRepository(r.im.Connect())
}

func (r repoManager) UserDetailsRepo() repository.UserDetailsRepository {
	//TODO implement me
	return repository.NewUserDetailsRepository(r.im.Connect())
}

func NewRepoManager(im InfraManager) RepoManager {
	return &repoManager{
		im: im,
	}
}
