package main

import "fmt"

func main() {
	result := User{"fahim", 10, true}
	// fmt.Println(result)

	// result.GetStatus()
	fmt.Println("Is user age 1:", result.age)
	result.NewMail()

	fmt.Println(result)
}

type User struct {
	name   string
	age    int
	status bool
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.status)
}

func (u User) NewMail() {
	u.age = 30
	// fmt.Println("Is user age:", u.age)
}
