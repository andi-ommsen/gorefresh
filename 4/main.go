package main

import (
	"fmt"
)

var productPrice = map[string]float64{
	"TSHIRTS": 20.00,
	"MUG":     12.59,
	"HAT":     18.79,
	"BOOK":    9.99,
}

func calculateItemPrice(item string, quantity int) (float64, bool) {
	price, exists := productPrice[item]
	if !exists {
		return 0.0, false
	}
	return price * float64(quantity), true
}

func main() {
	fmt.Println(calculateItemPrice("TSHIRTS", 2))
}
