// Package userslice helps scanning user input and convert it into a Slice
package userslice

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CreateAnIntSlice function helps creating a new slice
func CreateAnIntSlice(sl []int) []int {

	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter elements of the slice (Seperated by spaces): ")
	reader.Scan()
	userInput := strings.Split(reader.Text(), " ")

	for _, i := range userInput {
		num, err := strconv.Atoi(i)
		if err != nil {
			fmt.Printf("Error in converting '%s' to int: %v\n", i, err)
			continue
		}
		sl = append(sl, num)
	}

	return sl

}
