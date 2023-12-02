package postgres

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"restaurant/config"
	"restaurant/models"
)

func CreateTable(table models.Table) error {
	filepath := config.FilePathForTable

	file, err := OpenFile(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	encode := json.NewEncoder(file)

	err = encode.Encode(table)
	if err != nil {
		return err
	}

	return nil
}

func GetTables() ([]models.Table, error) {

	var tables []models.Table
	filepath := config.FilePathForTable

	file, err := OpenFile(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	for {
		var table models.Table

		err := decoder.Decode(&table)

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		tables = append(tables, table)
	}

	return tables, nil
}

func GetTableId(number uint8) (string, error) {
	tables, err := GetTables()
	for _, table := range tables {
		if table.Number == number {
			return table.Id, nil
		} else {
			err = errors.New("waiter not found")
		}
	}
	return "", err
}

func GetTable(id string) (models.Table, error) {
	tables, err := GetTables()
	for _, table := range tables {
		if table.Id == id {
			return table, nil
		} else {
			err = errors.New("waiter not found")
		}
	}
	return models.Table{}, err
}

func GetTableCheck(tableNumber int) (models.Table, models.Order, error) {
	tables, _ := GetTables()
	newTable := models.Table{}

	for _, table := range tables {
		if table.Number == uint8(tableNumber) {
			newTable = table
		}
	}

	order, err := GetTableOrder(newTable)
	if err == nil {
		return newTable, order, nil
	}

	return models.Table{}, models.Order{}, err

}

func GetTableOrder(table models.Table) (models.Order, error) {

	orders, err := GetOrders()

	for _, order := range orders {
		if order.TableId == table.Id {
			return order, nil
		}
	}

	fmt.Println("Buyurtma topilmadi!")

	return models.Order{}, err

}
