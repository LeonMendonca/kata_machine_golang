//mutable array in GO, so no pointers required
package main

import (
  "fmt"
  array "arrays/arraymethods"
)

func main() {
    arr := []int{1,2,3,4}
    //Get Index
    //value,result := array.GetIndex(arr,2)
    //fmt.Println(value,result)

    //Update array
    result := array.UpdateElement(arr,3,77)
    fmt.Println(result)

    //Delete array element
    //value,result := array.DeleteElement(arr,5)
    //fmt.Println(value,result)

    fmt.Println(arr)
}

