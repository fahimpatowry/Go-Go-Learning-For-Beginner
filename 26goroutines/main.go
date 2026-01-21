package main

import (
	"fmt"
	"time"
)

func main() {
	go greeter("Hello")
	greeter("World")
	greeter("111")
	greeter("!!!!!")
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
