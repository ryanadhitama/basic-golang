package main

import "fmt"

func main()  {
	var a = 100
	var b = 150

	fmt.Println(a,">",b,"=",a>b)
	fmt.Println(a,">=",b,"=",a>=b)
	fmt.Println(a,"<",b,"=",a<b)
	fmt.Println(a,"<=",b,"=",a<=b)
	fmt.Println(a,"==",b,"=",a==b)
	fmt.Println(a,"!=",b,"=",a!=b)
}