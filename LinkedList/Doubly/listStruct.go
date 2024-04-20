package doublylist


type Node struct {
  Value int
  Next, Prev *Node
}

type LinkedList struct {
  Head, Tail *Node
  Length int
}

//constructor
func DoublyLinkedList() *LinkedList {
  return &LinkedList {
    Head:nil,
    Tail:nil,
    Length:0,
  }
}
