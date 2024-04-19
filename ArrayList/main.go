/*
terms used :
length - count of elements
capacity - length of the array
*/
package main

import(
  "fmt"
  "arraylist/ArrayListMethods"
)


func main() {
    arraylist := arraylist.ArrayList() //creates an array with length 5, and no of elements 0
    //pushing elements
    for i:=1 ; i<=7 ; i++ {
        arraylist.Push(i)
    }
    //fmt.Println(arraylist.pop())
    arraylist.PrintArray()
    fmt.Println("length:",arraylist.Length,"capacity:",arraylist.Capacity)
    //fmt.Println(arraylist.GetIndex(9))
}
