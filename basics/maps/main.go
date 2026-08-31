package main

import "fmt"

func main() {

	// init a map with string keys and int values
	stundentGrade := map[string]int{
		"John":   90,
		"Paul":   85,
		"George": 95,
		"Ringo":  80,
	}
	fmt.Printf("%+v\n", stundentGrade)

	// change the value of a key in the map
	stundentGrade["John"] = 75
	fmt.Printf("%+v\n", stundentGrade)

	// check if a key exists in the map
	john, ok := stundentGrade["John"]
	if ok {
		fmt.Printf("John's grade is %d\n", john)
	} else {
		fmt.Println("John's grade not found")
	}

	var key = "George"

	// check if a key exists in the map
	if grade, ok := stundentGrade[key]; ok {
		fmt.Printf("%s's grade is %d\n", key, grade)
	} else {
		fmt.Printf("%s's grade not found\n", key)
	}

}
