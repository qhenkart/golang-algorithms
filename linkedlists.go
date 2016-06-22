package main

import "fmt"

func main() {
	list := NewList(1)
	list.appendToTail(5)
	list.appendToTail(5)
	list.appendToTail(3)
	list.appendToTail(5)
	list.appendToTail(5)
	list.removeDups()
	// list.deleteNode(6)
	// fmt.Printf("%+v")
	fmt.Printf("%+v", list)
}

type List struct {
	Head   *node
	Tail   *node
	Length int
}
type node struct {
	Next *node
	Data int
}

func NewList(d int) *List {
	n := &node{
		Next: nil,
		Data: d,
	}
	return &List{
		Head:   n,
		Tail:   n,
		Length: 1,
	}
}

func NewNode(d int) *node {
	return &node{
		Next: nil,
		Data: d,
	}
}

func (l *List) appendToTail(d int) {
	lastNode := NewNode(d)
	head := l.Head

	for head.Next != nil {
		head = head.Next
	}

	head.Next = lastNode
	l.Length++

}

func (l *List) deleteNode(d int) bool {
	n := l.Head

	if n.Data == d {
		l.Head = n.Next
		return true
	}

	var prev node
	for n.Next != nil || n.Data == d {
		if n.Data == d {
			prev.Next = n.Next
			l.Length--
			return true
		}

		prev = *n
		n = n.Next
	}
	return false
}

func (l *List) removeDups() {
	hash := make(map[int]int)

	n := l.Head
	prev := l.Head
	for n != nil {
		if _, ok := hash[n.Data]; ok {
			prev.Next = n.Next
			l.Length--
		} else {
			hash[n.Data]++
		}
		n = n.Next
	}

}
