package handler

import (
	"fmt"
	"log"
	"restaurant/database"
	"restaurant/models"
	"restaurant/ui"

	"github.com/google/uuid"
)

func CreateOrder() {
	var newOrder models.Order

	newOrder.Id = uuid.New().String()

	ui.Tprint("Enter of Table number: ")
	var number int
	fmt.Scan(&number)

	table, _ := database.GetTableByNumber(database.ConnectToDb(), number)
	newOrder.TableId = table.Id

	ui.Tprint("Enter of Waiter Name: ")
	var waiterName string
	fmt.Scan(&waiterName)

	waiter, _ := database.GetWaiterByName(database.ConnectToDb(), waiterName)
	newOrder.WaiterId = waiter.Id

	err := database.CreateOrder(database.ConnectToDb(), newOrder)
	if err != nil {
		fmt.Println("Order yaratilmadi! :", err)
		return
	}

	fmt.Println("Order muvaffaqiyatli yaratildi!")

	_, err = database.GetProducts(database.ConnectToDb())

	flag := true
	if err == nil {

		for flag {

			var orderProducts models.OrderProducts

			products, err := database.GetProducts(database.ConnectToDb())
			if err != nil {
				return
			}
			ui.PrintProducts(products)

			ui.Tprint("Buyurtmalaringizni tanlang")
			var orderProduct string
			fmt.Scan(&orderProduct)
			product, flag := database.GetProductByName(database.ConnectToDb(), orderProduct)
			if flag {
				orderProducts.Id = uuid.New().String()
				orderProducts.OrederId = newOrder.Id
				orderProducts.ProductId = product.Id

				var quantity int
				ui.Tprint("Nechta buyurtma berishni hohlaysiz: ")
				fmt.Scan(&quantity)
				orderProducts.Quantity = quantity

				database.CalculateProductsPrice(&orderProducts)

				database.CreateOrderProducts(database.ConnectToDb(), orderProducts)
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
	err = database.UpdateOrder(database.ConnectToDb(), newOrder)
	if err != nil {
		log.Println("error on update order: ", err)
		return
	}
}

func UpdateOrder() {
	ui.Tprint("Enter of Table number: ")
	var number int
	fmt.Scan(&number)
	table, err := database.GetTableByNumber(database.ConnectToDb(), number)
	if err != nil {
		fmt.Println("Table not found!")
		return
	}

	order, err := database.GetTableOrder(database.ConnectToDb(), table)
	if err != nil {
		fmt.Println("No orders in this table!")
	}

	products, err := database.GetProducts(database.ConnectToDb())

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

			product, flag := database.GetProductByName(database.ConnectToDb(), orderProduct)
			if flag {
				orderProducts.Id = uuid.New().String()
				orderProducts.OrederId = order.Id
				orderProducts.ProductId = product.Id

				var quantity int
				ui.Tprint("Nechta buyurtma berishni hohlaysiz: ")
				fmt.Scan(&quantity)
				orderProducts.Quantity = quantity

				database.CalculateProductsPrice(&orderProducts)

				database.CreateOrderProducts(database.ConnectToDb(), orderProducts)
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
	err = database.UpdateOrder(database.ConnectToDb(), order)
	if err != nil {
		fmt.Println("Order update failed! :", err)
		return
	}
}
