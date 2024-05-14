package main

import (
  "fmt"
  "tree/treeMethods"
)

func main() {
  fmt.Println("tree traversal")
  root := &btree.Treenode{}
  lchild := &btree.Treenode{}
  rchild := &btree.Treenode{}

  root.Data = 31
  root.TLeft = lchild
  root.TRight = rchild

  lchild.Data = 12
  lchild.TLeft = nil
  lchild.TRight = nil

  rchild.Data = 33
  rchild.TLeft = nil
  rchild.TRight = nil

  var path []int
  
  //run these function one at a time, since ive used pointers to the path array!

  //root.PreOrder(&path)
  //root.PostOrder(&path)
  //root.InOrder(&path)
  //fmt.Println(path)

  //build tree with an array
  tree := []int{1,2,3,4,5,6,7}
  root2 := btree.BinaryTree(tree, 0, 7)
  root2.InOrder(&path)
  fmt.Println(path)
}
