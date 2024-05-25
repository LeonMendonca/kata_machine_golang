package sortingalgorithm

func Partition(arr[]int, lo, hi int) int {
  println("low high",lo,hi)
  var pivot = hi //pivot assigned hi idx, not the value of hi
  var idx = lo-1 //for right sub array, -1 idx doesnt work

  for i:=lo ; i<hi ; i++ {
    if(arr[i] <= arr[pivot]) {
      idx++
      Swap(arr,idx,i)
    }
  }
  idx++
  Swap(arr,idx,pivot)

  return idx
}

func QuickSort(arr[]int, lo, hi int) {
  if(lo > hi) {
    return
  }
  var pivIdx int = Partition(arr,lo,hi)
  QuickSort(arr,lo,pivIdx-1)
  QuickSort(arr,pivIdx+1,hi)
}
