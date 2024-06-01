package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Mod in Golang")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")

	//running this as a server
	log.Fatal(http.ListenAndServe(":4000", r))
}
func greeter() {
	fmt.Println("Hello from Greeter!")
}
func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello from RESPONSE WRITER</h1>"))
}
