# Linked List

A linked list is like an array but, instead of being referenced by an index, the structure uses pointers to connect nodes. 

The linked list may be singly-linked, meaning it only has pointers in one direction. It may also be doubly-linked, pointing both forward and back in the list.

## Node

Here's the structure of a linked list node:

```go
type SinglyLinkedNode struct {
	val int
	next *SinglyLinkedNode
}
```
Typically, the list is referenced through the _HEAD_ node.

```go
head := &SinglyLinkedNode{}
```

## Access

You cannot access a node in constant time. There is no `list[0]`. Instead, to get the _nth_ node, you must traverse the list. This will be O(n) time complexity.

```go
// Get retrieves the nth element (0 indexed)
func (s *SinglyLinkedNode) Get(n int) *SinglyLinkedNode {
	node := s
  
	for i := 0; i < n; i++ {
		if node == nil {
			return nil
		}

		node = node.next
	}
  
	return node
}
```

## Insertion

Compared to an array, linked lists are effective at inserting nodes in the middle of the list. You can insert a node with an O(1) operation (constant time).

When you insert a new value into an array, every value after that element has to be moved. 
In a singly-linked list, you only have to update one node's `next` pointer, and add a pointer to the new node.

```go
// Insert will insert the given node right after the current node.
func (s *SinglyLinkedNode) Insert(n *SinglyLinkedNode) {
	if s == nil {
		return
	}
  
	old := s.next
	s.next = n
  
	if s.next != nil {
		s.next.next = old
	}
}
```

_Note:_ If you don't already have a pointer to where you want to insert, it will take an O(n) operation to reach that node, then insert.

Similarly, you can insert a node to the top of the list.

```go
head := &SinglyLinkedNode{}
old := head
head = &SinglyLinkedNode{
	next: old,
}
```

## Delete

Deletion is the opposite of insertion; it too can be done with an O(n) operation (assuming you already have a pointer to the node before it).

```go
func (s *SinglyLinkedNode) DeleteNext() error {
	if s == nil {
		return errors.New("cannot delete from nil node")
	}
	if s.next == nil {
		return errors.New("cannot delete nil node")
	}
  
	s.next = s.next.next
	return nil
}
```

If you do not have a pointer to the node before the one you want to delete, you must traverse to it.

```go
// Delete will remove the node at index n and will link to any remaining nodes after index n.
func (s *SinglyLinkedNode) Delete(n int) error {
	if n == 0 {
		return errors.New("cannot delete node 0")
	}
	if s == nil {
		return errors.New("cannot delete node: cannot delete from a nil head")
	}
  
	node := s.Get(n - 1)
	if node == nil {
		return errors.Errorf("cannot delete node: index out of bounds: %d", n)
	}
  
	return node.DeleteNext()
}
```

Deleting the first node is as easy as deleting or replacing the pointer to `head`.

```go
head := &SinglyLinkedList{}
// delete
head = nil
// replace
head = &SinglyLinkedList{}
```
