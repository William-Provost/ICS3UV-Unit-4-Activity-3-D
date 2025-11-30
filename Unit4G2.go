// Author: William Provost
// Version: 1.0.0
// Date: 2025-11-30
// Fileoverview: This program uses indecies in an array.

package main

import "fmt"

func main() {
	// constants and variables
	// Declares an array of 4 float64 elements
	var decimalValues [4]float64
	var index int = 0

	// input (Initialization of array elements)
	decimalValues[0] = 10.5
	decimalValues[1] = 20.5
	decimalValues[2] = 30.5
	decimalValues[3] = 40.5

	// output
	// set index variable to 3 so you can see
	// how to use a variable for an index value in an array
	index = 3

	// Prints decimalValues[3]
	fmt.Printf("Cell 3: %.2f\n", decimalValues[index])
	// Prints decimalValues[3 - 1] = decimalValues[2]
	fmt.Printf("Cell 2: %.2f\n", decimalValues[index-1])

	// index = 3 - 2 = 1
	index = index - 2

	// Prints decimalValues[1]
	fmt.Printf("Cell 1: %.2f\n", decimalValues[index])
	// Prints decimalValues[1 - 1] = decimalValues[0]
	fmt.Printf("Cell 0: %.2f\n", decimalValues[index-1])

	fmt.Println("\nDone.")
}
