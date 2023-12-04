package ui

import (
	"fmt"
	"restaurant/config"
	"restaurant/database"
	"restaurant/models"
)

func PrintApi() {
	fmt.Printf(`
	|---------------------------------------------------------------|
	|                          Restaurant
	|     -----------------------------------------------------     |
	|     restaurant:/info--
	|     restaurant:/create-table--
	|     restaurant:/get-tables--
	|     restaurant:/get-table-check
	|     restaurant:/create-order--
	|     restaurant:/update-order--
	|     restaurant:/get-products--
	|     restaurant:/get-product/id--
	|     restaurant:/create-product--
	|     restaurant:/delete-product--
	|     restaurant:/update-price-product--
	|     restaurant:/create-waiter--
	|     restaurant:/delete-waiter--
	|     restaurant:/get-waiters--
	|---------------------------------------------------------------|
	`)
	fmt.Println()
}

func PrintRestaurantinfo() {
	product, _ := database.GetProducts(database.ConnectToDb())
	products := len(product)
	waiter, _ := database.GetWaiters(database.ConnectToDb())
	waiters := len(waiter)
	table, _ := database.GetTables(database.ConnectToDb())
	tables := len(table)
	fmt.Printf(`
	|------------------Restaurant-------------------|
	| number of products: %d 
	| number of waiter: %d
	| number of table: %d
	|-----------------------------------------------|
	`, products, waiters, tables)

}

func Tprint(text string) {
	fmt.Printf(`
	~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	|   %s -->  `, text)
}

func PrintProduct(product models.Product) {
	fmt.Println("*********************************************************")
	fmt.Printf("\tName: %s\n\tPrice: %.2f\n", product.Name, product.Price)
	fmt.Println("*********************************************************")
}

func PrintProducts(products []models.Product) {
	fmt.Println("| ----------------------------  Menu  --------------------------- |")
	fmt.Println("| ---------------------------------------------------------------- |")
	index := 1
	for _, product := range products {
		fmt.Printf("| %d\t Name: %s\t\t\t Price: %.2f\n", index, product.Name, product.Price)
		index += 1
	}
	fmt.Println("| ---------------------------------------------------------------- |")
}

func PrintWaiter(waiters []models.Waiter) {
	fmt.Println("| --------------------------  Waiters  ------------------------- |")
	fmt.Println("*********************************************************")
	index := 1
	for _, waiter := range waiters {
		fmt.Printf("%d\t %s\n", index, waiter.Name)
		index += 1
	}
	fmt.Println("*********************************************************")
}

func PrintTables(tables []models.Table) {
	fmt.Println("| --------------------------  Tables  ------------------------- |")
	index := 1
	for _, table := range tables {
		fmt.Printf("\t%d\t", table.Number)
		if index%4 == 0 {
			fmt.Print("\n\n")
		}
		index += 1
	}
	fmt.Println("\n| ---------------------------------------------------------------- |")
}

func GetTableCheck(table models.Table, order models.Order) {
	waiter, _ := database.GetWaiterById(database.ConnectToDb(), order.WaiterId)
	fmt.Println("|----------------------------------------------|")
	fmt.Printf("					Table number: %d\n\n", table.Number)
	order_products, _ := database.GetOrderProducts(database.ConnectToDb(), order)
	fmt.Println(order_products)
	if len(order_products) > 0 {
		for _, order := range order_products {
			product, _ := database.GetOrderProductsProduct(database.ConnectToDb(), order)
			fmt.Printf("			--------%s--------		 	Jami\n", product.Name)
			fmt.Printf("			| %.2f  * %d       ", product.Price, order.Quantity)
			fmt.Printf("			 %.2f\n", order.Price)
		}
		fmt.Printf("			|------------------------------------------------\n")
		fmt.Printf("\n			 Waiter name:				 %s", waiter.Name)
		fmt.Printf("\n			|------------------------------------------------\n")
		fmt.Printf("			Jami					 %.2f", order.Price)
		fmt.Printf("\n			Servicce fee (19)			 %.2f", config.ServiceFee(order.Price))
		fmt.Printf("\n			Umumiy summa: 				 %.2f", order.Price+config.ServiceFee(order.Price))
	} else {
		fmt.Printf("			         Buyurtma yuq!\n")
	}

}
