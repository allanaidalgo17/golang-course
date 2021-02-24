package main

import "fmt"

type car struct {
	name         string
	currentSpeed int
}

type ferrari struct {
	car     // anonymous field
	turboOn bool
}

func main() {
	f := ferrari{}
	f.name = "F40"
	f.currentSpeed = 0
	f.turboOn = true

	fmt.Println(f)
}
