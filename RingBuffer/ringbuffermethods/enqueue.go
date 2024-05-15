package ringbuffermethods

/*
What if you increment the tail instead of check?

check - after series of increment, there is a chance that tail may point to the same index as head
And since if head == tail, then the list is said to be empty (i.e. initial point).
The PrintRingBuffer function does not print the Items
*/

func (rb *ringBuffer) Enqueue (item int) {
    //rb.tail = (rb.tail+1)%rb.bufferSize - Bug!
    check := (rb.tail+1)%rb.bufferSize
    if(rb.head == check) {
        //may create a large array buffer later, instead of this 
        println("queue is full cant insert",item)
        return
    }
    rb.tail = (rb.tail+1)%rb.bufferSize
    rb.array[rb.tail] = item
}
