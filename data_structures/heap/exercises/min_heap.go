package main

import (
	"errors"
	"log"
)

type Heap struct {
	values []int
}

func (h *Heap) Parent(index int) (int, error) {
	l := len(h.values)
	if l == 0 {
		return 0, errors.New("no parent: empty heap")
	}

	if index < 1 {
		return 0, errors.New("no parent: index too small")
	}

	loc := (index - 1) / 2
	if loc > l-1 {
		return 0, errors.New("no parent: index too high")
	}

	if loc < 0 {
		return 0, errors.New("no parent: index too small")
	}

	return h.values[loc], nil
}

func (h *Heap) Left(index int) (int, error) {
	l := len(h.values)
	if l == 0 {
		return 0, errors.New("no parent: empty heap")
	}

	loc := (index * 2) + 1
	if loc > l-1 {
		return 0, errors.New("no child: index too high")
	}

	return h.values[loc], nil
}

func (h *Heap) Right(index int) (int, error) {
	l := len(h.values)
	if l == 0 {
		return 0, errors.New("no parent: empty heap")
	}

	loc := (index * 2) + 2
	if loc > l-1 {
		return 0, errors.New("no child: index too high")
	}

	return h.values[loc], nil
}

func (h *Heap) swap(parent, child int) {
	p := h.values[parent]
	c := h.values[child]
	h.values[parent] = c
	h.values[child] = p
}

func (h *Heap) Insert(val int) {
	h.values = append(h.values, val)

	index := len(h.values) - 1
	parent, err := h.Parent(index)
	for err == nil && parent > val {
		parentIndex := (index - 1) / 2
		h.swap(parentIndex, index)
		index = parentIndex
		parent, err = h.Parent(index)
	}
}

func (h *Heap) Pop() (int, error) {
	l := len(h.values)
	if l == 0 {
		return 0, errors.New("empty heap")
	}

	head := h.values[0]
	if l == 1 {
		h.values = nil
		return head, nil
	}

	val := h.values[l-1]      // grab the last value
	index := 0                // start the index at the head
	h.values[0] = val         // move the last value to the head
	h.values = h.values[:l-1] // remove the last value from the heap

	for {
		min := val
		var childIndex int

		if child, err := h.Left(index); err == nil && child < min {
			min = child
			childIndex = (index * 2) + 1
		}

		if child, err := h.Right(index); err == nil && child < min {
			min = child
			childIndex = (index * 2) + 2

		}

		if min != val && childIndex != index {
			h.swap(index, childIndex)
			index = childIndex
		} else {
			break
		}
	}

	return head, nil
}

func main() {
	h := Heap{
		values: []int{1492, 1783, 1776, 1804, 1865, 1945, 1963, 1918, 2001, 1941},
	}

	h.Insert(2020)
	h.Insert(1400)

	v, err := h.Pop()
	for err == nil {
		log.Print(v)
		v, err = h.Pop()
	}
	log.Print(err)
}

/*
2009/11/10 23:00:00 1400
2009/11/10 23:00:00 1492
2009/11/10 23:00:00 1776
2009/11/10 23:00:00 1783
2009/11/10 23:00:00 1804
2009/11/10 23:00:00 1865
2009/11/10 23:00:00 1918
2009/11/10 23:00:00 1941
2009/11/10 23:00:00 1945
2009/11/10 23:00:00 1963
2009/11/10 23:00:00 2001
2009/11/10 23:00:00 2020
2009/11/10 23:00:00 empty heap
*/
