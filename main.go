package main

import (
	"fmt"
	"io/ioutil"
	"os"

	productPackage "github.com/musobarlab/go-protobuf-demo/product"
)

func main() {

	f, err := os.Open("aaa.jpg")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	imageData, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	product := &productPackage.Product{
		ID:       "1",
		Name:     "Samsung Galaxy S10",
		Quantity: 10,
		Image: productPackage.Image{
			URL:  "wuriyanto.com",
			Data: imageData,
		},
	}

	dataProto, err := product.ToProto()
	if err != nil {
		fmt.Println(err)
	}

	// -------------------------

	newProduct, _ := productPackage.FromProto(dataProto)
	fmt.Println(newProduct.Name)

	err = ioutil.WriteFile("out.jpg", newProduct.Image.Data, 0666)
	if err != nil {
		fmt.Println(err)
	}
}
