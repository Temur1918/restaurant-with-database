package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"restaurant/models"
)

func CreateTable(sqlDb *sql.DB, table models.Table) error {
	query := `
		INSERT INTO 
			table_ (id, number) 
			VALUES ($1, $2)
	`
	_, err := sqlDb.Exec(query, table.Id, table.Number)
	if err != nil {
		log.Println("error on add table to db: ", err)
		return err
	}

	return nil
}

func GetTables(sqlDb *sql.DB) ([]models.Table, error) {
	query := `
		SELECT * FROM 
			table_
	`
	rows, err := sqlDb.Query(query)
	if err != nil {
		log.Println("error on get tables to db: ", err)
		return nil, err
	}

	defer rows.Close()

	var tables []models.Table
	for rows.Next() {
		var table models.Table

		err := rows.Scan(&table.Id, &table.Number)
		if err != nil {
			log.Println("error on get tables to db: ", err)
			return nil, err
		}

		tables = append(tables, table)
	}

	return tables, nil
}

func GetTableById(sqlDb *sql.DB, id string) (models.Table, error) {
	query := `
		SELECT * FROM 
			table_
		WHERE
			id = $1
	`
	var table models.Table
	row := sqlDb.QueryRow(query, id)
	if row == nil {
		err := errors.New("table not found")
		return models.Table{}, err
	}
	row.Scan(&table.Id, &table.Number)

	return table, nil
}

func GetTableByNumber(sqlDb *sql.DB, number int) (models.Table, error) {
	query := `
		SELECT * FROM 
			table_
		WHERE
			number = $1
	`
	var table models.Table
	row := sqlDb.QueryRow(query, number)
	if row == nil {
		err := errors.New("table not found")
		return models.Table{}, err
	}
	row.Scan(&table.Id, &table.Number)

	return table, nil
}

func UpdateTable(sqlDb *sql.DB, table models.Table) error {
	query := `
		UPDATE 
			table_
		SET 
			number = $1 
		WHERE 
			id = $2
	`

	_, err := sqlDb.Exec(query, table.Number, table.Id)
	if err != nil {
		log.Println("error on updating data: ", err)
		return err
	}

	return nil
}

func DeleteTable(sqlDb *sql.DB, table models.Table) error {
	query := `
		DELETE FROM 
			table_
		WHERE
			id = $1
	`
	_, err := sqlDb.Exec(query, table.Id)
	if err != nil {
		log.Println("error on deleted table unsuccessfully: ", err)
		return err
	}

	return nil
}

func GetTableOrder(sqlDb *sql.DB, table models.Table) (models.Order, error) {
	query := `
		SELECT * FROM 
			order_
		WHERE order_.table_id = $1
	`
	order := models.Order{}
	row := sqlDb.QueryRow(query, table.Id)
	err := row.Scan(&order.Id, &order.Price, &order.Ispaid, &order.TableId, &order.WaiterId)

	if err != nil {
		log.Println("order does not find: ", err)
		return models.Order{}, err
	}

	fmt.Println("order, order: ", order)

	return order, nil
}

func GetTableCheck(sqlDb *sql.DB, tableNumber int) (models.Table, models.Order, error) {
	table, err := GetTableByNumber(sqlDb, tableNumber)
	if err != nil {
		log.Println("table does not find in db: ", err)
		return models.Table{}, models.Order{}, err
	}

	order, err := GetTableOrder(sqlDb, table)
	if err != nil {
		log.Println("table order does not find in db: ", err)
		return models.Table{}, models.Order{}, err
	}

	return table, order, nil

}
