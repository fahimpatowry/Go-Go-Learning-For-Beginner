package main

import "fmt"

func main() {
	defer fmt.Println("w1")
	defer fmt.Println("w2")
	defer fmt.Println("w3")

	fmt.Println("result")
	muDefer()
}

func muDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
