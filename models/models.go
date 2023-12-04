package models

type Table struct {
	Id     string `json:"id"`
	Number uint8  `json:"number"`
}

type Waiter struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Product struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Order struct {
	Id       string          `json:"id"`
	TableId  string          `json:"table_id"`
	Products []OrderProducts `json:"products"`
	WaiterId string          `json:"waiter_id"`
	Price    float64         `json:"price"`
	Ispaid   bool            `json:"is_paid"`
}

type OrderProducts struct {
	Id        string  `json:"id"`
	OrederId  string  `json:"order_id"`
	ProductId string  `json:"product"`
	Quantity  uint8   `json:"quantity"`
	Price     float64 `json:"price"`
}

func (o *OrderProducts) CalculateProductsPrice() {
	// product, _ := database.GetOrderProductsProduct(database.ConnectToDb(), o)
	// if o != nil {
	// 	o.Price = float64(o.Quantity) * product.Price
	// } else {
	// 	o.Price = 0
	// }
}

func (o *Order) CalculateOrderPrice() {
	if o != nil {
		for _, v := range o.Products {
			o.Price += v.Price
		}
	}
}
