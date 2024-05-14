package btree

func(treeNode *Treenode) InOrder(path *[]int) []int { 
  //pre
  if(treeNode == nil) {
    return *path
  }
  
  //recurse
  treeNode.TLeft.InOrder(path)
  *path = append(*path,treeNode.Data)
  treeNode.TRight.InOrder(path)

  //post
  return *path
}
