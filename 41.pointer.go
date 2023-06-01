package main

import "fmt"

type Address struct {
	City, Province, Nation string
}

func main() {
	address1 := Address{"Denpasar", "Bali", "Indonesia"}
	address2 := &address1
	address2.City = "Badung"

	var address3 *Address = &address1
	address3 = &Address{"Kuta", "Bali", "Indonesia"}
	address4 := &address1
	*address4 = Address{"Denpasar", "Bali", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address3)
	fmt.Println(address4)
}