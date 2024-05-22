package main

import "fmt"

func main() {
	fmt.Println("Functions!")
	greeter() //function call

	// //function inside function is not allowed
	// func greeterTwo() {
	// 	fmt.Println("Greeter function 2!")
	// }

	// greeterTwo()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proResult, sm := proAdder(2, 2, 2, 2, 2, 2)
	fmt.Println("Proresult is: ", proResult, "\n", sm)
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}

	return total, "something unnecessary"
}

func adder(v1 int, v2 int) int {
	return v1 + v2
}

func greeter() {
	fmt.Println("Greeter function!")
}
