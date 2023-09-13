package main

import (
	"fmt"
)

func main() {
	s := []string{"Gophers", "at", "FER", "at", "meetup"}
	clear(s) // HL

	fmt.Printf("len: %d\n", len(s))

	m := map[string]string{"location": "FER", "date": "2023-09-26"}
	clear(m) // HL
	fmt.Printf("len: %d\n", len(m))
}

// RUN OMIT
// NUR OMIT
