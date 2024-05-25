package arraylist

import (
  "fmt"
)

func (arrlt *arrayList) PrintArray() {
    for i:=0 ; i<arrlt.Length ; i++ {
        fmt.Print(arrlt.array[i]," ")
    }
    println()
}


