package models

type Ticket struct {
	ID               int    `json:"id"`
	StepItemID       int    `json:"stepItemId"`
	StatusItem       string `json:"statusItem"`
	ImageBarcode     string `json:"imageBarcode"`
}
