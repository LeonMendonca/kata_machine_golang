package btree


func (treeNode *Treenode) PreOrder(path *[]int) []int {
  if(treeNode == nil) {
    return *path
  }
  //pre
  *path = append(*path,treeNode.Data)
  
  //recurse
  *path = treeNode.TLeft.PreOrder(path)
  *path = treeNode.TRight.PreOrder(path)

  //post
  return *path
}

