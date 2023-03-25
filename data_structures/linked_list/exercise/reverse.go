// Reverse a singly-linked list
package main

import "log"

type Node struct {
	val  int
	next *Node
}

func Reverse(n *Node) *Node {
	var rev *Node

	for n != nil {
		head := &Node{
			val:  n.val,
			next: rev,
		}

		rev = head
		n = n.next
	}

	return rev
}

func main() {
	list := &Node{
		val: 1,
		next: &Node{
			val: 2,
			next: &Node{
				val: 3,
				next: &Node{
					val: 4,
				},
			},
		},
	}

	rev := Reverse(list)

	var ints []int
	var ints2 []int
	for rev != nil {
		ints = append(ints, rev.val)
		rev = rev.next
		ints2 = append(ints2, list.val)
		list = list.next
	}

	// []int{4, 3, 2, 1}, []int{1, 2, 3, 4}
	log.Printf("%#v, %#v", ints, ints2)
}
