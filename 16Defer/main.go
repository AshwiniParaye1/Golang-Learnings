package main

import "fmt"

func main() {
	// fmt.Println("Let's learn Defer today!")

	defer fmt.Println("World!")
	defer fmt.Println("One")
	defer fmt.Println("\nTwo")
	fmt.Println("Hello")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
