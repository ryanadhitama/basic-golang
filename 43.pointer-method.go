package main

import "fmt"

type Person struct {
	Name string
}

func (person *Person) Married() {
	person.Name = "Mr." + person.Name
}

func main() {
	ryan := Person{"Ryan"}
	ryan.Married()

	fmt.Println(ryan.Name)
}