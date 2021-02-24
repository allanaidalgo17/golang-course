package main

import "fmt"

type sport interface {
	turnOnTurbo()
}

type luxury interface {
	autoPark()
}

type sportLuxury interface {
	sport
	luxury
	// can add more methods
}

type bmw struct{}

func (b bmw) turnOnTurbo() {
	fmt.Println("turbo on...")
}

func (b bmw) autoPark() {
	fmt.Println("parking...")
}

func main() {
	var b sportLuxury = bmw{}
	b.turnOnTurbo()
	b.autoPark()
}
