package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.September, 10, 23, 20, 0, 0, time.UTC)
	fmt.Println("createdDate:", createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))

	// presentTime2 := time.Now().Nanosecond()
}
