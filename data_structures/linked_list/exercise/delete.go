// delete from a singly-linked list in constant time
package main

import (
	"fmt"
	"log"
)

type Node struct {
	val  int
	next *Node
}

func (n *Node) Push(n2 *Node) {

}

func (n *Node) Delete(n2 **Node) {
	if n2 == nil {
		log.Fatal("oops")
		return
	}

	n3 := *n2
	if n3 == nil || n3.next == nil {
		// this doesn't work
		*n2 = nil
		return
	}

	n3.val = n3.next.val
	n3.next = n3.next.next
}

func (n *Node) Ints() []int {
	var out []int

	curr := n
	for curr != nil {
		out = append(out, curr.val)
		curr = curr.next
	}

	return out
}

func main() {

	b := Node{
		val: 3,
	}

	n2 := &b

	a := Node{
		val:  2,
		next: n2,
	}

	n1 := &a

	list := Node{
		val:  1,
		next: n1,
	}

	fmt.Println(list.Ints())

	list.Delete(&n1)

	fmt.Println(list.Ints())

	list.Delete(&n2)

	fmt.Println(list.Ints())

}
