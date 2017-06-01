package main

import "time"

type Value struct {
	Id      int       `json:"id"`
	Value   string    `json:"value"`
	Created time.Time `json:"create"`
}

type Values []Value
type ValuesMap map[string]Value
