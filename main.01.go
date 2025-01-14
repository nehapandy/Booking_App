package main

import (
	"BOOKING-APP/Router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	Router.Initialize() // Calling a function from router.go

	// Initialize and start the server
	fmt.Println("MongoDB API")
	r := Router.Router()

	fmt.Println("Server is starting on port 4000...")
	log.Fatal(http.ListenAndServe(":4000", r))
}
