package main

import "fmt"

func main() {
	i := 1
	var p *int = nil

	p = &i //getting the variable's address
	*p++   //getting the value (*p), and incrementing it (++)
	i++

	fmt.Println(p, *p, i, &i)
}
