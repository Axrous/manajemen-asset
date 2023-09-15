package controller

import (
	"final-project-enigma-clean/delivery/middleware"
	"final-project-enigma-clean/model"
	"final-project-enigma-clean/model/dto"
	"final-project-enigma-clean/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StaffController struct {
	staffUC usecase.StaffUseCase
	rg      *gin.RouterGroup
}

func (s *StaffController) createHandlerStaff(c *gin.Context) {
	var staff model.Staff
	if err := c.ShouldBindJSON(&staff); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	err := s.staffUC.CreateNew(staff)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	response := gin.H{
		"message": "successfully created staff",
	}
	c.JSON(200, response)
}
func (s *StaffController) listHandlerStaff(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "5"))
	staff, paging, err := s.staffUC.Paging(dto.PageRequest{
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
		"message": "successfully get staff",
		"data":    staff,
		"paging":  paging,
	}
	c.JSON(200, response)
}
func (s *StaffController) getByIdteHandlerStaff(c *gin.Context) {
	nik_staff := c.Param("nik_staff")
	staff, err := s.staffUC.FindById(nik_staff)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	response := gin.H{
		"message": "successfully get by nik staff staff",
		"data":    staff,
	}
	c.JSON(200, response)
}

func (s *StaffController) getByNameteHandlerStaff(c *gin.Context) {
	name := c.Param("name")
	staff, err := s.staffUC.FindByName(name)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	response := gin.H{
		"message": "successfully get by name type asset",
		"data":    staff,
	}
	c.JSON(200, response)
}
func (s *StaffController) updateHandlerStaff(c *gin.Context) {
	var staff model.Staff
	if err := c.ShouldBindJSON(&staff); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	err := s.staffUC.Update(staff)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "successfully update staff",
	})
}
func (s *StaffController) deleteHandlerStaff(c *gin.Context) {
	nik_staff := c.Param("nik_staff")
	if err := s.staffUC.Delete(nik_staff); err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	message := fmt.Sprintf("successfulyy delete staff with nik %s", nik_staff)
	c.JSON(200, gin.H{
		"message": message,
	})
}

func (s *StaffController) DownloadlistStaffHandler(c *gin.Context) {
	c.Set("Content-Disposition", `attachment; filename="data-staff.csv"`)
	c.Set("Content-Type", "text/csv")

	// Memanggil metode usecase untuk mengunduh data staf dalam format CSV
	csvData, err := s.staffUC.DownloadAllStaff()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"Error": "Failed to download staff data"})
		return
	}
	c.Data(http.StatusOK, "text/csv", csvData)
}

func (s *StaffController) Route() {
	s.rg.POST("/staffs", middleware.AuthMiddleware(), s.createHandlerStaff)
	s.rg.GET("/staffs", middleware.AuthMiddleware(), s.listHandlerStaff)
	s.rg.GET("/staffs/:nik_staff", middleware.AuthMiddleware(), s.getByIdteHandlerStaff)
	s.rg.GET("/staffs/name/:name", middleware.AuthMiddleware(), s.getByNameteHandlerStaff)
	s.rg.PUT("/staffs", middleware.AuthMiddleware(), s.updateHandlerStaff)
	s.rg.DELETE("/staffs/:nik_staff", middleware.AuthMiddleware(), s.deleteHandlerStaff)
	s.rg.GET("/staffs/list-staff-download", s.DownloadlistStaffHandler)
}

func NewStaffController(staffUC usecase.StaffUseCase, rg *gin.RouterGroup) *StaffController {
	return &StaffController{
		staffUC: staffUC,
		rg:      rg,
	}
}
