// Package sortevenodd helps to sort the even and odd numbers of given slice of int data type
package sortevenodd

import (
	userslice "assignment4/userSlice"
	"fmt"
)

// SortEvenOdd function sorts and returns two separate slices with even and odd numbers respectively
func sortEvenOdd(s []int) (e []int, o []int) {

	for i := 0; i < len(s); i++ {

		if s[i]%2 == 0 {
			e = append(e, s[i])
		} else {
			o = append(o, s[i])
		}

	}

	return e, o

}

// PrintResult in sortevenodd package prints the resulting slices containing even and odd numbers.
func PrintResult() {

	var slice []int
	evenSlice, oddSlice := sortEvenOdd(userslice.CreateAnIntSlice(slice))

	fmt.Println("The slice with even numbers is:", evenSlice, "\nThe slice with odd numbers is:", oddSlice)

}
