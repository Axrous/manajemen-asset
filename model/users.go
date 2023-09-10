package model

import "time"

// model for user register request
type UserRegisterRequest struct {
	ID       string `json:"id"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" binding:"required"`
}

// model for user login request
type UserLoginRequest struct {
	ID       string `json:"id"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" binding:"required"`
}

// user details model
type UserDetails struct {
	ID          string    `json:"id" validate:"required"`
	UserID      string    `json:"user_id"`
	Name        string    `json:"name" validate:"required"`
	PhoneNumber string    `json:"phone_number" validate:"required"`
	Address     string    `json:"address" validate:"required"`
	BirthDate   time.Time `json:"birth_date" validate:"required"`
	ImgUrl      string    `json:"img_url" validate:"required"`
}
