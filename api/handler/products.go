package handler

import (
	"fmt"
	"log"
	"restaurant/database"
	"restaurant/models"
	"restaurant/ui"

	"github.com/google/uuid"
)

func CreateProduct() {
	var newProduct models.Product

	newProduct.Id = uuid.New().String()

	ui.Tprint("Enter product Name: ")
	fmt.Scan(&newProduct.Name)

	ui.Tprint("Enter product Price: ")
	fmt.Scan(&newProduct.Price)

	err := database.CreateProduct(database.ConnectToDb(), newProduct)
	if err != nil {
		fmt.Println("Product kiritilmadi! :", err)
		return
	}

	fmt.Println("Product bazaga qo'shildi")
}

func GetProducts() {
	products, err := database.GetProducts(database.ConnectToDb())
	if err != nil {
		log.Println("error does not products: ", err)
		return
	}
	ui.PrintProducts(products)
}

func GetProductId() {
	fmt.Print("Product Idsini kiriting: ")
	id := ""
	fmt.Scan(&id)

	product, err := database.GetProductById(database.ConnectToDb(), id)

	if err != nil {
		log.Println("product not found: ", err)
		return
	}

	ui.PrintProduct(product)
}

func GetProductName() {
	fmt.Print("Product Name ni kiriting: ")
	name := ""
	fmt.Scan(&name)

	product, flag := database.GetProductByName(database.ConnectToDb(), name)

	if !flag {
		log.Println("product not found!")
		return
	}

	ui.PrintProduct(product)
}

func DeleteProduct() {
	ui.Tprint("Enter Product Name: --> ")
	var name string
	fmt.Scan(&name)
	product, flag := database.GetProductByName(database.ConnectToDb(), name)
	if !flag {
		fmt.Println("Product not found")
		return
	}
	err := database.DeleteProduct(database.ConnectToDb(), product)
	if err != nil {
		fmt.Println("Product not found")
	} else {
		fmt.Println("Product deleted successfully")
	}
}

func UpdatePriceProduct() {
	ui.Tprint("Enter Product Name: --> ")
	var name string
	fmt.Scan(&name)
	ui.Tprint("Enter Product New Price: --> ")
	var newPrice float64
	fmt.Scan(&newPrice)
	product, flag := database.GetProductByName(database.ConnectToDb(), name)
	if !flag {
		fmt.Println("Product not found")
		return
	}
	product.Price = newPrice
	err := database.UpdateProduct(database.ConnectToDb(), product)
	if err != nil {
		fmt.Println("Product does not update")
	} else {
		fmt.Println("Product updated successfully")
	}
}
