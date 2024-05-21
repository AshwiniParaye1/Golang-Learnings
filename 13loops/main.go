package main

import "fmt"

func main() {
	fmt.Println("Oops! Loops!")

	days := []string{"Sunday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, days := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, days)
	// }

	rougueValue := 1

	// for rougueValue < 10 {
	// 	fmt.Println("value is: ", rougueValue)
	// 	rougueValue++
	// }

	// for rougueValue < 10 {
	// 	if rougueValue == 5 {
	// 		rougueValue++
	// 		break
	// 	}
	// 	fmt.Println("value is: ", rougueValue)
	// 	rougueValue++
	// }

	for rougueValue < 10 {

		if rougueValue == 3 {
			goto lco
		}

		if rougueValue == 5 {
			rougueValue++
			continue
		}
		fmt.Println("value is: ", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jumping at Goto statement!")

}
