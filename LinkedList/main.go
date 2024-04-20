package main

import (
  "fmt"
  "linkedlist/Singly"
  "linkedlist/Doubly"
)

func main() {
  fmt.Print("linked list\n\n")

  //SINGLY

  //constructor call
  SinglyLinkedlist := singlylist.SinglyLinkedList()

/*
  SinglyLinkedlist.Append(56)
  SinglyLinkedlist.Prepend(34)
  SinglyLinkedlist.InsertAt(99,5)
*/
/*
  fmt.Println("Popped Node:",SinglyLinkedlist.Pop())
  fmt.Println("Popped Node:",SinglyLinkedlist.Deque())
  fmt.Println("Popped Node:",SinglyLinkedlist.RemoveAt(3))
*/

  SinglyLinkedlist.PrintList()
  fmt.Println("length of SinglyLinkedlist",SinglyLinkedlist.Length)

  
  //DOUBLY
  
  //constructor call
  DoublyLinkedlist := doublylist.DoublyLinkedList()
/*
  DoublyLinkedlist.Append(1)
  DoublyLinkedlist.Prepend(11)
  DoublyLinkedlist.InsertAt(44,4) 
*/
/*
  fmt.Println(DoublyLinkedlist.Pop())
  fmt.Println(DoublyLinkedlist.Deque())
  fmt.Println(DoublyLinkedlist.RemoveAt(2))
*/

  DoublyLinkedlist.PrintDoublyLinkedListNext()
  DoublyLinkedlist.PrintDoublyLinkedListPrev()
  fmt.Println("Length of DoublyLinkedlist",DoublyLinkedlist.Length)

}
