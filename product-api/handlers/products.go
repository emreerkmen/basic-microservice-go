package handlers

import (
	"basic-microservice/hello/product-api/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	//catch others

	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
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
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
