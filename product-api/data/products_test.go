package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{Name: "kafee", Price: 2.4, SKU: "asd-asd-asd"}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
