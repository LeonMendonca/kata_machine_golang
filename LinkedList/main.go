package main

import (
  "fmt"
  "linkedlist/Singly"
)

func main() {
  fmt.Print("linked list\n\n")
  Linkedlist := list.LinkedList()

/*
  Linkedlist.Prepend(34)
  Linkedlist.Prepend(45)
  Linkedlist.Prepend(56)
*/
/*
  Linkedlist.InsertAt(99,5)
  Linkedlist.InsertAt(100,2)
*/
/*
  Linkedlist.Append(56)
*/

  Linkedlist.Prepend(34)
  Linkedlist.Prepend(45)
  Linkedlist.Prepend(56)
/*
  fmt.Println("Popped Node:",Linkedlist.Pop())
  fmt.Println("Popped Node:",Linkedlist.Deque())
  fmt.Println("Popped Node:",Linkedlist.RemoveAt(3))
*/

  Linkedlist.PrintList()
  fmt.Println("length is",Linkedlist.Length)
}
