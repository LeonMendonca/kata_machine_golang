package singlylist

import (
  "fmt"
)

func (list *linkedlist) PrintList() {
  current := list.head
  for current != nil {
    fmt.Print(current.value,"->")
    current = current.next
  }
  println("nil")
}
