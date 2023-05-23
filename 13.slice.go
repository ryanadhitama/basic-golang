package main

import "fmt"

func main() {
	var months = [...]string{
		"Jan",
		"Feb",
		"Mar",
		"Apr",
		"May",
		"Jun",
	}
	fmt.Println(months)
	
	var slice1 = months[2:4]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	thisArray := [3]int{1,2,3}
	thisSlice := []int{1,2,3}
	fmt.Println(thisArray)
	fmt.Println(thisSlice)
}