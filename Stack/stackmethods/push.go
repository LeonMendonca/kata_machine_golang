package stackmethods

func (s *stack) Push(item int) {
    newNode := &Node{value:item, prev:nil}
    if(s.head == nil) {
        s.Length++
        s.head = newNode
        return
    }

    newNode.prev = s.head
    s.head = newNode
    s.Length++
}
