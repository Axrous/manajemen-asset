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
	FindUserEmailPass(email string) (userPass model.ChangePasswordRequest, err error)
	ForgotPassword(email, newpass string) error
	GetUserPassword(email string) (string, error)
	CheckEmailExist(email string) bool
}

type userCredentialRepository struct {
	db *sql.DB
}

func (u userCredentialRepository) FindUserEmailPass(email string) (userPass model.ChangePasswordRequest, err error) {
	// Query SQL untuk mencari pengguna berdasarkan email
	query := "SELECT id, email, password FROM user_credential WHERE email = $1"
	var user model.ChangePasswordRequest

	err = u.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.OldPassword)
	if err != nil {
		// if email is not found
		if err == sql.ErrNoRows {
			return model.ChangePasswordRequest{}, fmt.Errorf("Invalid Credentials")
		}
		return model.ChangePasswordRequest{}, fmt.Errorf("Failed to run query: %v", err.Error())
	}

	return user, nil
}

// user register
func (u userCredentialRepository) UserRegister(user model.UserRegisterRequest) error {
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
func (u userCredentialRepository) UserLogin(user model.UserLoginRequest) (string, error) {
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
func (u userCredentialRepository) FindUserEmail(email string) (model.UserLoginRequest, error) {
	// Query SQL untuk mencari pengguna berdasarkan email
	query := "SELECT id, email, password FROM user_credential WHERE email = $1"
	var user model.UserLoginRequest

	err := u.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		// if email is not found
		if err == sql.ErrNoRows {
			return model.UserLoginRequest{}, fmt.Errorf("Invalid Credentials")
		}
		return model.UserLoginRequest{}, fmt.Errorf("Failed to run query: %v", err.Error())
	}

	return user, nil
}

func (u userCredentialRepository) ForgotPassword(email, newpass string) error {
	//TODO implement me
	query := "update user_credential set password = $2 where email = $1 "
	_, err := u.db.Exec(query, email, newpass)
	if err != nil {
		return fmt.Errorf("Failed to exec %v", err.Error())
	}
	return nil
}

func (u userCredentialRepository) GetUserPassword(email string) (string, error) {
	//TODO implement me
	query := "select password from user_credential where email = $1"

	var hashPassword string
	//do query row
	err := u.db.QueryRow(query, email).Scan(&hashPassword)
	if err != nil {
		return "", err
	}
	return hashPassword, nil
}

func (u userCredentialRepository) CheckEmailExist(email string) bool {
	//TODO implement me

	//do query untuk mencari apakah username sudah tersedia atau belum
	query := "select count(*) from user_credential where email=$1" // count(*) : menghitung jumlah baris

	var count int
	err := u.db.QueryRow(query, email).Scan(&count)
	if err != nil {
		fmt.Println(err.Error())
		return true //anggap username sudah ada jika error dalam query nya
	}
	return count > 0 //count > 0 mean is username already exist on database
}

func NewUserDetailsRepository(db *sql.DB) UserCredentialsRepository {
	return &userCredentialRepository{
		db: db,
	}
}
