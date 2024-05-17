package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study in golang")
	presentTime := time.Now()
	fmt.Print("full Date: ", presentTime)
	fmt.Print("\nformated Time: ", presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2024, time.January, 23, 18, 0, 0, 0, time.Local)
	fmt.Print("\ncreatedDate: ", createdDate.Format("01-02-2006 Monday"))
}
