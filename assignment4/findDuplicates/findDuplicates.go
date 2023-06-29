// Package findduplicates helps to find the duplicates of a given slice
package findduplicates

import (
	userinput "assignment4/userInput"
	"fmt"
)

// fundDuplicates function returns a map with duplicated values for a given slice.
func findDuplicates(sl []int) map[int]int {

	// cannot add values into uninitialized maps hence initialize an empty map of int-int pairs
	m := make(map[int]int)

	// check the slice and assign count value in map for corresponding number
	for _, j := range sl {

		m[j]++

		// 👇👇👇Alternative version which works too.

		// _, exist := m[j]

		// if exist == false {
		// 	m[j] = 1
		// } else {
		// 	m[j]++
		// }
	}

	// Iterate through the map and delete the key value pairs with no duplicate count
	for i, k := range m {
		if k < 2 {
			delete(m, i)
		}
	}

	return m

}

// PrintResult function in the findduplicates package prints the resulting map containing key value pairs with duplicate count.
func PrintResult() {

	var slice []int
	fmt.Println("This map contains key value pairs of duplicated numbers in the format of {number: count}: ", findDuplicates(userinput.CreateAnIntSlice(slice)))

}
