package services

import (
	"users/utils"
	"users/models"
)

func PayBill(payment models.Payment) error {
	query := "INSERT INTO payments (user_id, route_id, total_price, service_fee, method_payment_code) VALUES (?, ?, ?, ?, ?)"
	_, err := utils.GetDB().Exec(query, payment.UserID, payment.RouteID, payment.TotalPrice, payment.ServiceFee, payment.MethodPaymentCode)
	return err
}
