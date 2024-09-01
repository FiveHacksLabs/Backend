package services

import (
	"users/models"
	"users/utils"
)

func GetBarcodeTicket(stepItemID, statusItem string) (models.Ticket, error) {
	var ticket models.Ticket
	query := "SELECT id, step_item_id, status_item, image_barcode FROM tickets WHERE step_item_id = ? AND status_item = ?"
	row := utils.GetDB().QueryRow(query, stepItemID, statusItem)
	err := row.Scan(&ticket.ID, &ticket.StepItemID, &ticket.StatusItem, &ticket.ImageBarcode)
	return ticket, err
}
