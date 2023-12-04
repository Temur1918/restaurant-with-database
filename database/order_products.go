package database

import (
	"database/sql"
	"fmt"
	"log"
	"restaurant/models"
)

func CalculateProductsPrice(o *models.OrderProducts) {
	product, _ := GetProductById(ConnectToDb(), o.ProductId)
	o.Price = float64(o.Quantity) * float64(product.Price)
}

func CreateOrderProducts(sqlDb *sql.DB, order_products models.OrderProducts) error {
	query := `
		INSERT INTO 
			order_products (id, quantity, price, order_id, product_id)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := sqlDb.Exec(query, order_products.Id, order_products.Quantity, order_products.Price, order_products.OrederId, order_products.ProductId)
	fmt.Println("id: ", order_products.Id, "quantity: ", order_products.Quantity, "price: ", order_products.Price, "order_id: ", order_products.OrederId, "product_id: ", order_products.ProductId)
	if err != nil {
		log.Println("error on order_products add db: ", err)
		return err
	}

	return nil
}

func GetOrderProductsProduct(sqlDb *sql.DB, orderProduct models.OrderProducts) (models.Product, error) {
	query := `
		SELECT * FROM
			product
		WHERE
			id = $1
	`
	var product models.Product
	rows := sqlDb.QueryRow(query, orderProduct.ProductId)
	err := rows.Scan(&product.Id, &product.Name, &product.Price)
	if err != nil {
		log.Println("products not found: ", err)
		return models.Product{}, err
	}

	return product, nil

}

func GetOrderProducts(sqlDb *sql.DB, order models.Order) ([]models.OrderProducts, error) {
	query := `
		SELECT * FROM 
			order_products 
		WHERE 
			order_id = $1
	`
	rows, err := sqlDb.Query(query, order.Id)
	if err != nil {
		log.Println("error on get order_products: ", err)
		return []models.OrderProducts{}, err
	}

	var order_products []models.OrderProducts
	for rows.Next() {
		var order_product models.OrderProducts
		err := rows.Scan(&order_product.Id, &order_product.Quantity, &order_product.Price, &order_product.OrederId, &order_product.ProductId)
		if err != nil {
			log.Println("error on get order_product: ", err)
			return []models.OrderProducts{}, err
		}

		order_products = append(order_products, order_product)
	}

	return order_products, nil
}

func GetOrderProductsById(sqlDb *sql.DB, id string) (models.OrderProducts, error) {
	query := `
		SELECT * FROM 
			order_products
		WHERE
			id = $1
	`
	var order_products models.OrderProducts
	row := sqlDb.QueryRow(query, id)
	err := row.Scan(&order_products.Id, &order_products.Quantity, &order_products.Price, &order_products.OrederId, &order_products.ProductId)

	if err != nil {
		log.Println("error order_products not found: ", err)
		return models.OrderProducts{}, err
	}

	return order_products, nil
}

func UpdateOrderProducts(sqlDb *sql.DB, order_products models.OrderProducts) error {
	query := `
		UPDATE 
			order_products
		SET 
			quantity = $1, price = $2
		WHERE 
			id = $3
	`

	_, err := sqlDb.Exec(query, order_products.Quantity, order_products.Price, order_products.Id)
	if err != nil {
		log.Println("error on updating data: ", err)
		return err
	}

	return nil
}

func DeleteOrderProducts(sqlDb *sql.DB, order_products models.OrderProducts) error {
	query := `
		DELETE FROM 
			order_products
		WHERE
			id = $1
	`
	_, err := sqlDb.Exec(query, order_products.Id)
	if err != nil {
		log.Println("error on deleted order_products unsuccessfully: ", err)
		return err
	}

	return nil
}
