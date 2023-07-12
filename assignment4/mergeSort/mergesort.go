// Package mergesort helps to sort a given slice.
package mergesort

import (
	userinput "assignment4/userInput"
	"fmt"
)

// mergesort function helps to sort a given slice
func mergeSort(slice []int) {

	length := len(slice)

	if length < 2 {
		return
	}

	middle := length / 2
	leftSlice := make([]int, middle)
	rightSlice := make([]int, length-middle)

	i := 0
	j := 0

	for ; i < length; i++ {
		if i < middle {
			leftSlice[i] = slice[i]
		} else {
			rightSlice[j] = slice[i]
			j++
		}
	}
	mergeSort(leftSlice)
	mergeSort(rightSlice)
	merge(leftSlice, rightSlice, slice)
}

func merge(leftSlice []int, rightSlice []int, slice []int) {

	leftSize := len(slice) / 2
	rightSize := len(slice) - leftSize
	i, l, r := 0, 0, 0 //indices

	//check the conditions for merging
	for l < leftSize && r < rightSize {
		if leftSlice[l] < rightSlice[r] {
			slice[i] = leftSlice[l]
			i++
			l++
		} else {
			slice[i] = rightSlice[r]
			i++
			r++
		}
	}

	for l < leftSize {
		slice[i] = leftSlice[l]
		i++
		l++
	}

	for r < rightSize {
		slice[i] = rightSlice[r]
		i++
		r++
	}

}

// PrintResult in mergesort package prints the resulting sorted slice for a given slice.
func PrintResult() {

	var slice []int
	slice = userinput.CreateAnIntSlice(slice)
	mergeSort(slice)

	fmt.Println("\nThe sorted slice is: ", slice)

}
