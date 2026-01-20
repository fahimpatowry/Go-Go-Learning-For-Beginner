package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"` // - will decide this will not show in json
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {

	lcoCourses := []course{
		{"ReactJS Bootcamp", 255, "fahim.com", "abc12", []string{"web-dev", "js"}},
		{"MEARN Bootcamp", 450, "fahim.com", "abc12", []string{"web-full", "js"}},
		{"ANGULER Bootcamp", 50, "fahim.com", "abc12", nil},
	}

	// package this data as JSON data
	// interface
	// finalJson, err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", finalJson)

}

func DecodeJson() {
	jsonDateFromWeb := []byte(`
	{
		"coursename": "MEARN Bootcamp",
		"price": 450,
		"website": "je.com",
		"tags": ["fahim.com","js"]
	}
	`)

	var lcoCourse course

	checker := json.Valid(jsonDateFromWeb)

	if checker {
		fmt.Println("JSON was valid:")
		json.Unmarshal(jsonDateFromWeb, &lcoCourse)
		fmt.Printf("%#v \n", lcoCourse)
	} else {
		fmt.Println("water have error")
	}

	// some case where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDateFromWeb, &myOnlineData)

	fmt.Printf("%#v \n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value %v and type is: %T \n", k, v, v)
	}
}
