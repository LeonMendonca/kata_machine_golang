package ringbuffer

func (rb *ringBuffer) Enqueue (item int) {
    check := (rb.tail+1)%rb.bufferSize
    if(rb.head == check) {
        //may create a large array buffer later, instead of this 
        println("queue is full cant insert",item)
        return
    }
    rb.tail = (rb.tail+1)%rb.bufferSize
    rb.array[rb.tail] = item
}
