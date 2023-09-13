package main

import (
	"fmt"
	"slices"
)

// START OMIT
func main() {
	numbers := []int{42, 21, 0}

	i := slices.Index(numbers, 42)   // HL
	j := slices.Index(numbers, 0)    // HL
	k := slices.Index(numbers, 1111) // HL

	fmt.Println(i, j, k)
}

// END OMIT

// RUN OMIT
// NUR OMIT
