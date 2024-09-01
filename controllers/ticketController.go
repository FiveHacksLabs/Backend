package controllers

import (
	"users/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBarcodeTicket(c *gin.Context) {
	stepItemID := c.PostForm("stepItemId")
	statusItem := c.PostForm("statusItem")

	ticket, err := services.GetBarcodeTicket(stepItemID, statusItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": gin.H{"code": 500, "message": "Failed to get barcode ticket"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"code":    200,
			"message": "Success",
		},
		"data": ticket,
	})
}
