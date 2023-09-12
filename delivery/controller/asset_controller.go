package controller

import (
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/usecase"

	"github.com/gin-gonic/gin"
)

type AssetController struct {
	usecase usecase.AssetUsecase
	rg *gin.RouterGroup
}

func (a *AssetController) createAssetHandler(c *gin.Context) {
	
	var assetRequest model.AssetRequest
	err := c.ShouldBindJSON(&assetRequest)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"status": "Error", "message" : err.Error()})
	}

	err = a.usecase.Create(assetRequest)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"status": "Error", "message" : err.Error()})
	}

	c.JSON(201, gin.H{"status": "OK", "message": "successfully created Asset"})
}


func (a *AssetController) Route() {
	a.rg.POST("assets", a.createAssetHandler)
}


func NewAssetController(usecase usecase.AssetUsecase, rg *gin.RouterGroup) *AssetController {
	return &AssetController{
		usecase: usecase,
		rg:      rg,
	}
}