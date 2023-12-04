package database

import (
	"database/sql"
	"fmt"
	"log"
	"restaurant/models"
)

func CreateOrder(sqlDb *sql.DB, order models.Order) error {
	query := `
		INSERT INTO 
			order_ (id, price, is_paid, table_id, waiter_id)
		VALUES ($1, $2, $3, $4, $5)
	`
	fmt.Println("id: ", order.Id, "is_paid: ", order.Ispaid, "table_id: ", order.TableId, "waiter_id: ", order.WaiterId, "price: ", order.Price)
	_, err := sqlDb.Exec(query, order.Id, order.Price, order.Ispaid, order.TableId, order.WaiterId)

	if err != nil {
		log.Println("error on order add db: ", err)
		return err
	}

	return nil
}

func GetOrders(sqlDb *sql.DB) ([]models.Order, error) {
	query := `
		SELECT * FROM 
			order_
	`
	rows, err := sqlDb.Query(query)
	if err != nil {
		log.Println("error on get orders: ", err)
		return []models.Order{}, err
	}

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.Id, &order.Price, &order.Ispaid, &order.TableId, &order.WaiterId)
		if err != nil {
			log.Println("error on ge orders: ", err)
			return []models.Order{}, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func GetOrderById(sqlDb *sql.DB, id string) (models.Order, error) {
	query := `
		SELECT * FROM 
			order_
		WHERE
			id = $1
	`
	var order models.Order
	row := sqlDb.QueryRow(query, id)
	err := row.Scan(&order.Id, &order.Price, &order.Ispaid, &order.TableId, &order.WaiterId)

	if err != nil {
		log.Println("error order not found: ", err)
		return models.Order{}, err
	}

	return order, nil
}

func GetOrderByName(sqlDb *sql.DB, name string) (models.Order, error) {
	query := `
		SELECT * FROM 
			order_
		WHERE
			name = $1
	`
	var order models.Order
	row := sqlDb.QueryRow(query, name)
	err := row.Scan(&order.Id, &order.Price, &order.Ispaid, &order.TableId, &order.WaiterId)

	if err != nil {
		log.Println("error order not found: ", err)
		return models.Order{}, err
	}

	return order, nil
}

func UpdateOrder(sqlDb *sql.DB, order models.Order) error {
	query := `
		UPDATE 
			order_
		SET 
			price = $1, is_paid = $2 
		WHERE 
			id = $3
	`

	_, err := sqlDb.Exec(query, order.Price, order.Ispaid, order.Id)
	if err != nil {
		log.Println("error on updating data: ", err)
		return err
	}

	return nil
}

func DeleteOrder(sqlDb *sql.DB, order models.Order) error {
	query := `
		DELETE FROM 
			order_
		WHERE
			id = $1
	`
	_, err := sqlDb.Exec(query, order.Id)
	if err != nil {
		log.Println("error on deleted order unsuccessfully: ", err)
		return err
	}

	return nil
}
