package main

import "fmt"

func main() {
	var data interface{} = 1
	data = "hello"
	fmt.Println(data)
}