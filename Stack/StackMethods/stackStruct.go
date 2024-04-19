package stack

import (
  "fmt"
)

type Node struct {
    value int
    prev *Node
}

type stack struct {
    head *Node
    Length int
}

func Stack() *stack {
    return &stack{
      head:nil,
      Length:0,
    } 
}

func (s *stack) Printstack() {
    curr := s.head
    for curr != nil {
        fmt.Println("|_",curr.value,"_|")
        curr = curr.prev
    }
}
