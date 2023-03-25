package exercises

import (
	"log"
	"sort"
)

// Take a list of ints as input.
// Design an O(n log n) algorithm that partitions the numbers into pairs,
// where the partition minimizes the maximum sum of a pair.
// For example, say we are given the numbers (1, 3, 5, 9)
// The possible partitions are ((1,3), (5,9)), ((1,5),(3,9)), ((1,9),(3,5)).
// The pair sums for these partitions are (4,14), (6,12), and (10,8)
// Thus, the third partition has 10 as its maximum sum, which is the minimum of the three partitions.
func minMaxSumPair(ints []int) [][]int {
	sort.Ints(ints)

	var res [][]int

	// For odd slices, put the highest int by itself and create an even slice
	if len(ints)%2 != 0 {
		res = append(res, []int{ints[len(ints)-1]})
		ints = ints[:len(ints)-1]
	}

	for i := 0; i < len(ints)/2; i++ {
		res = append(res, []int{ints[i], ints[len(ints)-1-i]})
	}

	return res
}

func main() {
	ints := []int{1, 3, 5, 9, 100}
	log.Print(minMaxSumPair(ints))
}

// 2009/11/10 23:00:00 [[100] [1 9] [3 5]]
