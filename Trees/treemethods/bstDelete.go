package treemethods

func (treeNode *Treenode) BstDelete(value int) *Treenode {
  if(treeNode == nil) {
    return treeNode
  }
  if(treeNode.Data > value) {
    treeNode.TLeft = treeNode.TLeft.BstDelete(value)
  } else if(treeNode.Data < value) {
    treeNode.TRight = treeNode.TRight.BstDelete(value)
  } else {
    //if node found

    //case 1 and 2 : no children/1 child
    if(treeNode.TLeft == nil) {
      //println("left nil")
      return treeNode.TRight
    } else if(treeNode.TRight == nil) {
      //println("right nil")
      return treeNode.TLeft
    }
    //case 3
    twoChild(treeNode)
  }
  
  return treeNode
}

func twoChild(treeNode *Treenode) {
}
