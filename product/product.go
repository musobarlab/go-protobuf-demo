package product

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"
	protogo "github.com/musobarlab/go-protobuf-demo/protogo/product"
)

// Product type
type Product struct {
	ID       string
	Name     string
	Quantity uint64
	Images   []string
}

// ToProto function
func (p *Product) ToProto() ([]byte, error) {
	productProto := &protogo.Product{
		ID:       p.ID,
		Name:     p.Name,
		Quantity: p.Quantity,
		Images:   p.Images,
	}
	return proto.Marshal(productProto)
}

// ToJSON function
func (p *Product) ToJSON() ([]byte, error) {
	return json.Marshal(p)
}

// FromProto function
func FromProto(data []byte) (*Product, error) {
	var p protogo.Product

	err := proto.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}

	return &Product{
		ID:       p.ID,
		Name:     p.Name,
		Quantity: p.Quantity,
		Images:   p.Images,
	}, nil
}

// FromJSON function
func FromJSON(data []byte) (*Product, error) {
	var p Product

	err := json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// Products type list of Product
type Products []*Product
