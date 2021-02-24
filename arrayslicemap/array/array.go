package main

import "fmt"

func main() {
	//uniform(same data type) and static(fixed)
	var grades [3]float64
	fmt.Println(grades)

	grades[0], grades[1], grades[2] = 7.8, 4.3, 9.1
	fmt.Println(grades)

	total := 0.0
	for i := 0; i < len(grades); i++ {
		total += grades[i]
	}

	average := total / float64(len(grades))
	fmt.Printf("average = %.2f\n", average)
}
