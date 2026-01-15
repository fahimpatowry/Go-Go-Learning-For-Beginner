package main

import "fmt"

func main() {

	// var one int = 2
	// var ptr *int
	// fmt.Println(ptr)

	myNumber := 23
	var ptr = &myNumber // reference means &
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = *ptr * 2
	fmt.Println("new value is:", myNumber)

}
