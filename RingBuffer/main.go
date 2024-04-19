package main

import (
  "fmt"
  "ringbuffer/RingBufferMethods"
)

func main() {
    RingBuffer := ringbuffer.RingBuffer(5)
    fmt.Println("RingBuffer")
    RingBuffer.Enqueue(23)
    RingBuffer.Enqueue(55)
    RingBuffer.Enqueue(99)
    RingBuffer.Deque()
    RingBuffer.PrintRingBuffer()
    
}
