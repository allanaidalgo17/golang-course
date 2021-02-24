package main

import "fmt"

type product struct {
	name     string
	price    float64
	discount float64
}

// method: function with a receiver
func (p product) priceWithDiscount() float64 {
	return p.price * (1 - p.discount)
}

func main() {
	var product1 product
	product1 = product{
		name:     "p1",
		price:    100.0,
		discount: 0.05,
	}
	fmt.Println(product1, product1.priceWithDiscount())

	product2 := product{"p2", 1000.0, 0.10}
	fmt.Println(product2.priceWithDiscount())
}
