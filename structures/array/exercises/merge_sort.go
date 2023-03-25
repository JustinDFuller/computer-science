// Merge Sort Array
package exercises

import (
	"log"
	"sort"
)

// merge sorts two arrays in O(len(arr1) + len(arr2)) time.
// It does this by popping the smallest head of each array, then adding it to a sorted array.
// It uses O(len(arr1) + len(arr2)) space to hold an extra sorted array.
func merge(arr1, arr2 []int) []int {
	var sorted []int

	var i int
	var j int
	for i < len(arr1) || j < len(arr2) {
		if len(arr2) == 0 {
			sorted = append(sorted, arr1[0])
			arr1 = arr1[1:]
		} else if len(arr1) == 0 {
			sorted = append(sorted, arr2[0])
			arr2 = arr2[1:]
		} else if arr1[0] < arr2[0] {
			sorted = append(sorted, arr1[0])
			arr1 = arr1[1:]
		} else {
			sorted = append(sorted, arr2[0])
			arr2 = arr2[1:]
		}
	}

	return sorted
}

// MergeSort is a recursive sorting algorithm.
// It works by splitting the input array into two arrays.
// It recursively splits the input into two arrays until each array is < 2 length.
// An array of 0 or 1 items is sorted.
// After that, it merges the two sorted lists.
// It runs in O(n log n) time.
// log n because it splits the array into halves.
// n because it iterates over each element in each half for every split.
func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	return merge(MergeSort(arr[:len(arr)/2]), MergeSort(arr[len(arr)/2:]))
}

func main() {
	expected := []int{70, 500, 20, 10, 200, 404, 930, 4, 1, 20000, 294, 484}
	sort.Ints(expected)
	log.Print("Expected: ", expected)
	log.Print("Actual:   ", MergeSort([]int{70, 500, 20, 10, 200, 404, 930, 4, 1, 20000, 294, 484}))
}

// 2009/11/10 23:00:00 Expected: [1 4 10 20 70 200 294 404 484 500 930 20000]
// 2009/11/10 23:00:00 Actual:   [1 4 10 20 70 200 294 404 484 500 930 20000]
