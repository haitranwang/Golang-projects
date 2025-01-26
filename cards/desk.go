package main

import "fmt"

// Create a new type of 'desk'
// which is a slice of string
type desk []string

func (d desk) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
