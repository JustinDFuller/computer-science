/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
*/

package main

import (
	"log"
)

type Stack struct {
	values []rune
}

func (s *Stack) push(c rune) {
	s.values = append(s.values, c)
}

func (s *Stack) pop() rune {
	if len(s.values) == 0 {
		return 0
	}

	last := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return last
}

func (s *Stack) isEmpty() bool {
	return len(s.values) == 0
}

type Blocks struct {
	closers map[rune]rune
}

func (b Blocks) isCloser(r rune) bool {
	_, ok := b.closers[r]
	return ok
}

func (b Blocks) opener(r rune) rune {
	return b.closers[r]
}

var blocks Blocks = Blocks{
	closers: map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	},
}

func isValid(s string) bool {
	var stack Stack

	for _, c := range s {
		if !blocks.isCloser(c) {
			stack.push(c)
			continue
		}

		if last := stack.pop(); last != blocks.opener(c) {
			return false
		}
	}

	return stack.isEmpty()
}

func main() {
	log.Println("", isValid(""), true)
	log.Println("()", isValid("()"), true)
	log.Println("[]", isValid("[]"), true)
	log.Println("{}", isValid("{}"), true)
	log.Println("(}", isValid("(}"), false)
	log.Println("[}", isValid("[}"), false)
	log.Println("{)", isValid("{)"), false)
	log.Println("{", isValid("{)"), false)
}

/*
2009/11/10 23:00:00 "" true true
2009/11/10 23:00:00 () true true
2009/11/10 23:00:00 [] true true
2009/11/10 23:00:00 {} true true
2009/11/10 23:00:00 (} false false
2009/11/10 23:00:00 [} false false
2009/11/10 23:00:00 {) false false
2009/11/10 23:00:00 { false false
*/
