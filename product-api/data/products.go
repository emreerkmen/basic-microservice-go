package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// We create a Products type because now we can add metod for it
type Products []*Product

// We slide json produce responsibility to that struct from handler
// This json converter method is much faster
func (products *Products) ToJSON(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(products)
}

func (product *Product) FromJSON(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(product)
}

func AddProduct(product *Product) {
	product.ID = getIndex();
	productList = append(productList, product)
}

func getIndex() int {
	index := productList[len(productList)-1].ID
	return index+1
}

var productList = Products{
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

func GetProducts() Products {
	return productList
}
