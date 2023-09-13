package main

import "fmt"

// START OMIT
func main() { // HL12
	var x, y = 42, 21

	// RUN1 OMIT
	m := min(x)    // HL
	n := min(x, y) // HL
	o := max(x, y) // HL
	// NUR1 OMIT

	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(o)
} // HL12
//END OMIT
