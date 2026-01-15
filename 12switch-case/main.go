package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	// fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice move to 1 and you can open")
	case 2:
		fmt.Println("Dice move to 2 spot")
		fallthrough // will consider next case also
	case 3:
		fmt.Println("Dice move to 3 spot")
		fallthrough // will consider next case also
	case 4:
		fmt.Println("Dice move to 4 spot")
	case 5:
		fmt.Println("Dice move to 5 spot")
	case 6:
		fmt.Println("Dice move to 6 spot")
	default:
		fmt.Println("What was that!")
	}
}
