package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruitList = []string{}
	fmt.Println("type of data %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana", "Mango1", "Banana1")
	fmt.Println(fruitList)

	// fruitList = append(fruitList[1:])
	fruitList = append(fruitList[1:3])
	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 245
	highScores[2] = 234
	highScores[3] = 238

	highScores = append(highScores, 555, 66, 321)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// var vagList = [3]string{"portato", "beans", "mushroom"}
	// fmt.Println(vagList)

}
