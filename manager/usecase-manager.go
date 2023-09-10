package manager

import "final-project-enigma-clean/usecase"

type UsecaseManager interface {
	UDetailsUC() usecase.UserDetailsUsecase
}

type usecaseManager struct {
	rm RepoManager
}

func (u usecaseManager) UDetailsUC() usecase.UserDetailsUsecase {
	//TODO implement me
	return usecase.NewUserDetailsUsecase(u.rm.UserDetailsRepo())
}

func NewUsecaseManager(rm RepoManager) UsecaseManager {
	return &usecaseManager{
		rm: rm,
	}
}
