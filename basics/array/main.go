package main

import "fmt"

func main() {
	var numbers [2]int
	fmt.Printf("%+v\n", numbers)

	numbers[0] = 1
	numbers[1] = 2

	fmt.Printf("%+v\n", numbers)

	var primes = [5]int{2, 3, 5, 7, 11}
	fmt.Printf("%+v\n", primes)
}
