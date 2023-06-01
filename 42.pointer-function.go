package main

import "fmt"

type Address struct {
	City, Province, Nation string
}

func changeNationToMalay(address *Address) {
	address.Nation = "Malay"
}

func main() {
	address1 := Address{"Denpasar", "Bali", "Indonesia"}
	changeNationToMalay(&address1)
	
	fmt.Println(address1)
}