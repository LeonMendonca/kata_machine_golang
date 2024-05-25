package queuemethods

func (q *queue) Deque() int {
    if(q.head == nil) {
        println("empty queue")
        return -1
    }

    if(q.head.next == nil) {
        q.tail = nil
    }

    head := q.head
    q.head = q.head.next
    q.Length--

    head.next = nil
    return head.value
}
