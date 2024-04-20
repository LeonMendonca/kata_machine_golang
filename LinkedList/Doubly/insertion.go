package doublylist

func (list *LinkedList) Append(item int) {
  NewNode := &Node{Value:item,Next:nil,Prev:nil}

  list.Length++

  if(list.Head == nil && list.Tail == nil) {
    list.Head = NewNode
    list.Tail = NewNode
    return
  }

  NewNode.Prev = list.Tail
  list.Tail.Next = NewNode
  list.Tail = NewNode

}

func (list *LinkedList) Prepend(item int) {
  NewNode := &Node{Value:item,Next:nil,Prev:nil}

  list.Length++

  if(list.Head == nil && list.Tail == nil) {
    list.Head = NewNode
    list.Tail = NewNode
    return 
  }

  NewNode.Next = list.Head
  list.Head.Prev = NewNode
  list.Head = NewNode

}

func (list  *LinkedList) InsertAt(item, position int) {

  NewNode := &Node{Value:item,Next:nil,Prev:nil}

  if(list.Head == nil && list.Tail == nil) {
    list.Head = NewNode
    list.Tail = NewNode
    list.Length++
    return 
  }

  //validate the position
  if(position > list.Length+1 || position < 1) {
    println("wrong position") 
    return
  }
  
  if(position == 1) {
    list.Prepend(item)
  } else if (position == list.Length+1) {
    list.Append(item)
  } else {

    list.Length++

    current := list.Head
    for i:=1 ; i<position ; i++ {
      current = current.Next 
    }

    NewNode.Next = current
    NewNode.Prev = current.Prev

    current.Prev.Next = NewNode
    current.Prev = NewNode
  }

}
