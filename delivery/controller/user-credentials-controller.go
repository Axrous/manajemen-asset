package controller

import (
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/usecase"
	"final-project-enigma-clean/util/helper"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
)

type UserController struct {
	userUC usecase.UserCredentialUsecase
	rg     *gin.RouterGroup
}

func (u *UserController) RegisterUserHandler(c *gin.Context) {
	var userRegist model.UserRegisterRequest

	//bind json
	if err := c.ShouldBindJSON(&userRegist); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": "Bad JSON Format"})
		return
	}

	if err := u.userUC.RegisterUser(userRegist); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"Message": "Successfully Register"})

}

func (u *UserController) LoginUserHandler(c *gin.Context) {
	var userLogin model.UserLoginRequest

	// Bind JSON
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": err.Error()})
		return
	}

	userID, err := u.userUC.LoginUser(userLogin)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": err.Error()})
		return
	}

	slog.Infof("New user trying to login with email : %v and user id : %v", userLogin.Email, userID)

	c.JSON(200, gin.H{"Message": "Successfully Login, check your email for verification token"})
}

// login handler with otp
func (u *UserController) LoginOTPHandler(c *gin.Context) {
	//var userLogin model.UserLoginRequest

	var request struct {
		Email string `json:"email"`
		OTP   int    `json:"otp"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Bad json format"})
		return
	}

	//store otp
	storedOTP, exists := usecase.OTPMap[request.Email]
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "OTP not found or expired"})
		return
	}

	//stored otp and then we need to generate jwt
	if request.OTP == storedOTP {
		token, err := helper.GenerateJWT(request.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to create token"})
			return
		}
		delete(usecase.OTPMap, request.Email)

		c.JSON(200, gin.H{"Message": "Login successfully", "Data": token})
	}
}

func (u *UserController) ChangePasswordHandler(c *gin.Context) {
	//bind json
	var userLogin model.ChangePasswordRequest

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": "Bad json format"})
		return
	}

	//find email + login otp
	_, err := u.userUC.LoginUserChangePass(userLogin)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"Message": fmt.Sprintf("We have sent you an email to %v with password change instructions", userLogin.Email)})
}

func (u *UserController) ChangePassOTPHandler(c *gin.Context) {

	var request struct {
		ID          string `json:"id"`
		Email       string `json:"email"`
		OTP         int    `json:"otp"`
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": "Bad JSON Format"})
		return
	}

	//is email exist?
	u.userUC.EmailExist(request.Email)

	//store otp
	storedOTP, exists := usecase.OTPMap[request.Email]
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "OTP not found or expired"})
		return
	}

	//stored otp
	if request.OTP == storedOTP {
		delete(usecase.OTPMap, request.Email)

		//get user password
		hashedPass, err := u.userUC.GetUserPassword(request.Email)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"Error to get password": err.Error()})
			return
		}

		//compare
		if err = bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(request.OldPassword)); err != nil {
			c.AbortWithStatusJSON(500, gin.H{"Error": "Invalid credentials"})
			return
		}

		//if compare successfully,then weneed to hash new password
		newHashPassword, err := helper.HashPasswordForgotPass(request.NewPassword)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"Error": "something is wrong"})
			return
		}

		if err = u.userUC.ChangePassword(request.Email, newHashPassword); err != nil {
			c.AbortWithStatusJSON(500, gin.H{"Error": "Invalid Password"})
			return
		}

		c.JSON(200, gin.H{"Message": "Successfully change password"})
	}
}

func (u *UserController) ForgotPassHandler(c *gin.Context) {
	var request model.ForgotPassRequest

	//bind json nya
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": "Bad json format"})
		return
	}

	_, err := u.userUC.FindingUserEmail(request.Email)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": "Email not found"})
		return
	}

	u.userUC.ForgotPass(request.Email)
	c.JSON(200, gin.H{"Success": "Check your email for verification and follow the instruction"})
}

func (u *UserController) ForgotPassOTPHandler(c *gin.Context) {
	var request model.ForgotPassRequest

	//bind json
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": "Bad JSON Format"})
		return
	}
	//is email exist?
	u.userUC.EmailExist(request.Email)

	//now validate otp
	storedOTP, exists := usecase.OTPMap[request.Email]
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "OTP not found or expired"})
		return
	}

	if request.OTP == storedOTP {
		//confirm new password
		if err := u.userUC.ForgotPassRequest(request.Email, request.NewPassword, request.ConfirmNewPassword); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"Error": err.Error()})
			return
		}

		//delete otp
		delete(usecase.OTPMap, request.Email)
	}
	c.JSON(200, gin.H{"Message": "Success"})
}

// init route
func (u *UserController) Route() {
	{
		u.rg.POST("/register", u.RegisterUserHandler)
		u.rg.POST("/login", u.LoginUserHandler)
		u.rg.POST("/login/email-otp/start", u.LoginOTPHandler)
		u.rg.POST("/change-password", u.ChangePasswordHandler)
		u.rg.POST("/change-password/start", u.ChangePassOTPHandler)
		u.rg.POST("/forgot-password", u.ForgotPassHandler)
		u.rg.POST("/forgot-password/start", u.ForgotPassOTPHandler)
	}
}

func NewUserController(userUC usecase.UserCredentialUsecase, rg *gin.RouterGroup) *UserController {
	return &UserController{
		userUC: userUC,
		rg:     rg,
	}
}
