package heapmethods

func (heap *heap) Delete() int {
  if(heap.Length == 0) {
    return -1
  }
  //save the minheap
  minHeap := heap.heapArray[0]
  lastIdxEl := heap.heapArray[heap.Length-1]
  heap.heapArray[0] = lastIdxEl
  heap.Length--
  heap.heapArray = heap.heapArray[:heap.Length]
  HeapifyDown(heap.heapArray, 0, heap.Length)
  return minHeap
}
