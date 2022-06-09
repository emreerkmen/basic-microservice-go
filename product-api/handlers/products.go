package handlers

import (
	"basic-microservice/hello/product-api/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
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

	if r.Method == http.MethodPut {
		p.l.Println("PUT", r.URL.Path)
		// expect the id in the URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			p.l.Println("Invalid URI more than one id")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.l.Println("Invalid URI more than one capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Invalid URI unable to convert to numer", idString)
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.l.Println("Id",id)
		p.updateProducts(id, rw, r)
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

	err := data.ToJSON(listOfProducts,rw)

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
	err := data.FromJSON(product,r.Body)

	if err != nil {
		p.l.Printf("Product %v", err)
		http.Error(rw, "Unable to decode json", http.StatusBadRequest)
	}

	p.l.Printf("Product %#v", product)
	data.AddProduct(product)
}
func (p *Products) updateProducts(id int, rw http.ResponseWriter, r*http.Request) {
	p.l.Println("Handle PUT Product")

	prod := &data.Product{}

	err := data.FromJSON(prod,r.Body)
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
