package controller

import (
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/usecase"
	"final-project-enigma-clean/util/helper"

	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
<<<<<<< HEAD
	"go.uber.org/zap"
	"net/http"
=======
>>>>>>> 064dd7d6d89f1b6299f6efec1c16d78e5d153a22
)

type UserController struct {
	userUC usecase.UserCredentialUsecase
	rg     *gin.RouterGroup
}

// register handler
func (u *UserController) RegisterUserHandler(c *gin.Context) {
	var userRegist model.UserRegisterRequest

	//bind json
	if err := c.ShouldBindJSON(&userRegist); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": err.Error()})
		return
	}

	if err := u.userUC.RegisterUser(userRegist); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"Message": "Successfulyy Register"})

}

// login handler
func (u *UserController) LoginUserHandler(c *gin.Context) {
	var userLogin model.UserLoginRequest

	// Bind JSON
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": err.Error()})
		return
	}

	userID, err := u.userUC.LoginUser(userLogin)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"Error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
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

// init route
func (u *UserController) Route() {
	{
<<<<<<< HEAD
		ug.POST("/register", u.RegisterUserHandler)
		ug.POST("/login", u.LoginUserHandler)
		ug.POST("/login/email_otp/start", u.LoginOTPHandler)
=======
		u.rg.POST("/register", u.RegisterUserHandler)
		u.rg.POST("/login", u.LoginUserHandler)
>>>>>>> 064dd7d6d89f1b6299f6efec1c16d78e5d153a22
	}
}

func NewUserController(userUC usecase.UserCredentialUsecase, rg *gin.RouterGroup) *UserController {
	return &UserController{
		userUC: userUC,
		rg:     rg,
	}
}
