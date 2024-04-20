package queue

import ("fmt")

func (q *queue) PrintQueue() {
    curr := q.head
    for curr != nil {
        fmt.Print(curr.value,"->")
        curr = curr.next
    }
    println("nil")
}
