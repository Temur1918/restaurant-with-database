package postgres

import (
	"encoding/json"
	"io"
	"restaurant/config"
	"restaurant/models"
)

func CreateOrder(order models.Order) error {
	filepath := config.FilePathForOrder

	file, err := OpenFile(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	encode := json.NewEncoder(file)

	err = encode.Encode(order)
	if err != nil {
		return err
	}

	return nil
}

func GetOrders() ([]models.Order, error) {

	var orders []models.Order
	filepath := config.FilePathForOrder

	file, err := OpenFile(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	for {
		var order models.Order

		err := decoder.Decode(&order)

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func UpdateOrder(updateOrder models.Order) error {

	var orders []models.Order
	filepath := config.FilePathForOrder

	file, err := OpenFile(filepath)
	if err != nil {
		return nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	for {
		var order models.Order

		err := decoder.Decode(&order)

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil
		}

		if order.Id != updateOrder.Id {
			orders = append(orders, order)
		} else {
			orders = append(orders, updateOrder)
		}

	}

	err = config.ClearJSONFile(config.FilePathForOrder)
	if err != nil {
		return err
	}

	for _, order := range orders {
		err = CreateOrder(order)
	}
	if err != nil {
		return err
	}

	return nil
}
