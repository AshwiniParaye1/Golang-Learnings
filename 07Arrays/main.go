package main

import "fmt"

func main() {
	fmt.Println("let's learn Arrays today!")

	var fruitList [4]string

	fruitList[0] = "mango"
	fruitList[1] = "apple"
	fruitList[3] = "banana"

	fmt.Println("fruitlist : ", fruitList)
	fmt.Println("fruitlist : ", len(fruitList))

	var veggiesList = [4]string{"beans", "tomato"}

	fmt.Println("veggiesList : ", veggiesList)
	fmt.Println("veggiesList : ", len(veggiesList))
}
