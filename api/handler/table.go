package handler

import (
	"fmt"
	"restaurant/models"
	"restaurant/storage/postgres"
	"restaurant/ui"

	"github.com/google/uuid"
)

func CreateTable() {
	var newTable models.Table

	newTable.Id = uuid.New().String()

	// fmt.Print("Enter Table Number: ")
	ui.Tprint("Enter Table Number: ")
	fmt.Scan(&newTable.Number)

	err := postgres.CreateTable(newTable)
	if err != nil {
		fmt.Println("Table does not created! :", err)
		return
	}

	fmt.Println("Table Created")
}

func GetTables() {
	ui.PrintTables()
}

func GetTableCheck() {
	ui.Tprint("Enter the Table number  -->  ")
	var tableNumber int
	fmt.Scan(&tableNumber)

	table, order, err := postgres.GetTableCheck(tableNumber)

	if err == nil {

		ui.GetTableCheck(table, order)

	}
}
