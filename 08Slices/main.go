package main

import (
	"fmt"
)

func main() {
	fmt.Println("let's learn Slices today!")

	// var fruitList = []string{"mango", "tomato", "apple", "grapes", "cherry"}

	// fmt.Printf("type of fruit list is %T\n", fruitList)

	// fruitList = append(fruitList, "grapes", "xyzFruit")

	// fmt.Println(fruitList)

	// fruitList = append(fruitList[1:])

	// fmt.Println(fruitList)

	// fruitList = append(fruitList[1:3])

	// fmt.Println(fruitList)

	// fruitList = append(fruitList[:1])

	// fmt.Println(fruitList)

	// highScores := make([]int, 4)

	// highScores[0] = 234
	// highScores[1] = 324
	// highScores[2] = 4343
	// highScores[3] = 298
	// // highScores[4] = 228

	// fmt.Println(highScores)

	// highScores = append(highScores, 444, 555, 666, 888)

	// fmt.Println(highScores)

	// fmt.Println(sort.IntsAreSorted(highScores))

	// sort.Ints(highScores)
	// fmt.Println(highScores)

	// fmt.Println(sort.IntsAreSorted(highScores))

	//todo removing value  from slice based on index

	courses := []string{"reactjs", "javascript", "swift", "golang", "python", "ruby"}

	fmt.Println(courses)

	index := 2

	courses = append(courses[:index], courses[index+1:]...)
	//courses[:index]= 0-1 index        //courses[index+1:] = 3-last index

	fmt.Println(courses)
}
