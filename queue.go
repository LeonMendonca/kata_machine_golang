package main

import ("fmt")

type node struct {
    value int
    next *node
}

type queue struct {
    head, tail *node
    length int
}

func queueConstructor() (*queue) {
    return &queue{head:nil, tail:nil, length:0}
}

func (q *queue) PrintQueue() {
    curr := q.head
    for curr != nil {
        fmt.Print(curr.value,"->")
        curr = curr.next
    }
    println("nil")
}

func (q *queue) enqueue(item int) {
    newNode := &node{value:item, next:nil}
    if(q.tail == nil) {
        q.head = newNode
        q.tail = newNode
        q.length++
        return
    }
    
    q.tail.next = newNode
    q.tail = newNode
    q.length++

}

func (q *queue) deque() int {
    if(q.head == nil) {
        println("empty queue")
        return 0
    }

    head := q.head

    if(q.head.next == nil) {
        q.tail = nil
    }
    q.head = q.head.next
    q.length--

    head.next = nil
    return head.value
}

func (q *queue) peek() int {
    if(q.head != nil) {
        return q.head.value 
    }
    return 0
}

func main() {
    Queue := queueConstructor()   
    /*
    for i:=1; i<=5; i++ {
        Queue.enqueue(i)
    }
    */
    Queue.enqueue(1)
    fmt.Println("dequed node",Queue.deque())
    fmt.Println("head value is",Queue.peek())
    Queue.PrintQueue()
    fmt.Println("length is",Queue.length)
    fmt.Println("bug",Queue.tail)
}
