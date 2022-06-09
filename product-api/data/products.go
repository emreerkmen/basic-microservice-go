package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// the id for the product
	//
	// required: false
	// min: 1
	ID int `json:"id"` // Unique identifier for the product

	// the name for this poduct
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for this poduct
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// the price for the product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"required,gt=0"`

	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"sku"`
}

// We create a Products type because now we can add metod for it
type Products []*Product

// We slide json produce responsibility to that struct from handler
// This json converter method is much faster
// ToJSON serializes the given interface into a string based JSON format
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(i)
}

// FromJSON deserializes the object from JSON string
// in an io.Reader to the given interface
func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
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

// DeleteProduct deletes a product from the database
func DeleteProduct(id int) error {
	i := findIndexByProductID(id)
	if i == -1 {
		return ErrProductNotFound
	}

	productList = append(productList[:i], productList[i+1])

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

// findIndex finds the index of a product in the database
// returns -1 when no product can be found
func findIndexByProductID(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}

	return -1
}


var productList = Products{
	&Product{ID: 1,
		Name:        "Filtre",
		Description: "Sade",
		Price:       10.50,
		SKU:         "abc123"},
	&Product{ID: 2,
		Name:        "Turk Kahvesi",
		Description: "Sade",
		Price:       5.50,
		SKU:         "qwe123"},
}

func GetProducts() Products {
	return productList
}
