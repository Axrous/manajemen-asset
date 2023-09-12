package controller

import (
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/usecase"

	"github.com/gin-gonic/gin"
)

type TypeAssetController struct {
	typeAssetUC usecase.TypeAssetUseCase
	rg          *gin.RouterGroup
}

func (t *TypeAssetController) createHandlerTypeAsset(c *gin.Context) {
	var typeAsset model.TypeAsset
	if err := c.ShouldBindJSON(&typeAsset); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	err := t.typeAssetUC.CreateNew(typeAsset)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	response := gin.H{
		"message": "successfully created type asset",
	}
	c.JSON(200, response)
}
func (t *TypeAssetController) listHandlerTypeAsset(c *gin.Context) {

}
func (t *TypeAssetController) getByIdteHandlerTypeAsset(c *gin.Context) {

}
func (t *TypeAssetController) getByNameteHandlerTypeAsset(c *gin.Context) {

}
func (t *TypeAssetController) updateHandlerTypeAsset(c *gin.Context) {

}
func (t *TypeAssetController) deleteHandlerTypeAsset(c *gin.Context) {

}
func (t *TypeAssetController) Route() {
	t.rg.POST("/typeAsset", t.createHandlerTypeAsset)
	t.rg.GET("/typeAsset", t.listHandlerTypeAsset)
	t.rg.GET("/typeAsset/:id", t.getByIdteHandlerTypeAsset)
	t.rg.GET("/typeAsset/:name", t.getByNameteHandlerTypeAsset)
	t.rg.PUT("/typeAsset", t.updateHandlerTypeAsset)
	t.rg.DELETE("/typeAsset/:id", t.deleteHandlerTypeAsset)
}

func NewTypeAssetController(typeAssetUC usecase.TypeAssetUseCase, rg *gin.RouterGroup) *TypeAssetController {
	return &TypeAssetController{
		typeAssetUC: typeAssetUC,
		rg:          rg,
	}
}
