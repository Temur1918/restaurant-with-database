package handler

import (
	"fmt"
	"restaurant/models"
	"restaurant/storage/postgres"
	"restaurant/ui"

	"github.com/google/uuid"
)

func CreateOrder() {
	var newOrder models.Order

	newOrder.Id = uuid.New().String()

	ui.Tprint("Enter of Table number: ")
	var number uint8
	fmt.Scan(&number)
	tableId, _ := postgres.GetTableId(number)
	newOrder.TableId = tableId

	ui.Tprint("Enter of Waiter Name: ")
	var waiterName string
	fmt.Scan(&waiterName)
	waiterId, _ := postgres.GetWaiterId(waiterName)
	newOrder.WaiterId = waiterId

	fmt.Println("Order muvaffaqiyatli yaratildi!")

	_, err := postgres.GetProducts()

	flag := true
	if err == nil {

		for flag {

			var orderProducts models.OrderProducts

			// for _, product := range products {
			// 	ui.PrintProduct(product)
			// }
			ui.PrintProducts()

			ui.Tprint("Buyurtmalaringizni tanlang")
			var orderProduct string
			fmt.Scan(&orderProduct)
			product, flag := postgres.GetProductName(orderProduct)
			if flag {
				orderProducts.Id = uuid.New().String()
				orderProducts.OrederId = newOrder.Id
				orderProducts.Product = product

				var quantity uint8
				ui.Tprint("Nechta buyurtma berishni hohlaysiz: ")
				fmt.Scan(&quantity)
				orderProducts.Quantity = quantity

				orderProducts.CalculateProductsPrice()

				postgres.CreateOrderProducts(orderProducts)
				newOrder.Products = append(newOrder.Products, orderProducts)
			} else {
				fmt.Println("Bu buyurtma mavjud emas!")
				continue
			}
			var optionOrder string
			ui.Tprint("Yana buyurtma berishni hohlaysizmi (Y/N)")
			fmt.Scan(&optionOrder)
			if optionOrder == "N" || optionOrder == "n" {
				flag = false
				fmt.Println("Buyurtmalar qabul qilindi!")
				break
			}
		}
	}

	newOrder.CalculateOrderPrice()
	err = postgres.CreateOrder(newOrder)
	if err != nil {
		fmt.Println("Order yaratilmadi! :", err)
		return
	}
}

func UpdateOrder() {
	ui.Tprint("Enter of Table number: ")
	var number uint8
	fmt.Scan(&number)
	tableId, _ := postgres.GetTableId(number)
	table, err := postgres.GetTable(tableId)
	if err != nil {
		fmt.Println("Table not found!")
		return
	}
	order, err := postgres.GetTableOrder(table)
	if err != nil {
		fmt.Println("No orders in this table!")
	}
	order.Price = 0

	products, err := postgres.GetProducts()

	flag := true
	if err == nil {

		for flag {

			var orderProducts models.OrderProducts

			for _, product := range products {
				ui.PrintProduct(product)
			}

			ui.Tprint("Buyurtmalaringizni tanlang")
			var orderProduct string
			fmt.Scan(&orderProduct)
			product, flag := postgres.GetProductName(orderProduct)
			if flag {
				orderProducts.Id = uuid.New().String()
				orderProducts.OrederId = order.Id
				orderProducts.Product = product

				var quantity uint8
				ui.Tprint("Nechta buyurtma berishni hohlaysiz: ")
				fmt.Scan(&quantity)
				orderProducts.Quantity = quantity

				orderProducts.CalculateProductsPrice()

				postgres.CreateOrderProducts(orderProducts)
				order.Products = append(order.Products, orderProducts)
			} else {
				fmt.Println("Bu buyurtma mavjud emas!")
				continue
			}
			var optionOrder string
			ui.Tprint("Yana buyurtma berishni hohlaysizmi (Y/N)")
			fmt.Scan(&optionOrder)
			if optionOrder == "N" || optionOrder == "n" {
				flag = false
				fmt.Println("Buyurtmalar qabul qilindi!")
				break
			}
		}
	}

	order.CalculateOrderPrice()
	err = postgres.UpdateOrder(order)
	if err != nil {
		fmt.Println("Order update failed! :", err)
		return
	}
}
