package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:3000"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code:", response.Status)
	fmt.Println("Content length: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Bytecount: ", byteCount)
	fmt.Println(responseString.String())
	// fmt.Println(string(content))
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:3000/post"

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename" : "Let's go with golang",
			"price" : 0,
			"platform" : "LCO.in"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:3000/postform"

	// formdata

	data := url.Values{}
	data.Add("firstname", "bhavesh")
	data.Add("lastname", "gujar")
	data.Add("email", "bhavesh@go.dev")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
