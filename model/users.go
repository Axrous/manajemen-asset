package model

// user details model
type UserCredentials struct {
	ID       string `json:"id"`
	Email    string `json:"email,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
	Name     string `json:"name" validate:"required"`
	IsActive bool   `json:"is_active"`
}

type UserLoginRequest struct {
	ID       string `json:"id"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserRegisterRequest struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
	IsActive bool   `json:"is_active"`
}

type UserLoginOTPRequest struct {
	Email string `json:"email" validate:"required,email"`
	OTP   int    `json:"otp"`
}

type ChangePasswordRequest struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	OTP         int    `json:"otp"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ForgotPassRequest struct {
	Email              string `json:"email"`
	OTP                int    `json:"otp"`
	NewPassword        string `json:"new_password"`
	ConfirmNewPassword string `json:"confirm_password"`
}
