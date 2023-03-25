// The mode is the number that occurs most frequently.
// The set {4,6,2,4,3,1} has a mode of 4.
// Given an efficient and correct algorithm to compute the mode of a bag of n numbers.
package main

import "log"

func mode(ints []int) int {
	counts := map[int]int{}

	for _, i := range ints {
		counts[i] += 1
	}

	max := counts[ints[0]]
	maxVal := ints[0]
	for key, val := range counts {
		if val > max {
			max = val
			maxVal = key
		}
	}

	return maxVal
}

func main() {
	ints := []int{4, 6, 2, 4, 3, 1}
	log.Print(mode(ints))
}


// 2009/11/10 23:00:00 4
