package repository

import (
	"database/sql"
	"final-project-enigma-clean/model"
	"fmt"
)

type UserRepository interface {
	RegisterUser(request model.UserRegisterRequest) error
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) RegisterUser(request model.UserRegisterRequest) error {
	//TODO implement me

	//register logic
	query := "insert into users (id,email,password) values ($1, $2, $3)"
	_, err := u.db.Exec(query, request.ID, request.Email, request.Password)
	if err != nil {
		return fmt.Errorf("Failed to exec query %v", err.Error())
	}
	return nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
