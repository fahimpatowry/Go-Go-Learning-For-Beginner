package main

import "fmt"

func main() {
	greeter()

	result := adder(3, 5)
	fmt.Println(result)

	proResult, t := proAdder(2, 15, 3, 80, 7)
	fmt.Println(proResult)
	fmt.Println(t)

	//can't do this
	// func greeterTow() {
	// 	fmt.Println("hit this is 2")
	// }
}

func greeter() {
	fmt.Println("Hello!")
}

func adder(valueOne int, valueTow int) int { //int is function signeture
	return valueOne + valueTow
}

func proAdder(values ...int) (int, string) { //int is function signeture
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "hi"
}
