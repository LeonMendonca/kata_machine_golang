package main

import (
  "fmt"
  "tree/treemethods"
)

func main() {
  fmt.Println("tree traversal")
  root := &treemethods.Treenode{}
  lchild := &treemethods.Treenode{}
  rchild := &treemethods.Treenode{}

  root.Data = 31
  root.TLeft = lchild
  root.TRight = rchild

  lchild.Data = 12
  lchild.TLeft = nil
  lchild.TRight = nil

  rchild.Data = 33
  rchild.TLeft = nil
  rchild.TRight = nil

  //var path []int
  
  //run these function below one at a time, since ive used pointers to the path array!

  //root.PreOrder(&path)
  //root.PostOrder(&path)
  //root.InOrder(&path)
  //fmt.Println(path)

  //build tree with an array
  //tree := []int{1,2,3,4,5,6,7}
  //root2 := treemethods.BinaryTree(tree, 0, 7) //array, start idx, length
  //root2.InOrder(&path)
  //fmt.Println(path)
  root.Bfs(2) 
}
