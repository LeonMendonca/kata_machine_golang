package treemethods

import (
  //"fmt"
)

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

    //instead of pointing the parent node, to the node where the node is the child to the delNode
    // 15 -> 7 -> 4 :: 15 -> 4 delNode - 7
    // 15 is parent, 7 is del node, 4 is child of 7

    //Let the child node replace the delNode
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
func twoChild(treeNode *Treenode)  {
  delNode := treeNode
  curr := treeNode.TLeft
  var prev *Treenode
  for (curr.TRight != nil) {
    prev = curr
    curr = curr.TRight
  }
  //fmt.Println("curr",curr)
  //fmt.Println("prev",prev)
  delNode.Data = curr.Data
  if(prev != nil) {
    prev.TRight = curr.BstDelete(curr.Data)
  } else {
    delNode.TLeft = curr.BstDelete(curr.Data)
  }
}

