//insertion and deletion are in 2 different files, these are required by both files
package list

import ("fmt")

type Node struct {
  value int
  next *Node
}

type linkedlist struct {
  head, tail *Node
  Length int
}

func LinkedList() *linkedlist {
  return &linkedlist{
    head:nil,
    tail:nil,
    Length:0,
  }
}

func (list *linkedlist) PrintList() {
  current := list.head
  for current != nil {
    fmt.Print(current.value,"->")
    current = current.next
  }
  println("nil")
}
