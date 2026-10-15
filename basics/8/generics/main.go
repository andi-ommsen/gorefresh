package main

import "fmt"

// Generics sind ein neues Feature in Go 1.18, das es ermöglicht,
// Funktionen und Datentypen zu schreiben, die mit verschiedenen
// Typen arbeiten können, ohne dass der Code dupliziert werden muss.
// Generics ermöglichen es, Typparameter zu verwenden, um Funktionen
// und Datentypen zu definieren, die mit verschiedenen Typen arbeiten können.
func sum[T int | float64](a, b T) T {
	return a + b
}

func main() {

	// Beispielaufrufe der generischen Funktion sum
	// mit int
	fmt.Println(sum(5, 10))
	// mit float64
	fmt.Println(sum(5.5, 10.5))
}
