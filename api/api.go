package api

import (
	"fmt"
	"restaurant/api/handler"
	"restaurant/config"
	"restaurant/ui"
)

func Router() {
	for {
		ui.PrintApi()
		var path string
		fmt.Print("choose a path: ")
		fmt.Scan(&path)
		config.Clear()
		fmt.Scanln()
		switch path {
		case "info":
			{
				ui.PrintRestaurantinfo()
			}
		case "create-table":
			{
				handler.CreateTable()
			}
		case "get-tables":
			{
				handler.GetTables()
			}
		case "get-table-check":
			{
				handler.GetTableCheck()
			}
		case "create-order":
			{
				handler.CreateOrder()
			}
		case "update-order":
			{
				handler.UpdateOrder()
			}
		// case "create-order-products":
		// 	{
		// 		handler.CreateOrderProductsducts()
		// 	}
		case "get-products":
			{
				handler.GetProducts()
			}
		case "get-product/id":
			{
				handler.GetProductId()
			}
		case "create-product":
			{
				handler.CreateProduct()
			}

		case "delete-product":
			{
				handler.DeleteProduct()
			}
		case "update-price-product":
			{
				handler.UpdatePriceProduct()
			}
		case "create-waiter":
			{
				handler.CreateWaiter()
			}
		case "delete-waiter":
			{
				handler.DeleteWaiter()
			}
		case "get-waiters":
			{
				handler.GetWaiters()
			}
		}
	}
}
