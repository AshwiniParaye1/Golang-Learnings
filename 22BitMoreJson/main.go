package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"` //aliases
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Creating Json")
	EncodeJson()
}

func EncodeJson() {
	myCourses := []course{
		{"ReactJS Bootcamp", 299, "reactjs.in", "React@12", []string{"web-dev", "react"}},
		{"VueJS Bootcamp", 199, "vuejs.in", "Vue@12", []string{"full-stack", "vue"}},
		{"Golang Bootcamp", 499, "golang.in", "Golang@12", nil},
	}

	// package this data as json data
	// finalJson, err := json.Marshal(myCourses)
	finalJson, err := json.MarshalIndent(myCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", finalJson)
}
