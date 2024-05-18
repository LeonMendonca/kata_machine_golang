package treemethods

func (treeNode *Treenode) Dfs(value int) *Treenode {
  var foundNode *Treenode
  if(treeNode == nil) {
    return treeNode
  }
  if(treeNode.Data == value) {
    return treeNode
  }
  foundNode = treeNode.TLeft.Dfs(value)
  foundNode = treeNode.TRight.Dfs(value)

  return foundNode
}
