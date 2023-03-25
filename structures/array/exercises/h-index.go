// Give an efficient algorithm to take the array of citation counts (each count is a non-negative integer)
// of a researcher's papers, and compute the researcher's h-index.
// By definition, a scientist has index h if h of their papers have been cited at least h times,
// while the other n - h papers each have no more than h citations.
package main

import (
	"log"
	"sort"
)

// O(2n + n log n) == O(n log n) -> bottleneck is the sort.
func h(citations []int) int {
	if len(citations) == 0 {
		return 0
	}

	last := len(citations) - 1
	lowest := citations[last]
	count := 1

	// we need the total count to know when to stop
	var total int
	for _, c := range citations {
		total += c
	}

	sort.Ints(citations)

	for c := last - 1; c > 0; c-- {
		count++
		lowest = citations[c]
		total -= lowest
		if total < lowest {
			break
		}
	}

	if lowest < count {
		return lowest
	}

	return count
}

func main() {
	// h index is the minimum of: papers written, lowest citation count so far, maximum remaining citation
	// seems like it could be beneficial to walk through in reverse sorted order (highest to lowest)
	// where h is min(index + 1, citation count)
	// we stop when the next citation count <= h
	log.Print("h is limited by papers:        ", h([]int{2}), " expected: ", 1)
	log.Print("h is limited by citations:     ", h([]int{2, 2, 2}), " expected: ", 2)
	log.Print("remainder is less than lowest: ", h([]int{2, 4, 4, 5, 6, 7}), " expected: ", 4)
	log.Print(h([]int{}))
	log.Print(h(nil))
}

// 2009/11/10 23:00:00 h is limited by papers:        1 expected: 1
// 2009/11/10 23:00:00 h is limited by citations:     2 expected: 2
// 2009/11/10 23:00:00 remainder is less than lowest: 4 expected: 4
// 2009/11/10 23:00:00 0
// 2009/11/10 23:00:00 0
