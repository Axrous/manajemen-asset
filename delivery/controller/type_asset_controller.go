package controller

import (
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"
	"final-project-enigma-clean/usecase"
	"fmt"
	"strconv"

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
	// typeAsset.Id = helper.GenerateUUID()
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
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "5"))
	typeAsset, paging, err := t.typeAssetUC.Paging(dto.PageRequest{
		Page: page,
		Size: size,
	})
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	response := gin.H{
		"message": "successfully get type asset",
		"data":    typeAsset,
		"paging":  paging,
	}
	c.JSON(200, response)
}
func (t *TypeAssetController) getByIdteHandlerTypeAsset(c *gin.Context) {
	id := c.Param("id")
	typeAsset, err := t.typeAssetUC.FindById(id)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	response := gin.H{
		"message": "successfully get by id type asset",
		"data":    typeAsset,
	}
	c.JSON(200, response)
}

func (t *TypeAssetController) getByNameteHandlerTypeAsset(c *gin.Context) {
	name := c.Param("name")
	typeAsset, err := t.typeAssetUC.FindByName(name)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	response := gin.H{
		"message": "successfully get by name type asset",
		"data":    typeAsset,
	}
	c.JSON(200, response)
}
func (t *TypeAssetController) updateHandlerTypeAsset(c *gin.Context) {
	var typeAsset model.TypeAsset
	if err := c.ShouldBindJSON(&typeAsset); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	err := t.typeAssetUC.Update(typeAsset)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "successfully update type asset",
	})
}
func (t *TypeAssetController) deleteHandlerTypeAsset(c *gin.Context) {
	id := c.Param("id")
	if err := t.typeAssetUC.Delete(id); err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	message := fmt.Sprintf("successfulyy delete type asset with id %s", id)
	c.JSON(200, gin.H{
		"message": message,
	})
}
func (t *TypeAssetController) Route() {
	t.rg.POST("/typeAsset", t.createHandlerTypeAsset)
	t.rg.GET("/typeAsset", t.listHandlerTypeAsset)
	t.rg.GET("/typeAsset/:id", t.getByIdteHandlerTypeAsset)
	t.rg.GET("/typeAsset/name/:name", t.getByNameteHandlerTypeAsset)
	t.rg.PUT("/typeAsset", t.updateHandlerTypeAsset)
	t.rg.DELETE("/typeAsset/:id", t.deleteHandlerTypeAsset)
}

func NewTypeAssetController(typeAssetUC usecase.TypeAssetUseCase, rg *gin.RouterGroup) *TypeAssetController {
	return &TypeAssetController{
		typeAssetUC: typeAssetUC,
		rg:          rg,
	}
}
