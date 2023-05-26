package main

import "fmt"

func endApp() {
	fmt.Println("End Application")
}

func runApplication(err bool) {
	defer endApp()
	if err {
		panic("Application error")
	}
	fmt.Println("Run application")
}

func main() {
	runApplication(true)
}