package main

import (
	"fmt"
	"time"
)

//While does not exist in Go
func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Even ")
		} else {
			fmt.Print("Odd ")
		}
	}

	//endless loop
	for {
		fmt.Println("For... ever...")
		time.Sleep(time.Second)
	}
}
