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

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}
	//catch others

	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
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
