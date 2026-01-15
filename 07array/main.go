package main

import "fmt"

func main() {

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var vagList = [3]string{"portato", "beans", "mushroom"}
	fmt.Println(vagList)

}
