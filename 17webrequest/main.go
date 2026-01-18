package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://chatgpt.com"

func main() {

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("error!!!!!!!!!!")
		panic(err)
	}
	fmt.Printf("Response is of type: %T \n", response)

	defer response.Body.Close()

	dataByte, err1 := ioutil.ReadAll(response.Body)
	if err1 != nil {
		fmt.Println("error!!!!!!!!!!")
		panic(err)
	}
	content := string(dataByte)
	fmt.Printf(content)

}
