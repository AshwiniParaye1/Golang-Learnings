package main

import "fmt"

func main() {
	// fmt.Println("let's learn Pointers today!")

	// var ptr *int

	// fmt.Println("value of pointer is: ", ptr)

	myNumber := 22

	var newPtr = &myNumber

	// fmt.Println("value of new pointer is: ", newPtr)
	// fmt.Println("value of new pointer is: ", *newPtr)

	*newPtr = *newPtr + 2
	fmt.Println("value of new pointer is: ", myNumber)
}
