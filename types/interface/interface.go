package main

import "fmt"

type printable interface {
	toString() string
}

type person struct {
	name     string
	lastName string
}

type product struct {
	name  string
	price float64
}

// interfaces are implemented implicitly
func (p person) toString() string {
	return p.name + " " + p.lastName
}

func (p product) toString() string {
	return fmt.Sprintf("%s - %.2f", p.name, p.price)
}

func print(x printable) {
	fmt.Println(x.toString())
}

// another example
type sport interface {
	turnOnTurbo()
}

type ferrari struct {
	model        string
	currentSpeed int
	turboOn      bool
}

func (f *ferrari) turnOnTurbo() {
	f.turboOn = true
}

func main() {
	var v printable = person{"Name", "LastName"}
	fmt.Println(v.toString())
	print(v)

	v = product{"Product", 1000.0}
	fmt.Println(v.toString())
	print(v)

	p1 := product{"PS5", 4500.0}
	print(p1)

	// another example
	car1 := ferrari{"F40", 0, false}
	car1.turnOnTurbo()

	var car2 sport = &ferrari{"F40", 0, false}
	car2.turnOnTurbo()

	fmt.Println(car1, car2)
}
