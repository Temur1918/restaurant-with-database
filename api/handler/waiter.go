package handler

import (
	"fmt"
	"restaurant/models"
	"restaurant/storage/postgres"
	"restaurant/ui"

	"github.com/google/uuid"
)

func CreateWaiter() {
	var newWaiter models.Waiter

	newWaiter.Id = uuid.New().String()

	ui.Tprint("Enter Waiter Name: ")
	fmt.Scan(&newWaiter.Name)

	err := postgres.CreateWaiter(newWaiter)
	if err != nil {
		fmt.Println("Waiter kiritilmadi! :", err)
		return
	}

	fmt.Println("Waiter bazaga qo'shildi")
}

func GetWaiters() {
	ui.PrintWaiter()
}

func DeleteWaiter() {
	ui.Tprint("Enter Waiter Name: --> ")
	var name string
	fmt.Scan(&name)
	waiterId, err := postgres.GetWaiterId(name)
	if err != nil {
		fmt.Println("Waiter not found")
		return
	}
	waiter := postgres.GetWaiter(waiterId)
	err = postgres.DeleteWaiter(waiter)
	if err != nil {
		fmt.Println("Waiter not found")
	} else {
		fmt.Println("Waiter deleted successfully")
	}
}
