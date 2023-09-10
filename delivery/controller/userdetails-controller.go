package controller

import (
	"final-project-enigma-clean/delivery/controller/middleware"
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/usecase"
	"github.com/gin-gonic/gin"
)

type UserDetailsController struct {
	udetailsUC usecase.UserDetailsUsecase
	gin        *gin.Engine
}

func (u *UserDetailsController) SaveUserHandler(c *gin.Context) {
	var udetails model.UserDetails

	if err := c.ShouldBindJSON(&udetails); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"Data": udetails})
}

// define routing in here
func (u *UserDetailsController) Route() {
	//create a group
	ug := u.gin.Group("/app")
	ug.Use(middleware.AuthMiddleware()) // <---  init middleware dsini
	{
		ug.POST("/save-user", u.SaveUserHandler)
	}
}

func NewUserDetailsController(udetails usecase.UserDetailsUsecase, g *gin.Engine) *UserDetailsController {
	return &UserDetailsController{
		udetailsUC: udetails,
		gin:        g,
	}
}
