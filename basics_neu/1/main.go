package main

import "fmt"

func main() {
	// We just split a dinner bill. Now the group wants a weekend away.
	// Let's set the trip up first.
	//
	// TODO: create a string variable `destination` (where are we going?)
	// TODO: create an int variable `travelers` (how many friends are coming?)
	// TODO: create a bool variable `bookingConfirmed`
	//
	// Then print a one-line summary with fmt.Printf, for example:
	//   Trip to Lisbon for 4 friends. Booked: true
	destination := "Kappeln"
	travelers := 3
	bookingConfirmed := true
	fmt.Println("Let's plan a trip")
	fmt.Printf("Trip to %s for %d friends. Booked: %v\n", destination, travelers, bookingConfirmed)
}
