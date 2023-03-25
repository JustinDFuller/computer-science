// Binary Search
package exercises

import (
	"log"
)

func binarySearchLoop(ints []int, val int) int {
	low := 0
	high := len(ints) - 1

	for {
		middle := (low + high) / 2

		if low > high {
			return -1
		}

		if ints[middle] == val {
			return middle
		}

		if ints[middle] > val {
			high = middle - 1
		} else {
			low = middle + 1
		}

	}
}

func binarySearchRecurseInner(ints []int, val, low, high int) int {
	middle := (low + high) / 2

	if low > high {
		return -1
	}

	if ints[middle] == val {
		return middle
	}

	if ints[middle] > val {
		return binarySearchRecurseInner(ints, val, low, middle-1)
	}

	return binarySearchRecurseInner(ints, val, middle+1, high)
}

func binarySearchRecurse(ints []int, val int) int {
	return binarySearchRecurseInner(ints, val, 0, len(ints)-1)
}

func main() {
	ints := []int{1, 3, 4, 12, 52, 98, 104, 203, 400, 500, 700}

	for index, val := range ints {
		log.Printf("Loop Found: %d, wanted: %d", binarySearchLoop(ints, val), index)
		log.Printf("Recurse Found: %d, wanted: %d", binarySearchRecurse(ints, val), index)
	}

	log.Printf("Loop Found: %d, wanted: %d", binarySearchLoop(ints, 0), -1)
	log.Printf("Recurse Found: %d, wanted: %d", binarySearchRecurse(ints, 900), -1)
}

/*
2009/11/10 23:00:00 Loop Found: 0, wanted: 0
2009/11/10 23:00:00 Recurse Found: 0, wanted: 0
2009/11/10 23:00:00 Loop Found: 1, wanted: 1
2009/11/10 23:00:00 Recurse Found: 1, wanted: 1
2009/11/10 23:00:00 Loop Found: 2, wanted: 2
2009/11/10 23:00:00 Recurse Found: 2, wanted: 2
2009/11/10 23:00:00 Loop Found: 3, wanted: 3
2009/11/10 23:00:00 Recurse Found: 3, wanted: 3
2009/11/10 23:00:00 Loop Found: 4, wanted: 4
2009/11/10 23:00:00 Recurse Found: 4, wanted: 4
2009/11/10 23:00:00 Loop Found: 5, wanted: 5
2009/11/10 23:00:00 Recurse Found: 5, wanted: 5
2009/11/10 23:00:00 Loop Found: 6, wanted: 6
2009/11/10 23:00:00 Recurse Found: 6, wanted: 6
2009/11/10 23:00:00 Loop Found: 7, wanted: 7
2009/11/10 23:00:00 Recurse Found: 7, wanted: 7
2009/11/10 23:00:00 Loop Found: 8, wanted: 8
2009/11/10 23:00:00 Recurse Found: 8, wanted: 8
2009/11/10 23:00:00 Loop Found: 9, wanted: 9
2009/11/10 23:00:00 Recurse Found: 9, wanted: 9
2009/11/10 23:00:00 Loop Found: 10, wanted: 10
2009/11/10 23:00:00 Recurse Found: 10, wanted: 10
2009/11/10 23:00:00 Loop Found: -1, wanted: -1
2009/11/10 23:00:00 Recurse Found: -1, wanted: -1
*/
