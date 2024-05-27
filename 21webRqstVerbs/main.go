package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web Verb Request!")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code is: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	//! way 1 to read the response
	// content, _ := io.ReadAll(response.Body)

	// fmt.Println(string(content))

	//! way 2 to read the response
	var responseStringBuiler strings.Builder

	content, _ := io.ReadAll(response.Body)

	byteCount, _ := responseStringBuiler.Write(content)

	fmt.Println("Byte Count is: ", byteCount)
	fmt.Println(responseStringBuiler.String())

}

func PerformPostJsonRequest() {
	const postUrl = "http://localhost:8000/post"

	//dummy Json
	requestBody := strings.NewReader(`
	{
		"courseName" : "Let's Go with Golang",
		"platform" : "example.com",
		"price" : 0
	}
	`)

	response, err := http.Post(postUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const postFormUrl = "http://localhost:8000/postform"

	//formData

	data := url.Values{}
	data.Add("firstName", "Ashwini")
	data.Add("lastName", "Paraye")
	data.Add("email", "ashwiniparaye1@gmail.com")

	response, err := http.PostForm(postFormUrl, data)
	if err != nil {
		panic(err)
	}

	fmt.Println(response)

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))

}
