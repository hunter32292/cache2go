package cache

import (
	"log"
	"net/http"
	"strings"
)

func InitCacheService(host string, port string) {

	// Initialize a new router service
	router := NewRouter()
	listOfStrings := []string{host, port}
	address := strings.Join(listOfStrings, ":")
	log.Fatal(http.ListenAndServe(address, router))

}
