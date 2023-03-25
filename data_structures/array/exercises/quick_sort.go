// QuickSort
package main

import (
	"log"
	"math"
	"sort"
)

var count = 0

// Q: Does partition return the index of the element where every index before it has a lower value?
// [n] denotes the firsthigh which will be the new partition
// begin first full partition
// 	[]int{[70], 500, 20, 10, 200, 404, 930, 4, 1, 20000, 294, 484} start with this array
// 	[]int{[70], 500, 20, 10, 200, 404, 930, 4, 1, 20000, 294, 484} ints[p] = 484, ints[firsthigh] = 70, firsthigh = 0
// 	[]int{70, [500], 20, 10, 200, 404, 930, 4, 1, 20000, 294, 484} i = 0, ints[i] = 70 < ints[p] = 484, but firsthigh = i so nothing swaps, firsthigh = 1
// 	[]int{70, [500], 20, 10, 200, 404, 930, 4, 1, 20000, 294, 484} i = 1, ints[i] = 500 > ints[p] = 484, so nothing swaps, firsthigh = 1
// 	[]int{70, [500], 20, 10, 200, 404, 930, 4, 1, 20000, 294, 484} i = 2, ints[i] = 20 < ints[p] = 484, swap ints[2] and ints[1], firsthigh = 2
// 	[]int{70, 20, [500], 10, 200, 404, 930, 4, 1, 20000, 294, 484} i = 3, ints[i] = 10 < ints[p] = 484, swap ints[3] and ints[2], firsthigh = 3
// 	[]int{70, 20, 10, [500], 200, 404, 930, 4, 1, 20000, 294, 484} i = 4, ints[i] = 200 < ints[p] = 484, swap ints[4] and ints[3], firsthigh = 4
// 	[]int{70, 20, 10, 200, [500], 404, 930, 4, 1, 20000, 294, 484} i = 5, ints[i] = 404 < ints[p] = 484, swap ints[5] and ints[4], firsthigh = 5
// 	[]int{70, 20, 10, 200, 404, [500], 930, 4, 1, 20000, 294, 484} i = 6, ints[i] = 930 > ints[p] = 484, so nothing swaps, firsthigh = 5
// 	[]int{70, 20, 10, 200, 404, [500], 930, 4, 1, 20000, 294, 484} i = 7, ints[i] = 4 < ints[p] = 484, swap ints[7] and ints[5], firsthigh = 6
// 	[]int{70, 20, 10, 200, 404, 4, [930], 500, 1, 20000, 294, 484} i = 8, ints[i] = 1 < ints[p] = 484, swap ints[8] and ints[6], firsthigh = 7
// 	[]int{70, 20, 10, 200, 404, 4, 1, [500], 930, 20000, 294, 484} i = 9, ints[i] = 20000 > ints[p] = 484, so nothing swaps, firsthigh = 7
// 	[]int{70, 20, 10, 200, 404, 4, 1, [500], 930, 20000, 294, 484} i = 10, ints[i] = 294 < ints[p] = 484, swap ints[10] and ints[7], firsthigh = 8
// 	[]int{70, 20, 10, 200, 404, 4, 1, 294, [930], 20000, 500, 484} i = 11, ints[i] = 484 == ints[p] 484, so nothing swaps, firsthigh = 8
// partition returns 8 and the array is shuffled so everything before index 8 < 930 and everything after index 7 >= 484
// begin recursive quicksort between 0 and 7
// 	[]int{[70], 20, 10, 200, 404, 4, 1, 294} i = 0, ints[i] = 70 < ints[p] = 294, but firsthigh = i so nothing swaps, firsthigh = 1
// 	[]int{70, [20], 10, 200, 404, 4, 1, 294} i = 1, ints[i] = 20 < ints[p] = 294, but firsthigh = i so nothing swaps, firsthigh = 2
// 	[]int{70, 20, [10], 200, 404, 4, 1, 294} i = 2, ints[2] = 10 < ints[p] = 294, but firsthigh = i so nothing swaps, firsthigh = 3
// 	[]int{70, 20, 10, [200], 404, 4, 1, 294} i = 3, ints[3] = 200 < ints[p] = 294, but firsthigh = i so nothing swaps, firsthigh = 4
// 	[]int{70, 20, 10, 200, [404], 4, 1, 294} i = 4, ints[4] = 404 > ints[p] = 294, so nothing swaps, firsthigh = 4
// 	[]int{70, 20, 10, 200, [404], 4, 1, 294} i = 5, ints[5] = 4 < ints[p] = 294, swap ints[5] and ints[4], firsthigh = 5
// 	[]int{70, 20, 10, 200, 4, [404], 1, 294} i = 6, ints[6] = 1 < ints[p] = 294, swap ints[6] and ints[5], firsthigh = 6
// 	[]int{70, 20, 10, 200, 4, 1, [404], 294} i = 7, ints[7] = 294 == ints[p] = 294, so nothing swaps, firsthigh = 6
// 	partition returns 6 and the array is shuffled so everything before index 6 < index 7 and everything after index 6 >= 294
// begin recursive quicksort between 0 and 5
func partition(ints []int, low, high int) int {
	// The first time this runs, p/high is the end of the array
	p := high
	// The first time this runs, firsthigh/low is 0
	firsthigh := low

	// The first time this runs, it will run through the entire array
	for i := low; i < high; i++ {
		count++

		// if current index is smaller than the highest index,
		// swap it with the first high index, to ensure that every
		// index before firsthigh is lower than the partition we return
		if ints[i] < ints[p] {
			// because i can be greater than firsthigh
			ints[i], ints[firsthigh] = ints[firsthigh], ints[i]
			firsthigh++
		}
	}
	ints[p], ints[firsthigh] = ints[firsthigh], ints[p]

	return firsthigh
}

