package main

import "fmt"

func main() {
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	// fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for _, day := range days {
	for index, day := range days {
		// if i == 2 {
		// 	break
		// }
		fmt.Printf("index is %v and value is %v \n", index, day)
	}

	// like while loop
	rougeValue := 1
	for rougeValue < 10 {
		if rougeValue == 3 {
			goto lco
		}
		// if rougeValue == 4 {
		// 	rougeValue++
		// 	break
		// }
		// if rougeValue == 5 {
		// 	rougeValue++
		// 	continue
		// }
		fmt.Println("value is: ", rougeValue)
		rougeValue++
	}

lco:
	fmt.Println("jump")

}
