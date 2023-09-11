package repository

import (
	"database/sql"
	"final-project-enigma-clean/model"
	"fmt"
)

type UserCredentialsRepository interface {
	UserRegister(user model.UserRegisterRequest) error
	UserLogin(user model.UserLoginRequest) (string, error)
	FindUserEmail(email string) (user model.UserLoginRequest, err error)
}

type userDetailsRepository struct {
	db *sql.DB
}

// user register
func (u userDetailsRepository) UserRegister(user model.UserRegisterRequest) error {
	//register logic

	user.IsActive = true

	query := "insert into user_credential (id,email,password,name,is_active) values ($1, $2, $3, $4, $5)"
	_, err := u.db.Exec(query, user.ID, user.Email, user.Password, user.Name, user.IsActive)
	if err != nil {
		return fmt.Errorf("Failed to exec query %v", err.Error())
	}

	return nil
}

// user login
func (u userDetailsRepository) UserLogin(user model.UserLoginRequest) (string, error) {
	//TODO implement me

	var hashedPassword string
	query := "select password from user_credential where email = $1"
	err := u.db.QueryRow(query, user.Email).Scan(&hashedPassword)
	if err != nil {
		//if email is not found
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("Invalid Credential %v", err.Error())
		}
		return "", err
	}
	return hashedPassword, nil
}

// find by email
func (u userDetailsRepository) FindUserEmail(email string) (model.UserLoginRequest, error) {
	// Query SQL untuk mencari pengguna berdasarkan email
	query := "SELECT id, email, password FROM user_credential WHERE email = $1"
	var user model.UserLoginRequest

	err := u.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		// Jika email tidak ditemukan
		if err == sql.ErrNoRows {
			return model.UserLoginRequest{}, fmt.Errorf("Email tidak ditemukan")
		}
		return model.UserLoginRequest{}, fmt.Errorf("Gagal menjalankan query: %v", err.Error())
	}

	return user, nil // Pengguna dengan email yang sesuai ditemukan
}

func NewUserDetailsRepository(db *sql.DB) UserCredentialsRepository {
	return &userDetailsRepository{
		db: db,
	}
}
