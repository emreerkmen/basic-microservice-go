package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gte=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
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

func (product *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validetForSku)
	return validate.Struct(product)
}

// Structure
func validetForSku(fl validator.FieldLevel) bool {
	// sku example : asd-asd-afsdfa

	matched, err := regexp.Match(`[a-z]+-[a-z]+-[a-z]`, []byte(fl.Field().String()))
	fmt.Println(matched, err)
	return matched
}

func AddProduct(product *Product) {
	product.ID = getIndex()
	productList = append(productList, product)
}

func getIndex() int {
	index := productList[len(productList)-1].ID
	return index + 1
}

func UpdateProduct(id int, product *Product) error {
	_, position, err := findProduct(id)
	if err != nil {
		return err
	}

	product.ID = id
	productList[position] = product

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for index, product := range productList {
		if product.ID == id {
			return product, index, nil
		}
	}

	return nil, -1, ErrProductNotFound
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
