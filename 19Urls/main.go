package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "http://example.com:8080/api/resource?key1=value1&key2=value2"

func main() {
	fmt.Println("Handling URLs in Golang")
	fmt.Println(myUrl)

	//parsing the url
	result, _ := url.Parse(myUrl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())  ----- Port() is not a property, it is a method
	fmt.Println(result.RawQuery)

	qParams := result.Query()
	fmt.Printf("Type of Query params are: %T\n", qParams)

	fmt.Println(qParams["key2"])

	for _, val := range qParams {
		fmt.Println("Params are: ", val)
	}

	//constructing a url from parts
	partsOfUrl := &url.URL{ //have to pass a reference
		Scheme: "https",
		Host:   "medium.com",
		Path:   "@ashwini-paraye",
		// RawQuery: "user=ashwini",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

	// https://medium.com/@ashwini-paraye
}
