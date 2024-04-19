package ringbuffer

func (rb *ringBuffer) Deque () int {
    if(rb.head == rb.tail) {
        println("queue is empty")
        return -1
    }

    rb.head = (rb.head+1)%rb.bufferSize
    var dequedItem = rb.array[rb.head]
    rb.array[rb.head] = 0
    return dequedItem
}
