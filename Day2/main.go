package main

import "fmt"

func main() {
	fmt.Println("Hello its day 2")

	//string
	var username string = "Ashwini"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	//bool
	var isBoolVar bool = true
	fmt.Println(isBoolVar)
	fmt.Printf("variable is of type: %T \n", isBoolVar)

	//unit
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	//float
	var smallFloat float64 = 256.123456789121212
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	//default values and some alises
	var anotherIntVar int
	fmt.Println(anotherIntVar)
	fmt.Printf("variable is of type: %T \n", anotherIntVar)

	var anotherStrVar string
	fmt.Println(anotherStrVar)
	fmt.Printf("variable is of type: %T \n", anotherStrVar)
}
