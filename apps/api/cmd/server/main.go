package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/jacobbananaldev/quiklist/internal/router"
)

func main() {
	r := router.NewRouter()
	
	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
