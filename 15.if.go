package main

import "fmt"

func main() {
	var name string = "Adit"

	if name == "Ryan" {
		fmt.Println("Hello", name)
	} else {
		fmt.Println("You are not Ryan, but", name)
	}
}