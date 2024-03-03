package main

import "fmt"

type ringBuffer struct {
    array []int
    head, tail ,bufferSize int
}

/*
NOTE : You can use len(aray) instead of rb.bufferSize, but since we already 
assign the size to the array in the constructor its better to use that
*/

//Since, the head is always 0, because it indicates dead end for the insertion performed through tail
//we print the ringBuffer from head+1
func (rb ringBuffer) printRingBuffer () {
    //the if statement avoids unnecessary iteration
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

func constructorRingBuffer(bufferSize int) *ringBuffer {
    return &ringBuffer{make([]int,bufferSize), 0, 0, bufferSize}
}

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

func (rb *ringBuffer) Deque () int {
    if(rb.head == rb.tail) {
        println("queue is empty")
        return -1
    }

    rb.head = (rb.head+1)%rb.bufferSize
    var dequedItem int = rb.array[rb.head]
    rb.array[rb.head] = 0
    return dequedItem
}

func main() {
    RingBuffer := constructorRingBuffer(5)
    fmt.Println("RingBuffer")
    RingBuffer.Enqueue(23)
    RingBuffer.Enqueue(55)
    RingBuffer.Enqueue(99)
    RingBuffer.Deque()
    RingBuffer.printRingBuffer()
    
}
