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
	// EncodeJson()
	DecodeJson()
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

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"courseName": "ReactJS Bootcamp",
		"Price": 299,
		"website": "reactjs.in",
		"tags": ["web-dev", "react"]
	}
	`)

	var myCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json is Valid")
		json.Unmarshal(jsonDataFromWeb, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("JSON IS NOT Valid")
	}

	//some cases where we just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v And Type is %T\n", k, v, v)
	}
}
