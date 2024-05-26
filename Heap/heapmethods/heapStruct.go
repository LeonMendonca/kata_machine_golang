package heapmethods

type heap struct {
  heapArray []int
  Length int
}

func Heap() *heap {
  return &heap{heapArray:[]int{50,100,200,101,103,202,204}, Length:7}
}
