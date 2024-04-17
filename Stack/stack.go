package main

import(
    "fmt"
    "math"
)

type node struct {
    value int
    prev *node
}

type Stack struct {
    head *node
    length int
}

func stackConstructor() *Stack {
    return &Stack{head:nil,length:0} 
}

func (s *Stack) PrintStack() {
    curr := s.head
    for curr != nil {
        fmt.Println("|_",curr.value,"_|")
        curr = curr.prev
    }
}

func (s *Stack) push(item int) {
    newNode := &node{value:item, prev:nil}
    if(s.head == nil) {
        s.length++
        s.head = newNode
        return
    }

    newNode.prev = s.head
    s.head = newNode
    s.length++
}

func (s *Stack) pop() int {
    //if you try to pop an empty stack, this throws a runtime error
    s.length = int( math.Max(0, float64(s.length-1)) )
    if( s.length == 0 ) {
        top := s.head
        s.head = nil
        return top.value
    }

    //use this instead. uncomment the below
    /*
    if(s.head == nil) {
        println("nothing to pop")
        return 0
    }
    */
    
    top := s.head
    s.head = s.head.prev
    //s.length--

    return top.value
}

func (s *Stack) peek() int {
    if(s.head != nil) {
        return s.head.value 
    }
    return 0
}

func main() {
    Stack := stackConstructor();
    for i:=1 ; i<=3 ; i++ {
        Stack.push(i)
    }
    fmt.Println("popped element:",Stack.pop())
    fmt.Println("top of stack:",Stack.peek())
    fmt.Println("len:",Stack.length)
    Stack.PrintStack()
}
