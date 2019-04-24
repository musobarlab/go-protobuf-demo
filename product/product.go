package product

import (
	"github.com/golang/protobuf/proto"
	protogo "github.com/musobarlab/go-protobuf-demo/protogo/product"
)

// Product type
type Product struct {
	ID       string
	Name     string
	Quantity uint64
	Image    Image
}

// Image type
type Image struct {
	URL  string
	Data []byte
}

// ToProto function
func (p *Product) ToProto() ([]byte, error) {

	productProto := &protogo.Product{
		ID:       p.ID,
		Name:     p.Name,
		Quantity: p.Quantity,
		Image: &protogo.Image{
			URL:  p.Image.URL,
			Data: p.Image.Data,
		},
	}
	return proto.Marshal(productProto)
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
		Image: Image{
			URL:  p.Image.URL,
			Data: p.Image.Data,
		},
	}, nil
}

// Products type list of Product
type Products []*Product
