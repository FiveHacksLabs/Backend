package controllers

import (
	"users/models"
	"users/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PayBill(c *gin.Context) {
	var payment models.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": gin.H{"code": 400, "message": "Invalid request"}})
		return
	}

	if err := services.PayBill(payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": gin.H{"code": 500, "message": "Failed to process payment"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"code":    200,
			"message": "Success Payment",
		},
		"data": gin.H{
			"routeId": payment.RouteID,
		},
	})
}
