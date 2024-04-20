package ringbuffer

//Since, the head is always has 0 value, because it indicates dead end for the insertion performed through tail
//we print the ringBuffer from head+1
func (rb ringBuffer) PrintRingBuffer () {
    /*
    if ringBuffer is empty and this function is called
    avoid printing series of 0
    */
    if(rb.head == rb.tail) {
        return
    }

    i := rb.head+1
    for {
        print(rb.array[i]," ")
        if(i == rb.tail) {
            break
        }
        i = (i+1)%rb.bufferSize 
    }
    println()
}


