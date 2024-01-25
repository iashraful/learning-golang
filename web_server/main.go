package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, r.Method+" Not Allowed", http.StatusMethodNotAllowed)
	} else if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found!", http.StatusNotFound)
	} else {
		fmt.Fprintf(w, "Hello")
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Post request data: \n")
	fmt.Fprintf(w, "Name: %v\n", name)
	fmt.Fprintf(w, "Address: %v\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Listening on port 8081.\nAccess http://0.0.0.0:8081\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
