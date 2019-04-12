package product

import (
	"testing"
)

func BenchmarkSerializeProductToProto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		product := &Product{
			ID:       "1",
			Name:     "Samsung Galaxy S10",
			Quantity: 10,
			Images:   []string{"img1", "img2"},
		}

		product.ToProto()
	}
}

func BenchmarkSerializeProductToJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		product := &Product{
			ID:       "1",
			Name:     "Samsung Galaxy S10",
			Quantity: 10,
			Images:   []string{"img1", "img2"},
		}

		product.ToJSON()
	}
}

func BenchmarkDeserializeProductFromProto(b *testing.B) {
	product := &Product{
		ID:       "1",
		Name:     "Samsung Galaxy S10",
		Quantity: 10,
		Images:   []string{"img1", "img2"},
	}

	data, _ := product.ToProto()
	for i := 0; i < b.N; i++ {
		FromProto(data)
	}
}

func BenchmarkDeserializeProductFromJSON(b *testing.B) {
	product := &Product{
		ID:       "1",
		Name:     "Samsung Galaxy S10",
		Quantity: 10,
		Images:   []string{"img1", "img2"},
	}

	data, _ := product.ToJSON()
	for i := 0; i < b.N; i++ {
		FromJSON(data)
	}
}
