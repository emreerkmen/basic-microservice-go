package handlers

import (
	"basic-microservice/hello/product-api/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	listOfProducts := data.GetProducts()

	// That usage is much slower acording to NewEncodeer method that places in product structs.
	/*
		json, err := json.Marshal(listOfProducts)

		if err != nil {
			http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		}

		rw.Write(json)
	*/

	err := listOfProducts.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to encode json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")
	product := &data.Product{}

	// Golang use reader for request body. It actually buffered some stuff and read but not read all of the request at that time
	// It just read request body. Think of 10 gb file in request. In this scnerio you don't have to read all the respoanse expect body.
	// Actally it reads progressively. I wish there is a madium post about it.
	err := product.FromJSON(r.Body)

	if err != nil {
		p.l.Printf("Product %v", err)
		http.Error(rw, "Unable to decode json", http.StatusBadRequest)
	}

	p.l.Printf("Product %#v", product)
	data.AddProduct(product)
}
func (p *Products) UpdateProducts(rw http.ResponseWriter, r*http.Request) {

	//Gorilla extract variables from url automaticly.
	vars := mux.Vars(r);
	id, error := strconv.Atoi(vars["id"])
	if error != nil {
		http.Error(rw, "Unable to extract id", http.StatusBadRequest)
	}
	
	p.l.Println("Handle PUT Product",id)

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
