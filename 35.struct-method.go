package main

import "fmt"

type Customer struct {
	Name string
	Address string
	Job string
}

func (customer Customer) sayHello() {
	fmt.Println(customer.Name)
}

func main() {
	var ryan Customer
	ryan.Name = "Ryan"
	ryan.Address = "Denpasar"
	ryan.Job = "Programmer"

	ryan.sayHello()
}