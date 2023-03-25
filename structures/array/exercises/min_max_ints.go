package main

import (
	"log"
	"math"
	"sort"
)

/**
Let S be an unsorted array of n integers.
Given an algorithm that finds the pair x,y ∈ S that maximizes |x-y|.
Your algorithm must run in O(n) worst-case time.
**/
func findUnsortedMaxDifference(s []int) (int, int) {
	if s == nil || len(s) == 0 {
		return 0, 0
	}

	max := s[0]
	min := s[0]

	for _, i := range s {
		if i > max {
			max = i
		}

		if i < min {
			min = i
		}
	}

	return max, min
}

/**
Let S be a sorted array of n integers.
Give an algorithm that finds the pair x,y ∈ S that maximizes |x-y|.
Your algorithm must run in O(1) worst-case time.
**/
func findSortedMaxDifference(s []int) (int, int) {
	return s[len(s)-1], s[0]
}

/**
Let S be an unsorted array of n integers.
Give an algorithm that finds the pair x, y ∈ S that minimizes |x-y|, for x != y.
Your algorithm must run in O(n log n) worst-case time.
**/
func findUnsortedMinDifference(s []int) (int, int) {
	sort.Ints(s)
	return findSortedMinDifference(s)
}

/**
Let S be a sorted array of n integers.
Give an algorithm that finds the pair x, y ∈ S that minimizes |x-y|, for x != y.
Your algorithm must run in O(n) worst-case time.
**/
func findSortedMinDifference(s []int) (int, int) {
	if len(s) < 1 {
		return 0, 0
	}

	x := s[0]
	y := s[1]
	diff := math.Abs(float64(y - x))

	for index, val := range s[1:] {
		if index+2 == len(s) {
			break
		}

		if d := math.Abs(float64(s[index+2] - val)); d < diff {
			x = val
			y = s[index+2]
			diff = d
		}
	}

	return y, x
}

func main() {
	s := []int{50, 99, 2, 100, 1, 30}
	max, min := findUnsortedMaxDifference(s)
	log.Printf("Unsorted max difference: max %d, min: %d", max, min)

	sort.Ints(s)
	max, min = findSortedMaxDifference(s)
	log.Printf("Unsorted max difference: max %d, min: %d", max, min)

	s = []int{50, 99, 3, 100, 1, 30}
	max, min = findUnsortedMinDifference(s)
	log.Printf("Unsorted min difference: max %d, min: %d", max, min)

	sort.Ints(s)
	max, min = findSortedMinDifference(s)
	log.Printf("Unsorted min difference: max %d, min: %d", max, min)
}

/*
2009/11/10 23:00:00 Unsorted max difference: max 100, min: 1
2009/11/10 23:00:00 Unsorted max difference: max 100, min: 1
2009/11/10 23:00:00 Unsorted min difference: max 100, min: 99
2009/11/10 23:00:00 Unsorted min difference: max 100, min: 99
*/
