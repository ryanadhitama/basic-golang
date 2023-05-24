package main

import "fmt"

func main()  {
	for i := 1; i <= 10; i++ {
		// stop current looping to next index
		if i == 3 {
			continue
		}

		// stop looping
		if i == 9 {
			break
		}

		fmt.Println("loop:", i)
	}
}