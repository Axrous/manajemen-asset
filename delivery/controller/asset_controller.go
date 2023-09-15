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
		return
	}

	err = a.usecase.Create(assetRequest)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"status": "Error", "message" : err.Error()})
		return
	}

	c.JSON(201, gin.H{"status": "OK", "message": "successfully created Asset", "asset" : assetRequest})
	// c.JSON(201, assetRequest)
}

func (a *AssetController) ListAssetHandler(c *gin.Context) {
	name := c.Query("name")
	if name != "" {
		assets, err := a.usecase.FindByName(name)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"status": "Error", "message" : err.Error()})
			return
		}
	
		c.JSON(200, gin.H{
			"status" : "OK",
			"assets" : assets,
		})
		return
	}


	assets, err := a.usecase.FindAll()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"status": "Error", "message" : err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"status" : "OK",
		"assets" : assets,
	})
}

func (a *AssetController) findByIdHandler(c *gin.Context)  {
	
	id := c.Param("id")

	asset, err := a.usecase.FindById(id)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"status": "Error", "message" : err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status" : "OK",
		"assets" : asset,
	})
}

func (a *AssetController) updateHandler(c *gin.Context)  {
	
	var assetRequest model.AssetRequest
	err := c.ShouldBindJSON(&assetRequest)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"status": "Error", "message" : err.Error()})
		return
	}

	err = a.usecase.Update(assetRequest)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"status": "Error", "message" : err.Error()})
		return
	}

	c.JSON(201, gin.H{"status": "OK", "message": "successfully Update Asset"})
}

func (a *AssetController) deleteHandler(c *gin.Context) {

	id := c.Param("id")

	err := a.usecase.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"status": "Error", "message" : err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "OK", "message": "successfully delete asset"})
}


func (a *AssetController) Route() {
	a.rg.POST("/assets", a.createAssetHandler)
	a.rg.GET("/assets", a.ListAssetHandler)
	a.rg.GET("/assets/:id", a.findByIdHandler)
	a.rg.PUT("/assets", a.updateHandler)
	a.rg.DELETE("/assets/:id", a.deleteHandler)
}


func NewAssetController(usecase usecase.AssetUsecase, rg *gin.RouterGroup) *AssetController {
	return &AssetController{
		usecase: usecase,
		rg:      rg,
	}
}