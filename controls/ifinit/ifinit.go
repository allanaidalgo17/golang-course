package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rng() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := rng(); i > 5 {
		fmt.Println("Greater than 5; i =", i)
	} else {
		fmt.Println("Less than or equal to 5; i =", i)
	}
}
