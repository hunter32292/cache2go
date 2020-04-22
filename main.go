package main

import (
	"log"
	"net/http"

	cache "github.com/hunter32292/cachego/pkg"
)

func main() {

	// Initialize a new router service
	router := cache.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}
