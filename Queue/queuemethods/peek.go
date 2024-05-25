package queuemethods

func (q *queue) Peek() int {
    if(q.head != nil) {
        return q.head.value 
    }
    return 0
}
