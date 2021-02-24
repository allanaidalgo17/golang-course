package main

import "fmt"

func average(numbers ...float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}

	return total / float64(len(numbers))
}

func printApproved(approved ...string) {
	fmt.Println("Approved list:")
	for i, name := range approved {
		fmt.Printf("%d) %s\n", i, name)
	}
}

func main() {
	fmt.Printf("average = %.2f\n", average(7.7, 8.1, 5.9))
	fmt.Printf("average = %.2f\n", average(7.7, 8.1, 5.9, 9.9))

	approved := []string{"name1", "name2", "name3", "name4"}
	printApproved(approved...)
}
