package main

import "log"

type Stack struct {
	vals []rune
}

func (s *Stack) Pop() rune {
	if len(s.vals) == 0 {
		return 0
	}

	r := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return r
}

func (s *Stack) Push(r rune) {
	s.vals = append(s.vals, r)
}

func (s *Stack) Count() int {
	return len(s.vals)
}

func isBalanced(s string) bool {
	var stack Stack

	for _, c := range s {
		if c == '(' {
			stack.Push(c)
			continue
		}

		if last := stack.Pop(); last != '(' {
			return false
		}
	}

	if c := stack.Count(); c != 0 {
		return false
	}

	return true
}

func main() {
	tests := []struct {
		balanced bool
		parens   string
	}{
		{
			balanced: true,
			parens:   "(()((())))()()",
		},
		{
			balanced: false,
			parens:   ")(",
		},
		{
			balanced: false,
			parens:   "(",
		},
	}

	for _, test := range tests {
		if balanced := isBalanced(test.parens); balanced != test.balanced {
			log.Printf("Incorrect result for %s", test.parens)
		}
	}
}
