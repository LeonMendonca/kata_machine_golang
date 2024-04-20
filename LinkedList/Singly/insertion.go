package singlylist

func (list *linkedlist) Append(item int) {
  NewNode := &Node{value:item,next:nil} 

  if(list.head == nil && list.tail == nil) {
    list.head = NewNode
    list.tail = NewNode
    list.Length++
    return
  }
  
  list.tail.next = NewNode
  list.tail = NewNode
  list.Length++

}

func (list *linkedlist) Prepend(item int) {
  NewNode := &Node{value:item,next:nil} 

  if(list.head == nil && list.tail == nil) {
    list.head = NewNode
    list.tail = NewNode
    list.Length++
    return
  }
 
  NewNode.next = list.head
  list.head = NewNode
  list.Length++

}

func (list *linkedlist) InsertAt(item, position int) {

  NewNode := &Node{value:item,next:nil} 

  if(list.head == nil && list.tail == nil) {
    list.head = NewNode
    list.tail = NewNode
    list.Length++
    return
  }

  if(position > list.Length+1 || position < 1) {
    println("wrong position")
    return
  }

  if(position == 1) {
    list.Prepend(item)
  } else if (position == list.Length+1) {
    list.Append(item)
  } else {
    //iterate from head Node to the Position specified -1
    current := list.head
    for i:=1 ; i < position-1 ; i++ {
      current = current.next
    }
    NewNode.next = current.next
    current.next = NewNode
    list.Length++
  }

}
