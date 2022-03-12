package handlers

import (
	"basic-microservice/hello/product-api/data"
	"encoding/json"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products{
	return &Products{l}
}

func (g *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	listOfProducts := data.GetProducts()
	json, err := json.Marshal(listOfProducts)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	rw.Write(json)
}