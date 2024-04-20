package stack

import (
  "fmt"
)

func (s *stack) Printstack() {
    curr := s.head
    for curr != nil {
        fmt.Println("|_",curr.value,"_|")
        curr = curr.prev
    }
}
