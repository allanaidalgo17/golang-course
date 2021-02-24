package main

import (
	"fmt"
	"time"
)

func gradeToConcept(n float64) string {
	grade := int(n)
	switch grade {
	case 10:
		fallthrough //used to execute the next case
	case 9:
		return "A"
	case 7, 8:
		return "B"
	case 5, 6:
		return "C"
	case 3, 4:
		return "D"
	case 0, 1, 2:
		return "E"
	default:
		return "invalid"
	}
}

func daytime() {
	t := time.Now()

	switch { //switch true
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 18:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good night!")
	}
}

func getType(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case float32, float64:
		return "float"
	case func():
		return "function"
	default:
		return "invalid"
	}
}

func main() {
	fmt.Println(gradeToConcept(8.2))

	daytime()

	fmt.Println(getType(1))
	fmt.Println(getType(func() {}))
}
