package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Ryan")
	data.PushBack("Adhitama")
	data.PushBack("Putra")
	data.PushFront("Budi")

	// from front to back
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// from back to front
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}