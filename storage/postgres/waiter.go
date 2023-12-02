package postgres

import (
	"encoding/json"
	"errors"
	"io"
	"restaurant/config"
	"restaurant/models"
	"strings"
)

func CreateWaiter(waiter models.Waiter) error {
	filepath := config.FilePathForWaiter

	file, err := OpenFile(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	encode := json.NewEncoder(file)

	err = encode.Encode(waiter)
	if err != nil {
		return err
	}

	return nil
}

func GetWaiterId(waiterName string) (string, error) {
	waiters, err := GetWaiters()
	for _, waiter := range waiters {
		if strings.EqualFold(strings.TrimSpace(waiterName), waiter.Name) {
			return waiter.Id, nil
		} else {
			err = errors.New("waiter not found")
		}
	}
	return "", err
}

func GetWaiter(id string) (waiter models.Waiter) {
	waiters, _ := GetWaiters()
	for _, waiter := range waiters {
		if waiter.Id == id {
			return waiter
		}
	}
	return models.Waiter{}
}

func GetWaiterName(id string) string {
	waiters, _ := GetWaiters()
	for _, waiter := range waiters {
		if id == waiter.Id {
			return waiter.Name
		}
	}
	return ""
}

func GetWaiters() ([]models.Waiter, error) {

	var waiters []models.Waiter
	filepath := config.FilePathForWaiter

	file, err := OpenFile(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	for {
		var waiter models.Waiter

		err := decoder.Decode(&waiter)

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		waiters = append(waiters, waiter)
	}

	return waiters, nil
}

func DeleteWaiter(deleteWaiter models.Waiter) error {

	var waiters []models.Waiter
	filepath := config.FilePathForWaiter

	file, err := OpenFile(filepath)
	if err != nil {
		return nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	for {
		var waiter models.Waiter

		err := decoder.Decode(&waiter)

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil
		}

		if waiter.Id != deleteWaiter.Id {
			waiters = append(waiters, waiter)
		}

	}

	err = config.ClearJSONFile(config.FilePathForWaiter)
	if err != nil {
		return err
	}

	for _, waiter := range waiters {
		err = CreateWaiter(waiter)
	}
	if err != nil {
		return err
	}

	return nil
}
