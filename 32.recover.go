package main

import "fmt"

func endApp(){
	message := recover()
	if message != nil {
		fmt.Println("Error with message:", message)
	}
	fmt.Println("End Application")
}

func runApp(error bool){
	defer endApp()
	if error {
		panic("Application error")
	}
	fmt.Println("Run application")
}

func main() {
	runApp(true)
	fmt.Println("Ryan")
}