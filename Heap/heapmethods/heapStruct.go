package heapmethods

type heap struct {
  heapArray []int
  Length int
}

func Heap() *heap {
  return &heap{heapArray:[]int{}, Length:0}
}
