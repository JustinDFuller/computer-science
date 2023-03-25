// Given two sets S1 and S2 and a number x, describe an algorithm
// For finding whether there exists a pair of elements, one from S1 and another from S2
// that add up to x.
package main

import "log"

func pairExists(x int, s1, s2 []int) bool {
	vals := map[int]bool{}
	for _, i := range s1 {
		vals[i] = true
	}

	for _, i := range s2 {
		if vals[x-i] {
			return true
		}
	}

	return false
}

func main() {
	s1 := []int{4, 6, 2, 4, 3, 1}
	s2 := []int{3, 9, 1, 5, 2, 0}
	log.Print(pairExists(11, s1, s2))
	log.Print(pairExists(17, s1, s2))
}


// 2009/11/10 23:00:00 true
// 2009/11/10 23:00:00 false
