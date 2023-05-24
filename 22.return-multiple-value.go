package main

import "fmt"

func getFullName() (string, string) {
	return "Ryan", "Adhitama"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)
}