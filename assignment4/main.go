// Package main is executable package for this project.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// // call to create a slice and separate even and odd numbers into separate slices.
	// sortevenodd.PrintResult()

	// // call to create a slice and return a map with key value pairs of duplicated numbers for the slice.
	// findduplicates.PrintResult()

	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter string of your chocie: ")
	reader.Scan()
	input := strings.Replace(reader.Text(), " ", "", -1)

	fmt.Println(input)

	for i := range input {
		if input[i] != input[len(input)-i-1] {
			fmt.Println("Not a palindrome. Better luck next time!")
			return
		} else if i == (len(input) - 1) {
			fmt.Println("Hey, Nice catch! This word is a palindrome.")
		}
	}

}
