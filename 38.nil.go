package main

import "fmt"

func main() {
	var data map[string]string

	if data == nil {
		fmt.Println("Data nil")
	} else {
		fmt.Println("Data not nil")
	}
}