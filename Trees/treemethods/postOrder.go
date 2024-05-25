package treemethods

func(treeNode *Treenode) PostOrder(path *[]int) []int { 
  //pre
  if(treeNode == nil) {
    return *path
  }
  
  //recurse
  treeNode.TLeft.PostOrder(path)
  treeNode.TRight.PostOrder(path)

  //post
  *path = append(*path,treeNode.Data)
  return *path
}
