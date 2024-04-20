package queue


type Node struct {
    value int
    next *Node
}

type queue struct {
    head, tail *Node
    Length int
}

//constructor
func Queue() (*queue) {
    return &queue{
      head:nil, 
      tail:nil, 
      Length:0,
    }
}
