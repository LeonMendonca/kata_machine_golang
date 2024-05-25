package treemethods

func (treeNode *Treenode) BtreeInsert(value int) *Treenode {
  node := &Treenode{Data:value, TLeft:nil, TRight:nil}
  if(treeNode == nil) {
    return treeNode
  }
  if(treeNode.Data < value) {
    if(treeNode.TRight == nil) {
      treeNode.TRight = node
    } else {
      treeNode.TRight.BtreeInsert(value)
    }
  } else if (treeNode.Data >= value) {
    if(treeNode.TLeft == nil) {
      treeNode.TLeft = node
    } else {
      treeNode.TLeft.BtreeInsert(value)
    }
  }
  return treeNode
}
