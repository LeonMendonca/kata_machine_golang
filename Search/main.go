package main

import (
  "fmt"
  searchArr "searchAlgorithm/searchalgorithm"
)

func main() {
  fmt.Println("Search algorithm")
  arr := []int{1,4,6,9,15,17,20,22,27,32,35,37,38,49,52,105,107,171,200,201,202,301,444,555}

  //Binary Search
  resultB := searchArr.BinarySearch(arr,8)
  fmt.Println(resultB)

  //Linear Search
  resultL := searchArr.LinearSearch(arr,52)
  fmt.Println(resultL)
}
