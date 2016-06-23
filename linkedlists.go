package main

import "fmt"

func main() {
	list := NewList()
	list.appendToTail(5)
	list.appendToTail(6)
	list.appendToTail(4)
	list.appendToTail(2)
	list.appendToTail(7)
	list.appendToTail(8)
	list.appendToTail(3)

	// fmt.Println(list)
	fmt.Println(list.partitionList(5))
	// list.removeDups()
	//
	// for i := 2; i < 50000; i++ {
	// list.appendToTail(i)
	// }
	// list.deleteNode(6)
	// fmt.Printf("%+v", list)
	// fmt.Printf("\n\n%+v%d%d", a, b, c)
	// fmt.Printf("\n\n%d\n", <-a)
	//
	// result := testing.Benchmark(func(*testing.B) {
	// 	_ = list.kthToTheLast(2)
	// })
	//
	// fmt.Printf("\niterations %+v\ntime %+v\nBytes %+v\nMemAlocs %+v\nMemBytes %+v\n\n", result.N, result.T, result.Bytes, result.MemAllocs, result.MemBytes)
	//
	// result = testing.Benchmark(func(*testing.B) {
	// 	_, _ = list.kthToTheLastRecursive(list.Head, 0, 2)
	// })
	//
	// fmt.Printf("\niterations %+v\ntime %+v\nBytes %+v\nMemAlocs %+v\nMemBytes %+v\n\n", result.N, result.T, result.Bytes, result.MemAllocs, result.MemBytes)
	//
	// result = testing.Benchmark(func(*testing.B) {
	// 	_, _ = list.kthToTheLastChan(list.Head, 0, 2)
	// })
	//
	// fmt.Printf("\niterations %+v\ntime %+v\nBytes %+v\nMemAlocs %+v\nMemBytes %+v\n\n", result.N, result.T, result.Bytes, result.MemAllocs, result.MemBytes)
	// //

}

//List d
type List struct {
	Head   *node
	Tail   *node
	Length int
}
type node struct {
	Next *node
	Data int
}

//NewList d
func NewList() *List {
	return &List{
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
}

//NewNode sd
func newNode(d int) *node {
	return &node{
		Next: nil,
		Data: d,
	}
}

func (l *List) appendToTail(d int) {
	lastNode := newNode(d)
	if l.Head == nil {
		l.Head = lastNode
		l.Tail = lastNode
		l.Length++
		return
	}

	head := l.Head

	for head.Next != nil {
		head = head.Next
	}

	head.Next = lastNode
	l.Tail = lastNode
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

func (l *List) kthToTheLast(k int) *node {
	n := l.Head
	var counter int
	for n != nil {
		if counter == l.Length-k {
			return n
		}
		counter++
		n = n.Next
	}
	return nil
}

func (l *List) kthToTheLastRecursive(n *node, count, k int) (*node, int) {
	var bubble *node
	if n.Next == nil {
		return n, count - k
	}

	bubble, k = l.kthToTheLastRecursive(n.Next, count+1, k)
	if count == k {
		return bubble, k - 1
	}

	return n, k
}

func (l *List) kthToTheLastChan(n *node, count, k int) (chan *node, int) {
	if n == nil {
		return nil, count - k
	}

	ch, k := l.kthToTheLastChan(n.Next, count+1, k)
	if count == k {
		ch := make(chan *node, 1)
		ch <- n
		close(ch)
		return ch, k
	}

	return ch, k
}

func deleteNode(n *node) bool {
	if n.Next != nil {
		n.Data = n.Next.Data
		n.Next = n.Next.Next
		return true
	}
	return false
}

func (l *List) partitionList(x int) *List {
	head := NewList()
	tail := NewList()
	var higher *node
	var end *node

	n := l.Head
	for n != nil {
		if n.Data < x {
			if head.Head == nil {
				head.Head = &node{Data: n.Data}
				higher = head.Head
			} else {
				higher.Next = n
				higher = higher.Next
				higher.Data = n.Data
			}
		}
		if n.Data >= x {
			if tail.Head == nil {
				tail.Head = &node{Data: n.Data}
				end = tail.Head
			} else {
				end.Next = n
				end = end.Next
				end.Data = n.Data
			}
		}
		n = n.Next
	}

	higher.Next = tail.Head
	end.Next = nil

	result := &List{
		Length: l.Length,
		Head:   head.Head,
		Tail:   end,
	}

	return result
}
