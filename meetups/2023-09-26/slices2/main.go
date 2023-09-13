package main

import (
	"fmt"
	"slices"
)

// START OMIT
func main() {
	numbers := []int{42, 21, 0}

	numbers = slices.Delete(numbers, 1, 2) // HL

	fmt.Println(numbers)
}

// END OMIT

// RUN OMIT
// NUR OMIT
