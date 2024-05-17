package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	// comma ok or comma err syntax
	input, _ := reader.ReadString('\n') //read until there is \n
	fmt.Println("thanks for rating ", input)
	fmt.Printf("type of rating is = %T", input)

	// _, err     - here we care about the error
	// input, err - here we handle input as well as err just like try catch
}
