package main

import (
  "fmt"
  queue "queue/queuemethods"
)

func main() {
    Queue := queue.Queue()   
    /*
    for i:=1; i<=5; i++ {
        Queue.Enqueue(i)
    }
    */
    
    fmt.Println("dequed Node",Queue.Deque())
    fmt.Println("head value is",Queue.Peek())
    Queue.PrintQueue()
    fmt.Println("Length is",Queue.Length)
}
