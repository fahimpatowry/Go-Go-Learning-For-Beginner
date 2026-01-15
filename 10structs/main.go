package main

import "fmt"

func main() {

	// no inheritance in golang; No Super or parent

	fahim := User{"fahim", "fahim@go.dev", true, 21}
	fmt.Println(fahim)
	fmt.Printf("fahim details are: %+v \n", fahim)
	fmt.Printf("Name is: %v \n", fahim.Name)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
