package cache

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// Status, Explain what you are, Information
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!")
}

// Dump entire in-mem values
func ValueDump(w http.ResponseWriter, r *http.Request) {
	// Set for JSON and HTTP STATUS 200
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	// Tell the http writer to encode as JSON the value struct
	// If there is an error during encoding freak the f*** out :)
	if err := json.NewEncoder(w).Encode(values); err != nil {
		panic(err)
	}
}

//Find a specific value
func ValueFind(w http.ResponseWriter, r *http.Request) {
	// Get Reqest Varaibles
	vars := mux.Vars(r)
	var value Value
	value = RepoFindValue(vars["value"])
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		panic(err)
	}
}

// Add a new value
func ValueAdd(w http.ResponseWriter, r *http.Request) {
	var value Value
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &value); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	t := RepoCreateValue(value)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

// Remove a  value
func ValueRemove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var err error
	if err = RepoDestroyValue(vars["value"]); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Failed to Delete Value")
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Value Deleted")
}
