package database

import (
	"database/sql"
	"errors"
	"log"
	"restaurant/models"
)

func CreateProduct(sqlDb *sql.DB, product models.Product) error {
	query := `
		INSERT INTO 
			product (id, name, price) 
			VALUES ($1, $2, $3)
	`
	_, err := sqlDb.Exec(query, product.Id, product.Name, product.Price)
	if err != nil {
		log.Println("error on add product to db: ", err)
		return err
	}

	return nil
}

func GetProducts(sqlDb *sql.DB) ([]models.Product, error) {
	query := `
		SELECT * FROM 
			product
	`
	rows, err := sqlDb.Query(query)
	if err != nil {
		log.Println("error on get products to db: ", err)
		return nil, err
	}

	defer rows.Close()

	products := []models.Product{}
	for rows.Next() {
		var product models.Product

		err := rows.Scan(&product.Id, &product.Name, &product.Price)
		if err != nil {
			log.Println("error on get products to db: ", err)
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func GetProductById(sqlDb *sql.DB, id string) (models.Product, error) {
	query := `
		SELECT * FROM 
			product
		WHERE
			id = $1
	`
	var product models.Product
	row := sqlDb.QueryRow(query, id)
	if row == nil {
		err := errors.New("product not found")
		return models.Product{}, err
	}
	row.Scan(&product.Id, &product.Name, &product.Price)

	return product, nil
}

func GetProductByName(sqlDb *sql.DB, name string) (models.Product, bool) {
	query := `
		SELECT * FROM 
			product
		WHERE
			name = $1
	`
	var product models.Product
	row := sqlDb.QueryRow(query, name)

	err := row.Scan(&product.Id, &product.Name, &product.Price)
	if err != nil {
		log.Println("error product not found: ", err)
		return models.Product{}, false
	}

	return product, true
}

func UpdateProduct(sqlDb *sql.DB, product models.Product) error {
	query := `
		UPDATE 
			product
		SET 
			price = $1
		WHERE 
			id = $2
	`

	_, err := sqlDb.Exec(query, product.Price, product.Id)
	if err != nil {
		log.Println("error on updating data: ", err)
		return err
	}

	return nil
}

func DeleteProduct(sqlDb *sql.DB, product models.Product) error {
	query := `
		DELETE FROM 
			product
		WHERE
			id = $1
	`
	_, err := sqlDb.Exec(query, product.Id)
	if err != nil {
		log.Println("error on deleted product unsuccessfully: ", err)
		return err
	}

	return nil
}
