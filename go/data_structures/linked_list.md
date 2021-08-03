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

Compared to an array, linked lists are effective at inserting nodes in the middle of the list. 

When you insert a new value into an array, every value after that element has to be moved. 
In a singly-linked list, you only have to update one node's `next` pointer, and add a pointer to the new node.

```go
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
