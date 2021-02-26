package main

import (
	"fmt"
	"time"
)

// Channel - connects concurrent goroutines, sync point

func multiplyBy234(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)

	go multiplyBy234(2, c)
	fmt.Println("A")

	// <-c: sync point
	a, b := <-c, <-c // only continues when both values are ready

	fmt.Println("B")
	fmt.Println(a, b)

	fmt.Println(<-c)

	// deadlock
	fmt.Println(<-c)
}
