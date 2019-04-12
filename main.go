package main

import (
	"fmt"

	productPackage "github.com/musobarlab/go-protobuf-demo/product"
)

func main() {
	product := &productPackage.Product{
		ID:       "1",
		Name:     "Samsung Galaxy S10",
		Quantity: 10,
		Images:   []string{"img1", "img2"},
	}

	dataProto, err := product.ToProto()
	if err != nil {
		fmt.Println(err)
	}

	dataJSON, _ := product.ToJSON()
	fmt.Println(string(dataJSON))

	fmt.Println(dataProto)

	// -------------------------

	newProduct, _ := productPackage.FromProto(dataProto)
	fmt.Println(newProduct)
}
