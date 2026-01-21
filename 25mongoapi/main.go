package main

import (
	"fmt"
	"mongoapi/router"
	"net/http"
)

func main() {
	fmt.Println("hi")

	r := router.Router()
	http.ListenAndServe(":4000", r)

}
