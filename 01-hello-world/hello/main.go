package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/secrets", secrets)

	http.ListenAndServe(":5051", nil)
	fmt.Println("Server started on port 5051...")
}

func hello(w http.ResponseWriter, req *http.Request) {
	message := os.Getenv("message")
	fmt.Fprintf(w, "Message: %s\n", message)
}

func secrets(w http.ResponseWriter, req *http.Request) {
	username := os.Getenv("username")
	password := os.Getenv("password")
	fmt.Fprintf(w, "Username: %s, Password: %s\n", username, password)
}
