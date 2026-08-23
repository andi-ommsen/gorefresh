package main

import (
	"fmt"
)

func main() {

	// ---------------------------- IF/ELSE Statement ----------------------------

	tmp := 25

	// normales If/ELSE statement
	if tmp > 20 {
		fmt.Println("tmp is greater than 20")
	} else {
		fmt.Println("tmp is less than or equal to 20")
	}

	// Assoziative Arrays in PHP, in Go heißen sie "maps"
	access := map[string]bool{
		"admin": true,
		"user":  false,
	}

	// Initialisierung und Prüfung in einem Schritt
	if ok := access["admin"]; ok {
		fmt.Println("Access granted for admin")
	} else {
		fmt.Println("Access denied for admin")
	}

	// ---------------------------- SWITCH Statement ----------------------------

	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("It's Monday")
	case "Tuesday":
		fmt.Println("It's Tuesday")
	default:
		fmt.Println("It's another day")
	}

	typeCheck := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Println("Type is int:", v)
		case string:
			fmt.Println("Type is string:", v)
		case bool:
			fmt.Println("Type is bool:", v)
		default:
			fmt.Println("Unknown type")
		}

	}
	typeCheck(42)
	typeCheck("Hello, Go!")
	typeCheck(true)
	typeCheck(3.14) // Unknown type

}
