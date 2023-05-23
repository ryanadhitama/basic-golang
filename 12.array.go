package main

import "fmt"

func main()  {
	var name [3]string

	name[0] = "Ryan"
	name[1] = "Adhitama"
	name[2] = "Putra"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var value = [3]int{
		1,10,20,
	}

	fmt.Println(value[0])
	fmt.Println(value[1])
	fmt.Println(value[2])
}