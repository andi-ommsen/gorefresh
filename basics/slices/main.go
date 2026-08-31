package main

import "fmt"

func main() {
	// initialize a slice with 4 elements
	var names = []string{"John", "Paul", "George", "Ringo"}
	fmt.Printf("%+v\n", names)

	// append a new element to the slice
	names = append(names, "Pete")
	fmt.Printf("%+v\n", names)

	// create a slice with a length of 5 and a capacity of 10
	numbers := make([]int, 5, 10)
	fmt.Printf("%+v\n", numbers)

	numbers = append(numbers, 1, 2, 3, 4, 5)
	fmt.Printf("%+v\n", numbers)

	// slice a slcice. create  a new slice that contains elements 5 to 10 of the original slice
	sliced := numbers[5:10]
	fmt.Printf("%+v\n", sliced)
}
