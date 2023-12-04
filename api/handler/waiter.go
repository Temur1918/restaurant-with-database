package handler

import (
	"fmt"
	"log"
	"restaurant/database"
	"restaurant/models"
	"restaurant/ui"

	"github.com/google/uuid"
)

func CreateWaiter() {
	var newWaiter models.Waiter

	newWaiter.Id = uuid.New().String()

	ui.Tprint("Enter Waiter Name: ")
	fmt.Scan(&newWaiter.Name)

	err := database.CreateWaiter(database.ConnectToDb(), newWaiter)
	if err != nil {
		fmt.Println("Waiter kiritilmadi! :", err)
		return
	}

	fmt.Println("Waiter bazaga qo'shildi")
}

func GetWaiters() {
	waiters, err := database.GetWaiters(database.ConnectToDb())
	if err != nil {
		log.Println("waiters not found: ", err)
		return
	}
	ui.PrintWaiter(waiters)
}

func DeleteWaiter() {
	ui.Tprint("Enter Waiter Name: --> ")
	var name string
	fmt.Scan(&name)
	waiter, err := database.GetWaiterByName(database.ConnectToDb(), name)
	if err != nil {
		fmt.Println("Waiter not found")
		return
	}

	err = database.DeleteWaiter(database.ConnectToDb(), waiter)
	if err != nil {
		fmt.Println("Waiter not found")
		return
	}

	fmt.Println("Waiter deleted successfully")
}
