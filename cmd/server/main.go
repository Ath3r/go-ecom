package main

import (
	"fmt"
	"net/http"

	"github.com/Ath3r/go-ecom.git/internal/routes"
)

func main() {
	server := routes.CreateNewServer()

	server.MountHandlers()
	
	fmt.Println("Starting to listen on port 8080")
	if err := http.ListenAndServe(":8080", server.Router); err != nil {
		fmt.Printf("Error starting server: %v", err)
	}
}