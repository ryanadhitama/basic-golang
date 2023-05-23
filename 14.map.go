package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Ryan",
		"job": "Programmer",
	}

	fmt.Println(person["name"])
	fmt.Println(person["job"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Go Advanced"
	book["author"] = "Ryan"

	fmt.Println(book)

	delete(book, "author")
	fmt.Println(book)
}