package handler

import (
	"fmt"
	"restaurant/models"
	"restaurant/storage/postgres"
	"restaurant/ui"

	"github.com/google/uuid"
)

func CreateOrderProductsducts() {
	var newOrderProduct models.OrderProducts

	newOrderProduct.Id = uuid.New().String()

	ui.Tprint("Number of Order: ")
	fmt.Scan(&newOrderProduct.Quantity)

	err := postgres.CreateOrderProducts(newOrderProduct)
	if err != nil {
		fmt.Println("OrderProduct kiritilmadi! :", err)
		return
	}

	fmt.Println("OrderProduct bazaga qo'shildi")
}
