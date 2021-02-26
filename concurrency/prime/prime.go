package main

import (
	"fmt"
	"time"
)

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func primeNumbers(n int, c chan int) {
	initial := 2

	for i := 0; i < n; i++ {
		for prime := initial; ; prime++ {
			if isPrime(prime) {
				c <- prime
				initial = prime + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	// generates deadlock without the close() command
	close(c)
}

func main() {
	c := make(chan int, 30)
	go primeNumbers(60, c)

	for prime := range c {
		fmt.Printf("%d ", prime)
	}

	fmt.Println("The End!")
}
