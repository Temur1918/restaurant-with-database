package database

import (
	"database/sql"
	"errors"
	"log"
	"restaurant/models"
)

func CreateWaiter(sqlDb *sql.DB, waiter models.Waiter) error {
	query := `
		INSERT INTO 
			waiter (id, name) 
			VALUES ($1, $2)
	`
	_, err := sqlDb.Exec(query, waiter.Id, waiter.Name)
	if err != nil {
		log.Println("error on add waiter to db: ", err)
		return err
	}

	return nil
}

func GetWaiters(sqlDb *sql.DB) ([]models.Waiter, error) {
	query := `
		SELECT * FROM 
			waiter
	`
	rows, err := sqlDb.Query(query)
	if err != nil {
		log.Println("error on get waitrers to db: ", err)
		return nil, err
	}

	defer rows.Close()

	var waiters []models.Waiter
	for rows.Next() {
		var waiter models.Waiter

		err := rows.Scan(&waiter.Id, &waiter.Name)
		if err != nil {
			log.Println("error on get waiters to db: ", err)
			return nil, err
		}

		waiters = append(waiters, waiter)
	}

	return waiters, nil
}

func GetWaiterById(sqlDb *sql.DB, id string) (models.Waiter, error) {
	query := `
		SELECT * FROM 
			waiter
		WHERE
			id = $1
	`
	var waiter models.Waiter
	row := sqlDb.QueryRow(query, id)
	if row == nil {
		err := errors.New("waiter not found")
		return models.Waiter{}, err
	}
	row.Scan(&waiter.Id, &waiter.Name)

	return waiter, nil
}

func GetWaiterByName(sqlDb *sql.DB, name string) (models.Waiter, error) {
	query := `
		SELECT * FROM 
			waiter
		WHERE
			name = $1
	`
	var waiter models.Waiter
	row := sqlDb.QueryRow(query, name)
	if row == nil {
		err := errors.New("waiter not found")
		return models.Waiter{}, err
	}
	row.Scan(&waiter.Id, &waiter.Name)

	return waiter, nil
}

func UpdateWaiter(sqlDb *sql.DB, waiter models.Waiter) error {
	query := `
		UPDATE 
			waiter 
		SET 
			name = $1 
		WHERE 
			id = $2
	`

	_, err := sqlDb.Exec(query, waiter.Name, waiter.Id)
	if err != nil {
		log.Println("error on updating data: ", err)
		return err
	}

	return nil
}

func DeleteWaiter(sqlDb *sql.DB, waiter models.Waiter) error {
	query := `
		DELETE FROM 
			waiter
		WHERE
			id = $1
	`
	_, err := sqlDb.Exec(query, waiter.Id)
	if err != nil {
		log.Println("error on deleted waiter unsuccessfully: ", err)
		return err
	}

	return nil
}
