package main

import "fmt"

func printResult(grade float64) {
	if grade >= 7 {
		fmt.Println("The grade is greater than 7;  grade =", grade)
	} else {
		fmt.Println("The grade is less than 7; grade =", grade)
	}
}

func gradeToConcept(grade float64) string {
	if grade >= 9 && grade <= 10 {
		return "A"
	} else if grade >= 8 && grade < 9 {
		return "B"
	} else if grade >= 5 && grade < 8 {
		return "C"
	} else {
		return "E"
	}
}

func main() {
	printResult(7.3)
	printResult(5.1)

	fmt.Println(gradeToConcept(9.2))
}
