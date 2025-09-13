package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/hello", hello)

	portNumber := os.Getenv("PORT")

	if portNumber == "" {
		portNumber = "5051"
	}

	fmt.Printf("Listening on port %s...\n", portNumber)
	http.ListenAndServe(fmt.Sprintf(":%s", portNumber), nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello")
}
