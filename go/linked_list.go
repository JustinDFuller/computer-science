type LinkedList struct {
    val int
    next *LinkedList
}


/** Initialize your data structure here. */
func Constructor() LinkedList {
    return LinkedList{}
}


func (l *LinkedList) GetNode(index int) *LinkedList {
    node := l

    for i := 0; i < index; i++ {
        if node == nil {
            return nil
        }
        node = node.next
    }
    
    return node
}
/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (l *LinkedList) Get(index int) int {
    node := l.GetNode(index) 
    if node == nil {
        return -1
    }

    return node.val
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (l *LinkedList) AddAtHead(val int)  {
    old := l
    l = &LinkedList{
        val: val,
        next: old,
    }
}


/** Append a node of value val to the last element of the linked list. */
func (l *LinkedList) AddAtTail(val int)  {
    node := l
    if node == nil {
        l.AddAtHead(val)
        return
    }
    
    for node.next != nil {
        node = node.next
    }
    
    node.next = &LinkedList{
        val: val,
    }
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (l *LinkedList) AddAtIndex(index int, val int)  {
    node := l.GetNode(index - 1)
    if node == nil {
        return
    }
    
    if node.next == nil {
        node.next = &LinkedList{
            val: val,
        } 
    } else {
        node.next = &LinkedList{
            val: val,
            next: node.next.next,
        }
    }
    
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (l *LinkedList) DeleteAtIndex(index int)  {
    node := l.GetNode(index - 1)
    if node == nil {
        return
    }    
    if node.next == nil {
        return
    }
    node.next = node.next.next
}
