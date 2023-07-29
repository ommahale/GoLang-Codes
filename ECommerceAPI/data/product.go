package data

import "time"

type Product struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Supplier   string  `json:"supplier"`
	Price      float32 `json:"price"`
	Created_at string  `json:"createdAt"`
	Updated_at string  `json:"-"`
}

var productList = []*Product{
	{
		ID:         1,
		Name:       "IPhone 14",
		Supplier:   "Apple inc.",
		Price:      80000.0,
		Created_at: time.Now().UTC().String(),
		Updated_at: time.Now().UTC().String(),
	},
	{
		ID:         2,
		Name:       "POCO F4",
		Supplier:   "Xiomi",
		Price:      50000.0,
		Created_at: time.Now().UTC().String(),
		Updated_at: time.Now().UTC().String(),
	},
}

func GetInventory() []*Product {
	return productList
}
