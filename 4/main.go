package main

import (
	"fmt"
	"strings"
)

var productPrice = map[string]float64{
	"TSHIRTS": 20.00,
	"MUG":     12.59,
	"HAT":     18.79,
	"BOOK":    9.99,
}

func calculateItemPrice(itemCode string) (float64, bool) {
	basePrice, found := productPrice[itemCode]
	if !found {
		if strings.HasSuffix(itemCode, "_SALE") {
			originalItemCode := strings.TrimSuffix(itemCode, "_SALE")
			basePrice, found = productPrice[originalItemCode]
			if found {
				salePrice := basePrice * 0.8 // Apply 20% discount for sale items
				fmt.Printf("Item %s is on sale! Original price: %.2f, Sale price: %.2f\n", originalItemCode, basePrice, salePrice)

				return salePrice, true
			}

			fmt.Printf("Item %s not found in product list.\n", originalItemCode)
			return 0.0, false
		}

		return 0.0, false
	}

	return basePrice, true
}

func main() {
	orderItems := []string{"TSHIRTS", "MUG", "HAT_SALE", "HAT",}
	var totalPrice float64 
	for _, item := range orderItems {
		price, found := calculateItemPrice(item)
		if found {
			totalPrice += price
		} else {
			fmt.Printf("Item %s not found in product list.\n", item)
		}
		
	}
    fmt.Printf("Total price for the order: %.2f\n", totalPrice)
}
