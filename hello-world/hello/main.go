package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":5051", nil)
	fmt.Println("Server started on port 5051...")
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello, World\n")
}
