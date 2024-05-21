package main

import "fmt"

func main() {
	fmt.Println("Let's learn If-Else today!")

	loginCount := 10

	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch Out"
	} else {
		result = "exactly 10 login count"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}

	if num := 30; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is greater than 10")
	}

}
