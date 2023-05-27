package main

import "fmt"

func main() {
	type Customer struct {
		Name string
		Address string
		Job string
	}

	var ryan Customer
	ryan.Name = "Ryan"
	ryan.Address = "Denpasar"
	ryan.Job = "Programmer"
	fmt.Println(ryan)

	gempong := Customer{
		Name: "Bagus",
		Address: "Subak",
		Job: "Designer",
	}
	fmt.Println(gempong)

	yayak := Customer{"Yayak", "Soputan", "Designer"}
	fmt.Println(yayak)
}