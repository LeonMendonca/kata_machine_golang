package doublylist

import (
  "fmt"
)

func (list *LinkedList) PrintDoublyLinkedListNext() {
  current := list.Head
  for current != nil {
    fmt.Print(current.Value,"->")
    current = current.Next
  }
  println("nil")
}
func (list *LinkedList) PrintDoublyLinkedListPrev() {
  current := list.Tail
  for current != nil {
    fmt.Print(current.Value,"->")
    current = current.Prev
  }
  println("nil")
}
