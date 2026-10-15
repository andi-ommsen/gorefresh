package main

type Person interface {
	GetName() string
}

type Employee struct {
	id   int
	Name string
}

type BusinessPerson struct {
	id   int
	Name string
}

// displayPersonName würde nicht mit einem Employee und BusinessPerson funktionieren,
// da beide Strukturen keine gemeinsame Schnittstelle implementieren.
// In Go können Funktionen, die auf bestimmte Typen zugreifen,
// nicht direkt auf andere Typen angewendet werden, selbst wenn sie ähnliche Felder haben.

// Um die Funktionalität zu erreichen, können wir eine Schnittstelle (Interface) definieren,
// die von beiden Strukturen implementiert wird. In diesem Fall haben wir die Schnittstelle Person definiert,
// die die Methode GetName() enthält. Sowohl Employee als auch BusinessPerson implementieren diese Methode,
// wodurch sie beide die Schnittstelle Person erfüllen und somit in der Funktion displayPersonName verwendet werden können.

func (e Employee) GetName() string {
	return e.Name
}

func (e BusinessPerson) GetName() string {
	return e.Name
}

func displayPersonName(p Person) string {
	return p.GetName()
}

func main() {

	joe := Employee{
		id:   1,
		Name: "Joe",
	}

	jane := BusinessPerson{
		id:   2,
		Name: "Jane",
	}

	println(displayPersonName(joe))
	println(displayPersonName(jane))
}
