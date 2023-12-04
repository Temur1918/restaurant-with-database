package handler

import (
	"fmt"
	"restaurant/database"
	"restaurant/models"
	"restaurant/ui"

	"github.com/google/uuid"
)

func CreateOrderProductsducts() {
	var newOrderProduct models.OrderProducts

	newOrderProduct.Id = uuid.New().String()

	ui.Tprint("Number of Order: ")
	fmt.Scan(&newOrderProduct.Quantity)

	err := database.CreateOrderProducts(database.ConnectToDb(), newOrderProduct)
	if err != nil {
		fmt.Println("OrderProduct kiritilmadi! :", err)
		return
	}

	fmt.Println("OrderProduct bazaga qo'shildi")
}
