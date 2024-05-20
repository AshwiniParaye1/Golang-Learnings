package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let's learn Maps today!")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	// fmt.Println(languages)
	// fmt.Printf("type %T\n", languages)
	// fmt.Println("Accesing one value with key: ", languages["JS"])

	// //deleting value from maps

	// delete(languages, "RB")

	fmt.Println(languages)

	//todo -  Looping over maps

	// for key, value := range languages {
	// 	fmt.Printf("for key %v, value is %v\n", key, value)
	// }

	// //comma, ok syntax - skipping the key
	// for _, value := range languages {
	// 	fmt.Printf("for key xyz, value is %v\n", value)
	// }

	//comma, ok syntax - skipping the value
	for key, _ := range languages {
		fmt.Printf("for key %v, value is xyz \n", key)
	}

}
