package main

import "fmt"

func sayHelloFilter(name string, filter func(string) string) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Dog" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloFilter("Ryan", spamFilter)

	filter := spamFilter

	sayHelloFilter("Dog", filter)
}