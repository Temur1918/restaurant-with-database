package models

type Table struct {
	Id     string `json:"id"`
	Number int    `json:"number"`
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
	Products  Product `json:"products"`
	ProductId string  `json:"product"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

func (o *Order) CalculateOrderPrice() {
	if o != nil {
		for _, v := range o.Products {
			o.Price += v.Price
		}
	}
}
