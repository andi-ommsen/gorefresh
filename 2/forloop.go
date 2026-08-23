package main

import (
	"fmt"
)

func main() {

	// normal for loop
	for i := 0; i < 5; i++ {
		fmt.Println("for:", i)
	}

	//with while() style
	h := 5

	for h > 0 {
		fmt.Println("while:", h)
		h--
	}

	//range over slice
	nums := []int{1, 2, 3, 4, 5}
	for i, num := range nums {
		fmt.Printf("range: index=%d, value=%d\n", i, num)
	}

}
