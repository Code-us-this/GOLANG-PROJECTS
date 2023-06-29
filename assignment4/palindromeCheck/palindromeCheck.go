// Package palindromecheck helps checking if the user input string is a palindrome or not
package palindromecheck

import (
	userinput "assignment4/userInput"
	"fmt"
	"strings"
)

// Palindromecheck function checks if the user input string is a palindrome or not
func Palindromecheck() {

	input := strings.Replace(userinput.CreateAString(), " ", "", -1)

	for i := range input {
		if input[i] != input[len(input)-i-1] {
			fmt.Println("\n", false, ": Not a palindrome. Better luck next time!")
			return
		} else if i == (len(input) - 1) {
			fmt.Println("\n", true, ": Hey, Nice catch! This word is a palindrome.")
		}
	}

}
