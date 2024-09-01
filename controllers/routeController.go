package controllers

import (
	"users/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRoutes(c *gin.Context) {
	email := c.PostForm("email")
	price := c.PostForm("price")
	fromLocation := c.PostForm("fromLocation")
	toLocation := c.PostForm("toLocation")

	routes, err := services.GetRoutes(email, fromLocation, toLocation, price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": gin.H{"code": 500, "message": "Failed to get routes"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"code":    200,
			"message": "Success",
		},
		"data": gin.H{
			"listRoute": routes,
		},
	})
}

func GetRouteDetails(c *gin.Context) {
	routeID := c.PostForm("routeId")
	steps, err := services.GetRouteDetails(routeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": gin.H{"code": 500, "message": "Failed to get route details"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"code":    200,
			"message": "Success",
		},
		"data": gin.H{
			"stepRouteList": steps,
		},
	})
}
