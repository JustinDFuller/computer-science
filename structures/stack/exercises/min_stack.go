// Design a stack S that supports S.push(x), S.pop(), and S.findmin(), which returns the minimum element of S.
// All operations should run in constant time.
package main

import (
	"log"
	"sort"
)

type Stack struct {
	ints []int
}

// O(n log n) because of the sort
func (s *Stack) Push(i int) {
	if s.ints == nil {
		s.ints = []int{i}
		return
	}
	// Append is O(n) because of amortization
	unsorted := append(s.ints, i)
	// Go uses quickSort which is O(n log n)
	sort.Ints(unsorted)
	s.ints = unsorted
}

// O(1) because we don't have to reallocate the underlying array
func (s *Stack) Pop() int {
	if s.ints == nil {
		return 0
	}
	if len(s.ints) == 0 {
		return 0
	}
	i := s.ints[0]
	s.ints = s.ints[1:]
	return i
}

// O(1) because we just get the first element
func (s *Stack) Min() int {
	if s.ints == nil {
		return 0
	}
	if len(s.ints) == 0 {
		return 0
	}
	return s.ints[0]
}

func main() {
	var s Stack

	s.Push(1)
	s.Push(3)

	// 1
	log.Print(s.Min())
	// 1
	log.Print(s.Pop())

	// 3
	log.Print(s.Min())
	// 3
	log.Print(s.Pop())

	// 0
	log.Print(s.Pop())
	// 0
	log.Print(s.Min())

	s.Push(10)
	s.Push(100)

	// 10
	log.Print(s.Min())
	// 10
	log.Print(s.Pop())
}
