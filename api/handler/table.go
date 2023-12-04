package handler

import (
	"fmt"
	"log"
	"restaurant/database"
	"restaurant/models"
	"restaurant/ui"

	"github.com/google/uuid"
)

func CreateTable() {
	var newTable models.Table

	newTable.Id = uuid.New().String()

	ui.Tprint("Enter Table Number: ")
	fmt.Scan(&newTable.Number)

	err := database.CreateTable(database.ConnectToDb(), newTable)
	if err != nil {
		fmt.Println("Table does not created! :", err)
		return
	}

	fmt.Println("Table Created")
}

func GetTables() {
	tables, err := database.GetTables(database.ConnectToDb())
	if err != nil {
		log.Println("tables does not find: ", err)
		return
	}
	ui.PrintTables(tables)
}

func GetTableCheck() {
	ui.Tprint("Enter the Table number  -->  ")
	var tableNumber int
	fmt.Scan(&tableNumber)

	table, order, err := database.GetTableCheck(database.ConnectToDb(), tableNumber)

	if err != nil {
		log.Println("table check does not print: ", err)
		return
	}

	ui.GetTableCheck(table, order)
}
