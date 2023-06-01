package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Ryan Adhitama", "Eko"))
	fmt.Println(strings.Contains("Ryan Adhitama", "Budi"))

	fmt.Println(strings.Split("Ryan Adhitama", " "))

	fmt.Println(strings.ToLower("Ryan Adhitama"))
	fmt.Println(strings.ToUpper("Ryan Adhitama"))
	fmt.Println(strings.ToTitle("Ryan Adhitama"))

	fmt.Println(strings.Trim("      Ryan Adhitama     ", " "))
	fmt.Println(strings.ReplaceAll("Ryan Joko Ryan", "Ryan", "Budi"))
}