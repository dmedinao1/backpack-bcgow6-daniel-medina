package main

import "fmt"

const (
	small     string  = "small"
	medium    string  = "medium2"
	big       string  = "big"
	mantThree float64 = 0.03
	mantSix   float64 = 0.06
)

type Producto interface {
	CostCalculate() float64
}

type producto struct {
	prodType string
	price    float64
	name     string
}

type Ecomerce interface {
	Total() float64
	Agregar(Producto)
}

type tienda struct {
	products []Producto
}

func (t tienda) Total() float64 {
	var total float64

	for _, product := range t.products {
		total += product.CostCalculate()
	}

	return total
}

func (t *tienda) Agregar(p Producto) {
	t.products = append(t.products, p)
}

func (p producto) CostCalculate() float64 {
	switch p.prodType {
	case small:
		return p.price
	case medium:
		return p.price + (p.price * mantThree)
	case big:
		return p.price + (p.price * mantSix)
	default:
		return 0
	}
}

func main() {
	prod := producto{
		prodType: big,
		name:     "chocorramo",
		price:    2_000,
	}

	tid := tienda{
		products: []Producto{prod},
	}
	tid.Agregar(producto{
		prodType: small,
		name:     "ponymalta",
		price:    3_100,
	})

	fmt.Printf("%+v\n", tid)
	fmt.Print("+===================")
	fmt.Printf("%+v\n", tid.Total())
}
