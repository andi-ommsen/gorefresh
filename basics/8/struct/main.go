package main

import (
	"fmt"
	"time"
)

type Employee struct {
	id        int
	FirtsName string
	LastName  string
	Position  string
	Salary    float64
	IsActive  bool
	JoinedAt  time.Time
}

func main() {
	jane := Employee{
		id:        1,
		FirtsName: "Jane",
		LastName:  "Doe",
		Position:  "Software Engineer",
		Salary:    75000.00,
		IsActive:  true,
		JoinedAt:  time.Now(),
	}

	fmt.Println("Employee Details:")
	fmt.Println("ID:", jane.id)
	fmt.Println("First Name:", jane.FirtsName)
	fmt.Println("Last Name:", jane.LastName)
	fmt.Println("Position:", jane.Position)
	fmt.Println("Salary:", jane.Salary)
	fmt.Println("Is Active:", jane.IsActive)
	fmt.Println("Joined At:", jane.JoinedAt.Format(time.RFC1123))
}
