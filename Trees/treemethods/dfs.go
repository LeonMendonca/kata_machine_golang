package treemethods

func (treeNode *Treenode) Dfs(value int) bool {
  if(treeNode == nil) {
    return false
  }
  if(treeNode.Data == value) {
    return true
  }
  if(treeNode.TLeft.Dfs(value)) {
    return true
  } else if(treeNode.TRight.Dfs(value)) {
    return true
  }
  return false
}
