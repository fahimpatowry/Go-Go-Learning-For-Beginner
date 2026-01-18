package main

import (
	"fmt"
	"net/url"
)

const myurl = "http://10.70.57.33:3000/test?id=4534&key=343"

func main() {
	fmt.Println(myurl)
	result, _ := url.Parse(myurl)
	fmt.Println(result)
	fmt.Println(result.Scheme) // https
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Println(qparams)
	fmt.Println(qparams["key"])
	fmt.Printf("The type of query are: %T \n", qparams)

	for _, value := range qparams {
		fmt.Println("param is : ", value)
	}

	partsOfUrls := &url.URL{
		Scheme:   "https",
		Host:     "fahim.dev",
		RawPath:  "/test",
		RawQuery: "id=dfsfs",
	}
	fmt.Println("partsOfUrls:", partsOfUrls)

	anotherUrl := partsOfUrls.String()
	fmt.Println(anotherUrl)
}
