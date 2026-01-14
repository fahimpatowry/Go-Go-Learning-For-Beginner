package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "wecmome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for your pizza")

	// comma ok || err ok syntax
	input, _ := reader.ReadString('\n')
	// _, error := reader.ReadString('\n')
	// input, error := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
}
