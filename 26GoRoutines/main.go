package main

import (
	"fmt"
	"net/http"
	"sync"
)

// must be pointers
var wg sync.WaitGroup

func main() {
	// go greeter("Hello")
	// greeter("World")

	websiteList := []string{
		"https://stackoverflow.com/",
		"https://soundcloud.com",
		"https://www.dropbox.com/",
		"https://www.bbc.com/",
		"https://medium.com/",
		"https://www.linkedin.com/",
		"https://github.com/",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {

	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	}
	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
}
