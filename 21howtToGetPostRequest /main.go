package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const myurl = "http://10.70.57.33:3000/test?id=4534&key=343"

func main() {
	PerformJsonRequest()
}

func PerformJsonRequest() {
	const myurl = "http:localhost:8000/post"

	// fake form data
	data := url.Values{}
	data.Add("firstname", "fahim")
	data.Add("lastname", "pat")
	data.Add("email", "pat.com")

	// fake json payload

	requestBody := strings.NewReader(`
	{
		"coursename":"let's go with golan",
		"price": 0,
		"platfrom": "fahim.com"
	}`)

	res, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(content))
}
