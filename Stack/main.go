package main

import (
  "fmt"
  stack "stack/stackmethods"
)

func main() {
    fmt.Print("Stack\n\n")
    Stack := stack.Stack();
    /*
    Stack.Push(3)
    fmt.Println("popped element:",Stack.Pop())
    fmt.Println("top of Stack:",Stack.Peek())
    fmt.Println("length:",Stack.Length)
    */
    Stack.Printstack()
}
