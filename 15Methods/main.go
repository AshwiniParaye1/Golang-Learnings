package main

import "fmt"

func main() {
	fmt.Println("Let's learn Methods today!")

	//no inheritance / super / prent in golang

	ashwini := User{"Ashwini", "ashwini@go.dev", true, 18}

	fmt.Println(ashwini)
	fmt.Printf("my details are: %+v\n\n", ashwini)                              //+ in %v provides more details
	fmt.Printf("Name is %v and Email is %v: \n\n", ashwini.Name, ashwini.Email) //accessing single values

	ashwini.GetStatus() //calling the method
	ashwini.NewMail()

	fmt.Printf("Name is %v and Email is %v: \n\n", ashwini.Name, ashwini.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "xyz@xyz.co"
	fmt.Println("New Email of the user is: ", u.Email)
}
