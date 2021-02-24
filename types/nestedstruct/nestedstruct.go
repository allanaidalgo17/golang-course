package main

import "fmt"

type item struct {
	productID int
	quantity  int
	price     float64
}

type order struct {
	userID int
	items  []item
}

func (o order) totalValue() float64 {
	total := 0.0
	for _, item := range o.items {
		total += item.price * float64(item.quantity)
	}

	return total
}

func main() {
	order := order{
		userID: 1,
		items: []item{
			{productID: 1, quantity: 2, price: 100.0},
			{productID: 2, quantity: 1, price: 200.0},
			{productID: 3, quantity: 12, price: 50.0},
		},
	}

	fmt.Printf("The total values is: %.2f\n", order.totalValue())
}
