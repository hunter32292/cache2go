package main

import (
	"log"
	"net/http"

	"github.com/hunter32292/cachego/pkg/cache"
)

func main() {

	// Initialize a new router service
	router := cache.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}
