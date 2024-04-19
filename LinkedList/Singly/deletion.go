package list

func (list *linkedlist) Pop() int {
  //no node
  if(list.head == nil && list.tail == nil) {
    println("empty list")
    return -1
  }

  //single node
  if(list.head.next == nil && list.tail.next == nil) {
    var popNode = list.head
    list.head = nil
    list.tail = nil
    list.Length--
    return popNode.value
  }

  //more than one node
  var current = list.head
  for current.next.next  != nil {
    current = current.next
  }
  var popNode = current.next
  current.next = current.next.next
  list.Length--

  return popNode.value
}

func (list *linkedlist) Deque() int {
  //no node
  if(list.head == nil && list.tail == nil) {
    println("empty list")
    return -1
  }

  //single node, update the tail
  if(list.head.next == nil) {
    list.tail = nil
  }

  //more than one node
  var popNode = list.head
  list.head = list.head.next
  list.Length--

  popNode.next = nil //free
  return popNode.value
}

func (list *linkedlist) RemoveAt(position int) int {
  //no node
  if(list.head == nil && list.tail == nil) {
    println("empty list")
    return -1
  }
 
  //single node
  if(list.head.next == nil && list.tail.next == nil) {
    var popNode = list.head
    list.head = nil
    list.tail = nil
    list.Length--

    return popNode.value
  }

  //more than one node
  if(position == 1) {
    var popNode = list.head
    list.head = list.head.next
    list.Length--

    popNode.next = nil //free
    return popNode.value
  }
  
  current := list.head
  for i:=1 ; i<position-1 && current != nil ; i++ {
    current = current.next
  }

  if(current.next == nil || position < 1) {
    println("wrong position")
    return 0
  }

  var popNode = current.next
  current.next = current.next.next
  list.Length--

  return popNode.value
}
