package heapmethods

import (
  "fmt"
)

func (heap *heap) Insert(value int) {
  heap.heapArray= append(heap.heapArray,value)
  heap.Length++
  HeapifyUp(heap.heapArray,heap.Length-1)
  fmt.Println(heap.heapArray)
}
