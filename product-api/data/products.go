package data

import (
	"time"
)

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	SKU         string
	CreatedOn   string
	UpdatedOn   string
	DelatedOn   string
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Price:       2.45,
		Description: "A classic coffee drink made with espresso and steamed milk.",
		SKU:         "123ABC",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Price:       1.99,
		Description: "A strong shot of espresso made from the finest Arabica beans.",
		SKU:         "456DEF",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
