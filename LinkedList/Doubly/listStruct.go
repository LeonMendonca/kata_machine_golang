package list

type Node struct {
  value int
  next, prev *Node
}

type LinkedList struct {
  head, tail *Node
  length int
}
