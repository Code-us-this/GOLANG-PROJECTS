// Package implbinarysearch used to implement binary search algorithm
package implbinarysearch

// ImplBinarySearch implements binary search on a slice for a given number
func ImplBinarySearch(sl []int, item int) (mid int) {

	low := 0
	high := len(sl) - 1

	for low <= high {

		mid = (low + high) / 2
		guess := sl[mid]

		if guess < item {
			low = mid + 1
		} else if guess > item {
			high = mid - 1
		} else {
			break
		}
	}
	return mid
}
