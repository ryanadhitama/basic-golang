package main

import (
	"basic-golang/helper"
	"fmt"
)

func main() {
	db := helper.GetDatabase()
	fmt.Println(db)
}