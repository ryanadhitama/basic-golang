package main

import "fmt"

func getHello(name string) string {
	return "Hello, "+name
}

func main()  {
	name := getHello("Ryan")
	fmt.Println(name)
}