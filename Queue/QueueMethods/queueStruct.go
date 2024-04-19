package queue

import ("fmt")

type Node struct {
    value int
    next *Node
}

type queue struct {
    head, tail *Node
    Length int
}

//constructor
func Queue() (*queue) {
    return &queue{
      head:nil, 
      tail:nil, 
      Length:0,
    }
}

func (q *queue) PrintQueue() {
    curr := q.head
    for curr != nil {
        fmt.Print(curr.value,"->")
        curr = curr.next
    }
    println("nil")
}