func quicksort(i []int, low, high int) {
	if low < high {
		p := partition(i, low, high)
		quicksort(i, low, p-1)
		quicksort(i, p+1, high)
	}
}

func QuickSort(i []int) int {
	count = 0

	quicksort(i, 0, len(i)-1)

	return count
}

func main() {
	expected := []int{70, 500, 20, 10, 200, 404, 930, 4, 1, 20000, 294, 484}
	sort.Ints(expected)
	log.Print("Expected: ", expected)

	actual := []int{70, 500, 20, 10, 200, 404, 930, 4, 1, 20000, 294, 484}
	total := QuickSort(actual)
	log.Print("Recurses: ", total, ", n log n = ", int(float64(len(actual))*math.Log2(float64(len(actual)))))
	log.Print("Actual:   ", actual)

	actual = []int{1, 4, 10, 20, 70, 200, 294, 404, 484, 500, 930, 20000}
	total = QuickSort(actual)
	log.Print("Recurses: ", total, ", n log n = ", int(float64(len(actual))*math.Log2(float64(len(actual)))), ", n^2: ", math.Pow(float64(len(actual)), 2))

	actual = []int{20000, 930, 500, 484, 404, 294, 200, 70, 20, 10, 4, 1}
	total = QuickSort(actual)
	log.Print("Recurses: ", total, ", n log n = ", int(float64(len(actual))*math.Log2(float64(len(actual)))), ", n^2: ", math.Pow(float64(len(actual)), 2))

	actual = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	total = QuickSort(actual)
	log.Print("Recurses: ", total, ", n log n = ", int(float64(len(actual))*math.Log2(float64(len(actual)))), ", n^2: ", math.Pow(float64(len(actual)), 2))
}

/**
2009/11/10 23:00:00 Expected: [1 4 10 20 70 200 294 404 484 500 930 20000]
2009/11/10 23:00:00 Recurses: 32, n log n = 43
2009/11/10 23:00:00 Actual:   [1 4 10 20 70 200 294 404 484 500 930 20000]
2009/11/10 23:00:00 Recurses: 66, n log n = 43, n^2: 144
2009/11/10 23:00:00 Recurses: 66, n log n = 43, n^2: 144
2009/11/10 23:00:00 Recurses: 66, n log n = 43, n^2: 144
**/
