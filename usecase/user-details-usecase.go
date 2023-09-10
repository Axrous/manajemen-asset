package usecase

import (
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/repository"
	"fmt"
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
	if udetails.ID == "" {
		return fmt.Errorf("ID is required")
	} else if udetails.Name == "" {
		return fmt.Errorf("Name is required")
	} else if udetails.PhoneNumber == "" {
		return fmt.Errorf("Phone Number is required")
	} else if udetails.Address == "" {
		return fmt.Errorf("Address is required")
	} else if udetails.BirthDate.IsZero() {
		return fmt.Errorf("Birth Date is required")
	} else if udetails.ImgUrl == "" {
		return fmt.Errorf("Image URL is required")
	}

	return nil
}

func NewUserDetailsUsecase(udetailsRepo repository.UserDetailsRepository) UserDetailsUsecase {
	return &userDetailUsecase{
		udetailsRepo: udetailsRepo,
	}
}
