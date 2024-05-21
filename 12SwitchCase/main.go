package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case!")

	// rand.Seed(time.Now().UnixNano())
	// diceNumber := rand.Intn(6) + 1

	//Create a new random source and random number generator
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Generate a random dice number between 1 and 6
	diceNumber := r.Intn(6) + 1
	fmt.Println("Value of a dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can Open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
	case 4:
		fmt.Println("You can move 4 spot")
	case 5:
		fmt.Println("You can move 5 spot")
	case 6:
		fmt.Println("You can move 2 spot and roll dice again!")
	}
}
