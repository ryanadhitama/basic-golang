package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func sayHello(hasName HasName) {
	fmt.Println("Hello,", hasName.GetName())
}

func main() {
	var ryan Person
	ryan.Name = "Ryan"

	sayHello(ryan)
}