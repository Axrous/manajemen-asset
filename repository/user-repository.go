package repository

import (
	"database/sql"
	"final-project-enigma-clean/model"
	"fmt"
)

type UserRepository interface {
	RegisterUser(request model.UserRegisterRequest) error
	LoginUser(logreq model.UserLoginRequest) (string, error)
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) LoginUser(logreq model.UserLoginRequest) (string, error) {
	//TODO implement me

	var hashedPass string
	query := "select password from users where email = $1"
	err := u.db.QueryRow(query, logreq.Email).Scan(&hashedPass)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("Invalid credential")
		}
	}
	return hashedPass, err

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
