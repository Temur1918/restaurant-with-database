package handler

import (
	"fmt"
	"restaurant/models"
	"restaurant/storage/postgres"
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

	err := postgres.CreateProduct(newProduct)
	if err != nil {
		fmt.Println("Product kiritilmadi! :", err)
		return
	}

	fmt.Println("Product bazaga qo'shildi")
}

func GetProducts() {
	ui.PrintProducts()
}

func GetProductId() {
	fmt.Print("Product Idsini kiriting: ")
	id := ""
	fmt.Scan(&id)
	product, flag := postgres.GetProduct(id)
	if flag {
		ui.PrintProduct(product)
	} else {
		fmt.Println("Ushbu Idga tegishli Product topilmadi!!!")
	}
}

func DeleteProduct() {
	ui.Tprint("Enter Product Name: --> ")
	var name string
	fmt.Scan(&name)
	productId, err := postgres.GetProductId(name)
	if err != nil {
		fmt.Println("Product not found")
		return
	}
	product, _ := postgres.GetProduct(productId)
	err = postgres.DeleteProduct(product)
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
	productId, err := postgres.GetProductId(name)
	if err != nil {
		fmt.Println("Product not found")
		return
	}
	product, _ := postgres.GetProduct(productId)
	product.Price = newPrice
	err = postgres.UpdatePriceProduct(product)
	if err != nil {
		fmt.Println("Product not found")
	} else {
		fmt.Println("Product updated successfully")
	}
}
