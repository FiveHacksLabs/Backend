package main

import (
	"users/controllers"
	"users/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitDB()
	defer utils.GetDB().Close()

	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("/getRoute", controllers.GetRoutes)
	// r.POST("/getDetailRoute", controllers.GetRouteDetails)
	r.POST("/payBill", controllers.PayBill)
	r.POST("/getBarcodeTicket", controllers.GetBarcodeTicket)

	r.Run(":8080")
}
