package main

import "fmt"

func modifyValue(val int) {
	val = val * 10
	fmt.Printf("modifyValue: %+v\n", val)
}

func modifyPointer(val *int) {
	if val == nil {
		fmt.Printf("modifyPointer: pointer is nil\n")
		return
	} else {
		*val = *val * 10
		fmt.Printf("modifyPointer: %+v\n", *val)
	}
}
func main() {
	// Pointer is a variable that stores the memory address of another variable.
	age := 30 //
	var agePointer *int = &age

	// Dereferencing a pointer " * " means accessing the value stored at the memory address the pointer is pointing to.
	// Here we are dereferencing the pointer and changing the value of the variable it points to.
	*agePointer = 31

	println(age)

	num := 10
	modifyValue(num)
	fmt.Printf("main: %+v\n", num)

	modifyPointer(&num)
	fmt.Printf("main: %+v\n", num)
}
