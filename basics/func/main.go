package main

import "fmt"

func greeting(name string) {
	fmt.Println("Hello, " + name + "!")
}

func main() {
	println("Hello, World!")
	greeting("jens")
}
