package main

import (
	"basic-microservice/hello/client"
	"basic-microservice/hello/client/products"
	"fmt"
	"testing"
)

func TestOurClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := products.NewListProductsParams()
	test, prod, err := c.Products.ListProducts(params)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v", test.GetPayload()[0])
	fmt.Printf("%#v", prod)
	t.Fail()
}