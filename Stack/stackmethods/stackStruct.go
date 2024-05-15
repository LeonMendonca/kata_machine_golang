package stackmethods

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


