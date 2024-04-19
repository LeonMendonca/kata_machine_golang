package main

import (
  "fmt"
  "stack/StackMethods"
)

func main() {
    Stack := stack.Stack();
    for i:=1 ; i<=3 ; i++ {
        Stack.Push(i)
    }
    fmt.Println("popped element:",Stack.Pop())
    fmt.Println("top of Stack:",Stack.Peek())
    fmt.Println("len:",Stack.Length)
    Stack.Printstack()
}
