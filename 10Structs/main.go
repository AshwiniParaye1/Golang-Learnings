package main

import "fmt"

func main() {
	fmt.Println("Let's learn Structs today!")

	//no inheritance / super / prent in golang

	ashwini := User{"Ashwini", "ashwini@go.dev", true, 18}

	fmt.Println(ashwini)
	fmt.Printf("my details are: %+v\n", ashwini)                            //+ in %v provides more details
	fmt.Printf("Name is %v and Email is %v: ", ashwini.Name, ashwini.Email) //accessing single values
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
