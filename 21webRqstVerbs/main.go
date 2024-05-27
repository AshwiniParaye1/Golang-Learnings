package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Web Verb Request!")
	PerformGetRequest()
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
