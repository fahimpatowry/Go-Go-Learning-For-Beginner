package main

import (
	"fmt"
)

func main() {

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["RJS"] = "reactjs"
	languages["SW"] = "swift"
	languages["PY"] = "python"

	fmt.Println(languages)
	fmt.Println(languages["JS"])

	delete(languages, "SW")
	fmt.Println(languages)

	// loops are intersting in golang
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v \n", key, value)
		// fmt.Println(key, ":", value)
	}

}
