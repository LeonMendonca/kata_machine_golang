//insertion and deletion are in 2 different files, these are required by both files
package singlylist


type Node struct {
  value int
  next *Node
}

type linkedlist struct {
  head, tail *Node
  Length int
}

//constructor
func SinglyLinkedList() *linkedlist {
  return &linkedlist{
    head:nil,
    tail:nil,
    Length:0,
  }
}


