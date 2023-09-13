package main

import "fmt"

// START OMIT
func main() { // HL
	var x, y = 42, 21

	// RUN OMIT
	m := min(x, y, 100) // HL
	// NUR OMIT

	fmt.Println(m)
} // HL
// END OMIT
