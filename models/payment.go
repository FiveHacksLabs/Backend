package models

type Payment struct {
	ID                int     `json:"id"`
	UserID            int     `json:"userId"`
	RouteID           int     `json:"routeId"`
	TotalPrice        float64 `json:"totalPrice"`
	ServiceFee        float64 `json:"serviceFee"`
	MethodPaymentCode string  `json:"methodPaymentCode"`
}
