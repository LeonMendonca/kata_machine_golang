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

  var path []int
  
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

  //Breadth first search
  //root.Bfs(2) 

  //Binary tree comparison
  /*
  arr1 := []int{1,2,3}
  arr2 := []int{1,2,34}
  */
  //treeA := treemethods.BinaryTree(arr1,0,3)
  //treeB := treemethods.BinaryTree(arr2,0,3)
  //var result = treemethods.Btcomp(treeA,treeB)

  //Bt insert and delete
  bst := &treemethods.Treenode{} //decalred and allocated memory
  bst.BtreeInsert(2)
  bst.BtreeInsert(4)
  bst.BtreeInsert(11)
  var left, right int
  fmt.Println("h",bst.Height(left,right))

  bst.BstDelete(11)
  fmt.Println(bst.PreOrder(&path))


}
