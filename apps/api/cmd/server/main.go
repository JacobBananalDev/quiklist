package main

import (
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ready")
}

func main() {

	http.HandleFunc("/health", healthHandler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Issue connecting to server")
		return
	}
}
