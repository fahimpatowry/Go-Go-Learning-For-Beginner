package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const myurl = "http://10.70.57.33:3000/test?id=4534&key=343"

func main() {
	performGetRequest()
}

func performGetRequest() {
	const myurl = "http://10.70.57.33:3000"

	res, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("Status code:", res.StatusCode)
	fmt.Println("Content length:", res.ContentLength)

	// content, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(string(content))

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(res.Body)

	byteCount, _ := responseString.Write(content)
	fmt.Println(string(byteCount))
	fmt.Println(responseString.String())

}
