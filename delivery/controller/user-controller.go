package controller

import (
	"final-project-enigma-clean/delivery/middleware"
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUC usecase.UserUsecase
	gin    *gin.Engine
}

func (u *UserController) RegisterUserHandler(c *gin.Context) {
	var userloginReq model.UserRegisterRequest

	if err := u.userUC.UserRegist(userloginReq); err != nil {
		c.AbortWithStatusJSON(500, gin.H{"Error": err.Error()})
		return
	}

	//bind json
	if err := c.ShouldBindJSON(&userloginReq); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"Message": "Successfully register"})
}

func (u *UserController) LoginUserHandler() {

}

func (u *UserController) Route() {
	ug := u.gin.Group("/auth")
	ug.Use(middleware.AuthMiddleware())
	{
		ug.POST("/register", u.RegisterUserHandler)
	}
}

func NewUserController(useruc usecase.UserUsecase, g *gin.Engine) *UserController {
	return &UserController{
		userUC: useruc,
		gin:    g,
	}
}
