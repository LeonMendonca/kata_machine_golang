package heapmethods

func HeapifyDown(arr []int, idx, count int) {
  leftIn := ToLeft(idx)
  rightIn := ToRight(idx)

  if(leftIn >= count || rightIn >= count) {
    return
  }


  //swap with left
  if(arr[leftIn] < arr[rightIn] && arr[leftIn] < arr[idx]) {
    tmp := arr[leftIn]
    arr[leftIn] = arr[idx]
    arr[idx] = tmp
    HeapifyDown(arr, leftIn, count)
    //swap with right
  } else if (arr[rightIn] < arr[leftIn] && arr[rightIn] < arr[idx]) {
    tmp := arr[rightIn]
    arr[rightIn] = arr[idx]
    arr[idx] = tmp
    HeapifyDown(arr, rightIn, count)
  }
}
