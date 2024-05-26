package heapmethods

func HeapifyUp(arr []int, idx int) {
  if(idx == 0) {
    return
  }
  p := ToParent(idx)
  if(arr[p] > arr[idx]) {
    tmp := arr[p]
    arr[p] = arr[idx]
    arr[idx] = tmp
  }
  HeapifyUp(arr,p)
}
