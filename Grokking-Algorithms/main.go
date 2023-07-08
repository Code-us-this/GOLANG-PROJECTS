// Package main is executable package for this project.
package main

import (
	implbinarysearch "algorithms/binarySearch"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var slice []int
	slice = createAnIntSlice(slice)

	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("\nEnter the number you want to implement binary search for")
	reader.Scan()
	input := reader.Text()

	searchNumber, _ := strconv.Atoi(input)

	pos := implbinarysearch.ImplBinarySearch(slice, searchNumber)
	fmt.Println("index of the number you searched for is:", pos)

}

// CreateAnIntSlice function creates a new slice based on user input.
func createAnIntSlice(sl []int) []int {

	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("\nEnter the upper limit number ")
	reader.Scan()
	input := reader.Text()

	upperLimit, _ := strconv.Atoi(input)

	for i := 1; i <= upperLimit; i++ {
		sl = append(sl, i)
	}

	return sl

}
