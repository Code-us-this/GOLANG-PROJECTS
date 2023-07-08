// Package main is executable package for this project.
package main

import (
	findcharacterindex "assignment4/findCharacterIndex"
	findduplicates "assignment4/findDuplicates"
	palindromecheck "assignment4/palindromeCheck"
	sortevenodd "assignment4/sortEvenOdd"
	"bufio"
	"fmt"
	"os"
)

func main() {

	inMenu := true

	for inMenu {

		reader := bufio.NewScanner(os.Stdin)

		fmt.Println("\nWelcome, we have some fun apps below. select one to begin.\n\nType 1 for sorting even and odd numbers.\nType 2 for finding duplicate numbers.\nType 3 for checking if a word is palindrome or not.\nType 4 to check the middle index of given character")

		reader.Scan()
		input := reader.Text()

		switch option := input; option {
		case "1":
			// call to create a slice and separate even and odd numbers into separate slices.
			sortevenodd.PrintResult()
		case "2":
			// call to create a slice and return a map with key value pairs of duplicated numbers for the slice.
			findduplicates.PrintResult()
		case "3":
			// call to create a string and check if the given input is a palindrome or not.
			palindromecheck.Palindromecheck()
		case "4":
			// call to create a string, find and print middle index of the given character of a string.
			findcharacterindex.FindCharacterIndex()
		default:
			fmt.Println("\nPut on some glasses and try again!")
		}

		newReader := bufio.NewScanner(os.Stdin)

		fmt.Println("\nDo you want to continue Type Y for yes, N for no: ")

		newReader.Scan()
		isContinue := newReader.Text()

		if isContinue != "y" && isContinue != "Y" {
			inMenu = false
		}

	}
}
