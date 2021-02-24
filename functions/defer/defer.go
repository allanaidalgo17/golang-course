package main

import "fmt"

func getValue(number int) int {
	//defer: makes the command be the last thing
	//the function will do before the return
	defer fmt.Println("the end!")

	if number > 5000 {
		fmt.Println("greater than 5000...")
		return 5000
	}
	fmt.Println("less than 5000...")
	return number
}

func main() {
	fmt.Println(getValue(6000))
}
