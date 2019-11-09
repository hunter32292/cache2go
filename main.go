package main

import (
	"log"
	"net/http"
)

func main() {

	// Initialize a new router service
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}
