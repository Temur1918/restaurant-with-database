package ui

import (
	"fmt"
	"restaurant/config"
	"restaurant/models"
	"restaurant/storage/postgres"
)

func PrintApi() {
	fmt.Printf(`
	|---------------------------------------------------------------|
	|                          Restaurant
	|     -----------------------------------------------------     |
	|     restaurant:/info
	|     restaurant:/create-table
	|     restaurant:/get-tables
	|     restaurant:/get-table-check
	|     restaurant:/create-order
	|     restaurant:/update-order
	|     restaurant:/get-products
	|     restaurant:/get-product/id
	|     restaurant:/create-product
	|     restaurant:/delete-product
	|     restaurant:/update-price-product
	|     restaurant:/create-waiter
	|     restaurant:/delete-waiter
	|     restaurant:/get-waiters
	|---------------------------------------------------------------|
	`)
	fmt.Println()
}

func PrintRestaurantinfo() {
	product, _ := postgres.GetProducts()
	products := len(product)
	waiter, _ := postgres.GetWaiters()
	waiters := len(waiter)
	table, _ := postgres.GetTables()
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

func PrintProducts() {
	fmt.Println("| ----------------------------  Menu  --------------------------- |")
	fmt.Println("| ---------------------------------------------------------------- |")
	products, err := postgres.GetProducts()
	if err != nil {
		return
	}
	index := 1
	for _, product := range products {
		fmt.Printf("| %d\t Name: %s\t\t\t Price: %.2f\n", index, product.Name, product.Price)
		index += 1
	}
	fmt.Println("| ---------------------------------------------------------------- |")
}

func PrintWaiter() {
	fmt.Println("| --------------------------  Waiters  ------------------------- |")
	fmt.Println("*********************************************************")
	waiters, err := postgres.GetWaiters()
	if err != nil {
		return
	}
	index := 1
	for _, waiter := range waiters {
		fmt.Printf("%d\t %s\n", index, waiter.Name)
		index += 1
	}
	fmt.Println("*********************************************************")
}

func PrintTables() {
	fmt.Println("| --------------------------  Tables  ------------------------- |")
	tables, err := postgres.GetTables()
	if err != nil {
		return
	}
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
	fmt.Println("|----------------------------------------------|")
	fmt.Printf("					Table number: %d\n\n", table.Number)
	if len(order.Products) > 0 {
		for _, order := range order.Products {
			fmt.Printf("			--------%s--------		 	Jami\n", order.Product.Name)
			fmt.Printf("			| %.2f  * %d       ", order.Product.Price, order.Quantity)
			fmt.Printf("			 %.2f\n", order.Price)
		}
		fmt.Printf("			|------------------------------------------------\n")
		fmt.Printf("\n			 Waiter name:				 %s", postgres.GetWaiterName(order.WaiterId))
		fmt.Printf("\n			|------------------------------------------------\n")
		fmt.Printf("			Jami					 %.2f", order.Price)
		fmt.Printf("\n			Servicce fee (19)			 %.2f", config.ServiceFee(order.Price))
		fmt.Printf("\n			Umumiy summa: 				 %.2f", order.Price+config.ServiceFee(order.Price))
	} else {
		fmt.Printf("			         Buyurtma yuq!\n")
	}

}
