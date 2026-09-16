// Handlers for the pizza order tracking application


package main 

import (
	"pizza-tracker-go/internal/models"
)


type Handler struct {
	order *models.OrderModel
}


// when we creat the handels 
func NewHandler(dbModel *models.DBModel) *Handler {
	return &Handler{
		order: &dbModel.Order,
	}
}