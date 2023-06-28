// Package main is executable package for this project.
package main

import (
	findduplicates "assignment4/findDuplicates"
	sortevenodd "assignment4/sortEvenOdd"
)

func main() {

	// call to create a slice and separate even and odd numbers into separate slices.
	sortevenodd.PrintResult()

	// call to create a slice and return a map with key value pairs of duplicated numbers for the slice.
	findduplicates.PrintResult()

}
