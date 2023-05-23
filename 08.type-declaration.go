package main

import "fmt"

func main() {
	type IdentityNumber string

	var noIdentity IdentityNumber = "099878978979"
	fmt.Println(noIdentity)
}