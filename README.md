## Benchmark Antara JSON VS Protocol Buffer

Pada pembahasan kali ini kita akan mencoba mengadu dua metode untuk `Serialization` yaitu `JSON` vs `Protocol Buffer`.
Simpel saja, kita akan menggunakan contoh data product dengan bentuk seperti ini

```go
type Product struct {
	ID       string
	Name     string
	Quantity uint64
	Images   []string
}
```

Untuk Protocol Buffer Serialization, kita butuh satu external dependency. Silahkan install terlebih dahulu.

```shell
$ glide get github.com/golang/protobuf/proto
```

Nanti data ini akan kita serialize ke JSON dan Protocol Buffer. Kita akan buat dua `method` di `Product` untuk masing-masing serializationnya dan deserializationnya.

Method `ToProto` untuk serialiaze dari data `Product` ke Protocol Buffer
```go
func (p *Product) ToProto() ([]byte, error) {
	productProto := &protogo.Product{
		ID:       p.ID,
		Name:     p.Name,
		Quantity: p.Quantity,
		Images:   p.Images,
	}
	return proto.Marshal(productProto)
}
```

Method `ToJSON` untuk serialiaze dari data `Product` ke JSON
```go
func (p *Product) ToJSON() ([]byte, error) {
	return json.Marshal(p)
}
```

Fungsi `FromProto` untuk serialiaze dari Protocol Buffer ke data `Product`
```go
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
```

Fungsi `FromJSON` untuk serialiaze dari JSON ke data `Product`
```go
// FromJSON function
func FromJSON(data []byte) (*Product, error) {
	var p protogo.Product

	err := json.Unmarshal(data, &p)
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
```

### Benchmark

Selanjutnya kita akan melakukan `Benchmarking` dari dua metode serialization tadi. Kita buat file baru dengan nama `product_test.go`.

`product_test.go`
```go
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
```

Kita jalankan secara bersamaan untuk serialize dan deserializenya.
```shell
$ make test_benc
```

Hasil
```shell
goos: darwin
goarch: amd64
pkg: github.com/musobarlab/go-protobuf-demo/product

BenchmarkSerializeProductToProto-8       	 5000000	       358 ns/op
BenchmarkSerializeProductToJSON-8        	 2000000	       718 ns/op
BenchmarkDeserializeProductFromProto-8   	 3000000	       569 ns/op
BenchmarkDeserializeProductFromJSON-8    	  500000	      2470 ns/op
```

Kita bisa lihat, untuk `serialization` Protocol Buffer menang telak dengan melakukan `5000000` kali iteration dengan waktu `358 ns` per operasi. Begitu juga dengan `deserialization`, Protocol Buffer menang telak dengan melakukan `3000000` kali iteration dengan waktu `569 ns` per operasi.