package postgres

import (
	"encoding/json"
	"errors"
	"io"
	"restaurant/config"
	"restaurant/models"
	"strings"
)

func CreateProduct(product models.Product) error {
	filepath := config.FilePathForProduct

	file, err := OpenFile(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	encode := json.NewEncoder(file)

	err = encode.Encode(product)
	if err != nil {
		return err
	}

	return nil
}

func GetProducts() ([]models.Product, error) {

	var products []models.Product
	filepath := config.FilePathForProduct

	file, err := OpenFile(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	for {
		var product models.Product

		err := decoder.Decode(&product)

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func GetProductId(productName string) (string, error) {
	products, err := GetProducts()
	for _, product := range products {
		if strings.EqualFold(strings.TrimSpace(productName), product.Name) {
			return product.Id, nil
		} else {
			err = errors.New("waiter not found")
		}
	}
	return "", err
}

func GetProduct(pk string) (models.Product, bool) {
	products, _ := GetProducts()
	for _, product := range products {
		if product.Id == pk {
			return product, true
		}
	}
	return models.Product{}, false
}

func GetProductName(name string) (models.Product, bool) {
	products, _ := GetProducts()
	for _, product := range products {
		// TrimSpace (bo'shliqlarni olib tashlaydi)
		// EqualFold (ikkala so'zni solisg=htiradi katta kichik harflar ahamiyattga ega emas)
		if strings.EqualFold(strings.TrimSpace(name), product.Name) {
			return product, true
		}
	}
	return models.Product{}, false
}

func DeleteProduct(deleteProduct models.Product) error {

	var products []models.Product
	filepath := config.FilePathForProduct

	file, err := OpenFile(filepath)
	if err != nil {
		return nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	for {
		var product models.Product

		err := decoder.Decode(&product)

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil
		}

		if product.Id != deleteProduct.Id {
			products = append(products, product)
		}

	}

	err = config.ClearJSONFile(config.FilePathForProduct)
	if err != nil {
		return err
	}

	for _, product := range products {
		err = CreateProduct(product)
	}
	if err != nil {
		return err
	}

	return nil
}

func UpdatePriceProduct(updateProduct models.Product) error {

	var products []models.Product
	filepath := config.FilePathForProduct

	file, err := OpenFile(filepath)
	if err != nil {
		return nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	for {
		var product models.Product

		err := decoder.Decode(&product)

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil
		}

		if product.Id != updateProduct.Id {
			products = append(products, product)
		} else {
			products = append(products, updateProduct)
		}

	}

	err = config.ClearJSONFile(config.FilePathForProduct)
	if err != nil {
		return err
	}

	for _, product := range products {
		err = CreateProduct(product)
	}
	if err != nil {
		return err
	}

	return nil
}
