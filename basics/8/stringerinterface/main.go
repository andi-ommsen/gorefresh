package main

import "fmt"

type Person interface {
	GetName() string
}

type BusinessPerson struct {
	id   int
	Name string
}

// Ist jetzt eine Methode, die den String-Interface implementiert
// dadruch wird bei fmt.Println() automatisch diese String()-Methode aufgerufen
// Heisst jeder Type, der die Methode String() implementiert,
// implementiert automatisch das Stringer-Interface

func (e BusinessPerson) String() string {
	return fmt.Sprintf("BusinessPerson: %s", e.Name)
}

func main() {

	jane := BusinessPerson{
		id:   1,
		Name: "Jane",
	}

	fmt.Println(jane)
}
