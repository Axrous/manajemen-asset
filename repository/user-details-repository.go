package repository

import (
	"database/sql"
	"final-project-enigma-clean/model"
	"fmt"
)

type UserDetailsRepository interface {
	CreateNewUserDetails(udetails model.UserDetails) error
}

type userDetailsRepository struct {
	db *sql.DB
}

func (u *userDetailsRepository) CreateNewUserDetails(udetails model.UserDetails) error {
	//TODO implement me
	query := "insert into user_details (id,user_id,name,phone_number,address,birth_date,img_url) values ($1, $2, $3, $4, $5, $6, $7)"
	_, err := u.db.Exec(query, udetails.ID, udetails.UserID, udetails.Name, udetails.PhoneNumber, udetails.Address, udetails.BirthDate, udetails.ImgUrl)
	if err != nil {
		return fmt.Errorf("Failed to query %v", err.Error())
	}
	return nil
}

func NewUserDetailsRepository(db *sql.DB) UserDetailsRepository {
	return &userDetailsRepository{
		db: db,
	}
}
