package main

import "fmt"

// jwtToken := 3000 //not allow
// var jwtToken = 3000 //allow

const LoginToken string = "dfsdfsdsf" // capital first later(L) is mean public variable

func main() {
	var username string = "Fahim"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggendIn bool = false
	fmt.Println(isLoggendIn)
	fmt.Printf("Variable is of type: %T \n", isLoggendIn)

	// var smallVal uint8 = 256 //255
	var smallVal int = 256 //255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	// var smallFloat float32 = 256.4554155 //255
	var smallFloat float64 = 256.4554155 //255
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	// var anotherVariable string
	// var anotherVariable float64
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "learngo.com"
	fmt.Println(website)

	// no var style
	numberOfUser := 30000
	fmt.Println(numberOfUser)

	// print public key
	fmt.Println(LoginToken)

}
