package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"users/controllers"
	"users/services"
	"users/utils"
)

func main() {
	utils.InitDB()
	defer utils.GetDB().Close()

	transactionService := services.NewTransactionService(utils.GetDB())
	kycService := services.NewKYCService(utils.GetDB())

	router := gin.Default()
	api := router.Group("/api/v1")
	{
		// Endpoint transaksi
		api.GET("/transactions/:id", controllers.NewTransactionController(transactionService).InquiryTransaction)
		api.POST("/transactions/payment", controllers.NewTransactionController(transactionService).PaymentTransaction)
		api.GET("/transactions/report", controllers.NewTransactionController(transactionService).ReportTransaction)

		// Endpoint KYC
		api.GET("/kyc/:cif", controllers.NewKYCController(kycService).GetKYC)
		api.PUT("/kyc/:cif", controllers.NewKYCController(kycService).UpdateKYC)
		api.POST("/kyc", controllers.NewKYCController(kycService).CreateKYC)
	}

	if err := router.Run(":8081"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}

