// Package findcharacterindex helps finds and prints middle index of the given character of a string
package findcharacterindex

import (
	userinput "assignment4/userInput"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// FindCharacterIndex function find and prints middle index of the given character of a string
func FindCharacterIndex() {

	input := strings.ToLower(userinput.CreateAString())

	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter a single character for which you want to find its middle index: ")
	reader.Scan()
	char := strings.ToLower(reader.Text())

	charMap := make(map[int]int)
	key := 0

	for i := range input {
		if string(input[i]) == char {
			key++
			charMap[key] = i
		}
	}

	if len(charMap)%2 == 0 && len(charMap) != 0 {
		fmt.Println("The middle index of given character in the given string is: ", charMap[len(charMap)/2])
		fmt.Println(input)
	} else {

		if len(charMap) == 0 {
			fmt.Println("The given character doesnt exist in string or the given character is not a single character.")
		} else {
			fmt.Println("The middle index of given character in the given string is: ", charMap[(len(charMap)+1)/2])
		}

	}
}
