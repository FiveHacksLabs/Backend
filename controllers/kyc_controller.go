package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"users/models"
	"users/services"
)

// KYCController menangani endpoint KYC.
type KYCController struct {
	Service services.KYCService
}

// NewKYCController mengembalikan instance KYCController.
func NewKYCController(s services.KYCService) *KYCController {
	return &KYCController{Service: s}
}

// GetKYC menangani GET /kyc/:cif
func (ctrl *KYCController) GetKYC(c *gin.Context) {
	cif := c.Param("cif")
	kyc, err := ctrl.Service.GetKYCByCIF(cif)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, kyc)
}

// UpdateKYC menangani PUT /kyc/:cif
func (ctrl *KYCController) UpdateKYC(c *gin.Context) {
	cif := c.Param("cif")
	var data models.KYCData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.Service.UpdateKYC(cif, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "KYC updated successfully"})
}

// CreateKYC menangani POST /kyc
func (ctrl *KYCController) CreateKYC(c *gin.Context) {
	var data models.KYCData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.Service.CreateKYC(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, data)
}
