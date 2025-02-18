package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"users/services"
)

// TransactionController menangani endpoint transaksi.
type TransactionController struct {
	Service services.TransactionService
}

// NewTransactionController mengembalikan instance TransactionController.
func NewTransactionController(s services.TransactionService) *TransactionController {
	return &TransactionController{Service: s}
}

// InquiryTransaction menangani GET /transactions/:id
func (ctrl *TransactionController) InquiryTransaction(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction id"})
		return
	}
	header, details, err := ctrl.Service.InquiryTransaction(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"header":  header,
		"details": details,
	})
}

// PaymentTransaction menangani POST /transactions/payment
func (ctrl *TransactionController) PaymentTransaction(c *gin.Context) {
	var req services.PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	header, err := ctrl.Service.PaymentTransaction(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, header)
}

// ReportTransaction menangani GET /transactions/report?start=YYYY-MM-DD&end=YYYY-MM-DD
func (ctrl *TransactionController) ReportTransaction(c *gin.Context) {
	startStr := c.Query("start")
	endStr := c.Query("end")
	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date"})
		return
	}
	end, err := time.Parse("2006-01-02", endStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date"})
		return
	}
	headers, err := ctrl.Service.ReportTransaction(start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, headers)
}
