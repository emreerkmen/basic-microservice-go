package data

import "time"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	SKU         string
	CreatedOn   string
	UpdatedOn   string
	DeletedOn   string
}

var productList = []*Product{
	&Product{ID: 1,
		Name:        "Filtre",
		Description: "Sade",
		Price:       10.50,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String()},
	&Product{ID: 2,
		Name:        "Turk Kahvesi",
		Description: "Sade",
		Price:       5.50,
		SKU:         "qwe123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String()},
}

func GetProducts() []*Product {
	return productList
}
