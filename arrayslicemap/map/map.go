package main

import "fmt"

func main() {
	//maps must be initialized
	m := make(map[int]string)

	m[123] = "123"
	m[456] = "456"
	m[789] = "789"

	fmt.Println(m)

	for key, value := range m {
		fmt.Printf("%d) %s\n", key, value)
	}

	delete(m, 456)
	fmt.Println(m)

	//other examples
	prices := map[string]float64{
		"product1": 1234.0,
		"p2":       345.6,
		"p3":       1000.0,
	}

	prices["new product"] = 500.0
	delete(prices, "non existent")

	productsPerLetter := map[string]map[string]float64{
		"G": {
			"game1": 4500.0,
			"game2": 4000.0,
		},
		"J": {
			"j1": 1000.2,
		},
		"P": {
			"p1": 123.4,
			"p2": 567.8,
		},
	}

	delete(productsPerLetter, "P")

	for letter, products := range productsPerLetter {
		fmt.Println(letter, products)
	}

}
