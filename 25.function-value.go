package main

import "fmt"

func getGoodBye(name string) string {
	return "Hello "+name
}

func main() {
	sayGoodBye := getGoodBye

	result := sayGoodBye("Ryan")
	fmt.Println(result)
}