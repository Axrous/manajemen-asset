package controller

import (
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/usecase"
	"final-project-enigma-clean/util/helper"
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"go.uber.org/zap"
)

type UserController struct {
	userUC usecase.UserCredentialUsecase
	gin    *gin.Engine
	logger *zap.Logger
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

	// Generate JWT
	token, err := helper.GenerateJWT(userID)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"Error": "Failed to generate jwt"})
		return
	}

	slog.Infof("New user with email : %v and jwt : %v", userLogin.Email, token)
	c.JSON(200, gin.H{"Message": "Successfully Login", "Token": token})
}

// init route
func (u *UserController) Route() {
	//grouping
	ug := u.gin.Group("/auth")
	//define middleware in here if u need it
	//ex : ug.Use(middleware.AuthMiddleware())
	{
		ug.POST("/register", u.RegisterUserHandler)
		ug.POST("/login", u.LoginUserHandler)
	}
}

func NewUserController(useruc usecase.UserCredentialUsecase, g *gin.Engine) *UserController {
	return &UserController{
		userUC: useruc,
		gin:    g,
	}
}
