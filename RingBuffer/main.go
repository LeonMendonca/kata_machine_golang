package main

import (
  "fmt"
  "ringbuffer/RingBufferMethods"
)

func main() {
    RingBuffer := ringbuffer.RingBuffer(5)
    fmt.Println("RingBuffer")
    /*
    RingBuffer.Enqueue(23)
    RingBuffer.Deque()
    */
    RingBuffer.PrintRingBuffer()
}
