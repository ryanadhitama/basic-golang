package main

import "fmt"

func main()  {
	counter := 1
	
	for counter <= 10 {
		fmt.Println("loop:", counter)
		counter++
	}

	for i := 0; i <= 5; i++ {
		fmt.Println("loop i:", i)
	}

	// range
	slice := []string{"Bagus", "Adit", "Faris"}

	for i, s := range slice {
		fmt.Println("name", i, "=",s)
	}

	exMap := make(map[string]string)
	exMap["name"] = "Ryan"
	exMap["job"] = "Programmer"

	for key, value := range exMap {
		fmt.Println(key, "=", value)
	}
}