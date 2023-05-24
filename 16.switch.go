package main

import "fmt"

func main()  {
	var name string = "Ryans"

	switch name {
		case "Bagus":
			fmt.Println("Hello, Bagus")
		case "Ryan":
			fmt.Println("Hello, Ryan")
		case "Adi":
			fmt.Println("Hello, Adi")
		default:	
			fmt.Println("Who are you?")
	}
}