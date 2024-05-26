package main

import (
  "fmt"
  heap "heap/heapmethods"
)

func main() {
  Heap := heap.Heap()
  fmt.Println(Heap.Delete())
  fmt.Println(Heap.Delete())
  fmt.Println(Heap.Delete())
  fmt.Println(Heap.Delete())
  fmt.Println(Heap.Delete())
  fmt.Println(Heap.Delete())
  fmt.Println(Heap.Delete())
  fmt.Println(Heap.Delete())
  fmt.Println(Heap)
  //Heap.Insert(8)
}
