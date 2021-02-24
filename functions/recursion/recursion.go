package main

import (
	"fmt"
)

func factorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("invalid number: %d", n)
	case n == 0:
		return 1, nil
	default:
		previousFactorial, _ := factorial(n - 1)
		return n * previousFactorial, nil
	}
}

//simple form, using uint
func factorial2(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * factorial2(n-1)
}

func main() {
	result, _ := factorial(5)
	fmt.Println(result)

	_, err := factorial(-1)
	if err != nil {
		fmt.Println(err)
	}

	result2 := factorial2(5)
	fmt.Println(result2)

}
