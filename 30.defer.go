package main

import "fmt"

func logging() {
	fmt.Println("Logging")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println(result)
}

func main() {
	runApplication(0)
}