package usecase

import (
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/repository"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type UserDetailsUsecase interface {
	NewUserDetails(udetails model.UserDetails) error
}

type userDetailUsecase struct {
	udetailsRepo repository.UserDetailsRepository
}

func (u userDetailUsecase) NewUserDetails(udetails model.UserDetails) error {
	//TODO implement me

	//business logic happen here
	//implement validate struct
	val := validator.New()
	if err := val.Struct(udetails); err != nil {
		return fmt.Errorf("Missing require info")
	}

	if err := u.udetailsRepo.CreateNewUserDetails(udetails); err != nil {
		return err
	}

	return nil
}

func NewUserDetailsUsecase(udetailsRepo repository.UserDetailsRepository) UserDetailsUsecase {
	return &userDetailUsecase{
		udetailsRepo: udetailsRepo,
	}
}
