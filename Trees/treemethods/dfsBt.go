package treemethods

func (treeNode *Treenode) DfsBt(value int) bool {
  if(treeNode == nil) {
    return false
  }
  if(treeNode.Data == value) {
    return true
  }
  if(treeNode.Data < value) {
    return treeNode.TRight.DfsBt(value)
  }
  return treeNode.TLeft.DfsBt(value)
}
