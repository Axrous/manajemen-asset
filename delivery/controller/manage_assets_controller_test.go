package controller

import (
	"final-project-enigma-clean/__mock__/usecasemock"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type ManageAssetsControllerSuite struct {
	suite.Suite
	controller *ManageAssetController
	usecase    *usecasemock.ManageAssetsMock
	r          *gin.Engine
}

func (suite *ManageAssetsControllerSuite) SetupTest() {
	suite.usecase = new(usecasemock.ManageAssetsMock)
	suite.r = gin.New()
	rg := suite.r.Group("/api/v1")
	suite.controller = NewManageAssetController(suite.usecase, rg)
}
