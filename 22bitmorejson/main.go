package main

import (
	"encoding/json"
	"fmt"
)

type courses struct {
	Name     string `json:"coursename"`
	price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to Json ")
	EncodeJson()
}
func EncodeJson() {
	lcoCourses := []courses{
		{"ReactJS Bootcamp1", 299, "LCO.in", "abc123", []string{"web-dev", "js"}},
		{"ReactJS Bootcamp2", 299, "LCO.in", "abc123", []string{"web-dev", "js"}},
		{"ReactJS Bootcamp3", 299, "LCO.in", "abc123", nil},
	}

	//package this data as JSON data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp1",
		"Price": 299,
		"website": "LCO.in",
		"tags": ["web-dev","js"]
	}
	`)

}
