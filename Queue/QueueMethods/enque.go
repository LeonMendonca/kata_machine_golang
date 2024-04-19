package queue

func (q *queue) Enqueue(item int) {
    newNode := &Node{value:item, next:nil}
    if(q.tail == nil) {
        q.head = newNode
        q.tail = newNode
        q.Length++
        return
    }
    
    q.tail.next = newNode
    q.tail = newNode
    q.Length++

}
