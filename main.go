package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to my Home Page server!!")
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {

	exampleUser := User{}

	err := json.NewDecoder(r.Body).Decode(&exampleUser)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(exampleUser)

}

func GetHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "Application/json")

	exampleUser := User{
		Name: "Shubham Raj",
		Age:  22,
	}

	err := json.NewEncoder(w).Encode(exampleUser)

	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	fmt.Println("Server is up and running")

	http.HandleFunc("/", HomePageHandler)
	http.HandleFunc("/post-user", PostHandler)
	http.HandleFunc("/get-user", GetHandler)

	http.ListenAndServe("localhost:8080", nil)
}
