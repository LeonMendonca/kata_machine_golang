package list

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

  //update the head
  if(position == 1) {
    NewNode.next = list.head 
    list.head = NewNode
    list.Length++
    return
  }
 
  //iterate from head Node to the Position specified -1
  current := list.head
  for i:=1 ; i < position-1 && current != nil ; i++ {
    current = current.next
  }
  //fmt.Println(current) //debugging print
  if(current == nil || position < 1) {
    println("wrong position")
    return
  }
 
  //if last Node, then update the tail
  if(current.next == nil) {
    list.tail = NewNode
  }

  NewNode.next = current.next
  current.next = NewNode
  list.Length++
}
