package main

import (
	"fmt"
	"net/http"
	"github.com/jacobbananaldev/quiklist/internal/handlers"
)

func main() {

	http.HandleFunc("/health", handlers.HealthHandler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Issue connecting to server")
		return
	}
}
