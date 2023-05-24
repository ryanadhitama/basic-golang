package main

import "fmt"

func getName() (firstName, lastName string) {
	firstName = "ryan"
	lastName = "adhitama"

	return
}

func main() {
	a, b := getName()

	fmt.Println(a)
	fmt.Println(b)
}