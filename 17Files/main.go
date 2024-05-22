package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Working with Files")

	content := "This is the content that needs to go in a file"

	//creating a file
	file, err := os.Create("./firstTextFile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	//writing into the file
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length is: ", length)
	defer file.Close()

	readFile("./firstTextFile.txt")
}

// reading a file
func readFile(filenameAlongWithPath string) {
	dataByte, err := os.ReadFile(filenameAlongWithPath)
	checkNilErr(err)

	fmt.Println("Text data inside the file is:= ", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
