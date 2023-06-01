package main

import (
	"fmt"
	"errors"
)

func divide(nilai int, div int) (int, error) {
	if div == 0 {
		return 0, errors.New("Divider should not 0")
	} else {
		return nilai/div, nil
	}
}

func main() {
	var exampleError = errors.New("New error")
	fmt.Println(exampleError.Error())

	hasil, err := divide(10, 2)
	if err == nil {
		fmt.Println("Hasil:", hasil)
	} else {
		fmt.Println("Error:", err.Error())
	}
}