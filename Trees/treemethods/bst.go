package treemethods

func (treeNode *Treenode) Bst(value int) bool {
  if(treeNode == nil) {
    return false
  }
  if(treeNode.Data == value) {
    return true
  }
  if(treeNode.Data < value) {
    return treeNode.TRight.Bst(value)
  }
  return treeNode.TLeft.Bst(value)
}
