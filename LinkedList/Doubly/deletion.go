package doublylist
func (list *LinkedList) Pop() int {

  //no node
  if(list.Head == nil && list.Tail == nil) {
    println("empty list")
    return -1
  }
  
  //single node
  if(list.Head.Next == nil && list.Tail.Next == nil) {
    list.Length--
    var popNode = list.Tail
    list.Head = nil
    list.Tail = nil
    return popNode.Value
  }
  
  //more than one node
  list.Length--
  var popNode = list.Tail
  list.Tail.Prev.Next = nil
  list.Tail = list.Tail.Prev
  list.Tail.Next = nil

  popNode.Prev = nil // free
  return popNode.Value
}

func (list *LinkedList) Deque() int {
  
  //no nodee
  if(list.Head == nil && list.Tail == nil) {
    println("empty list")
    return -1
  }

  //single node
  if(list.Head.Next == nil && list.Tail.Next == nil) {
    list.Length--
    var popNode = list.Head
    list.Head = nil
    list.Tail = nil
    return popNode.Value
  }

  //more than one node
  list.Length--
  var popNode = list.Head
  list.Head.Next.Prev = nil
  list.Head = list.Head.Next
  list.Head.Prev = nil
 
  popNode.Next = nil // free
  return popNode.Value
}

func (list *LinkedList) RemoveAt(position int) int  {
  
  //no node
  if(list.Head == nil && list.Tail == nil) {
    println("empty list")
    return -1
  }

  //validate the position
  if(position > list.Length || position < 1) {
    println("wrong position")
    return -1
  }

  if(position == 1) {
    return list.Deque()
  } else if (position == list.Length) {
    return list.Pop()
  }

  current := list.Head
  for i:=1 ; i<position ; i++ {
    current = current.Next
  }
 
  list.Length--
  var popNode = current
  current.Prev.Next = current.Next
  current.Next.Prev = current.Prev

  //free
  current.Next = nil
  current.Prev = nil

  return popNode.Value
}
