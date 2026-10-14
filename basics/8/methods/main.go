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

func FullName(e Employee) string {
	return e.FirtsName + " " + e.LastName
}

func deactivateEmployee(e *Employee) {
	e.IsActive = false
}

func NewEmployee(id int, firstName, lastName, position string, salary float64, isActive bool) Employee {
	return Employee{
		id:        id,
		FirtsName: firstName,
		LastName:  lastName,
		Position:  position,
		Salary:    salary,
		IsActive:  isActive,
		JoinedAt:  time.Now(),
	}
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
	fmt.Println("Full Name:", FullName(jane))
	deactivateEmployee(&jane)
	fmt.Println("Is Active after deactivation:", jane.IsActive)
}
