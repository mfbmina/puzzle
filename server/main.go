package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("starting server: http://localhost:8080")

	err := http.ListenAndServe(":8080",
		http.FileServer(http.Dir("./assets")))
	if err != nil {
		panic(fmt.Errorf("Failed to start server: %w", err))
	}
}
