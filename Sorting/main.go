package main

import (
  "fmt"
  sortArray "sortingAlgorithm/sortingalgorithm"
)

func main() {
  fmt.Println("Sorting algorithm")
  arr := []int{8,5,1,2,7}
  //sortArray.BubbleSort(arr) 
  sortArray.SelectionSort(arr)
  fmt.Println(arr)
}
