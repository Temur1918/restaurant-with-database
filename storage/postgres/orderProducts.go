package postgres

import (
	"encoding/json"
	"restaurant/config"
	"restaurant/models"
)

func CreateOrderProducts(orderProducts models.OrderProducts) error {
	filepath := config.FilePathForOrderProducts

	file, err := OpenFile(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	encode := json.NewEncoder(file)

	err = encode.Encode(orderProducts)
	if err != nil {
		return err
	}

	return nil
}
