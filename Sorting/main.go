package main

import (
  "fmt"
  sortArray "sortingAlgorithm/sortingalgorithm"
)

func main() {
  fmt.Println("Sorting algorithm")
  //arr := []int{6,5,1,2,7}
  //arr := []int{24,9,29,14,19,27}
  arr := []int{8,7,6,4,5}

  //sortArray.BubbleSort(arr) 
  //sortArray.SelectionSort(arr)
  sortArray.QuickSort(arr,0,len(arr)-1)
  //println(len(arr)-1)
  fmt.Println(arr)
}
