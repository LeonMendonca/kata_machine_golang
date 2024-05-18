package treemethods

import (
  "math"
)

func (treeNode *Treenode) Height(l, r int) int {
  if(treeNode == nil) {
    return 0
  }
  if(treeNode.TLeft != nil) {
    l++
    return  treeNode.TLeft.Height(l, r)
  } else if (treeNode.TRight != nil) {
    r++
    return treeNode.TRight.Height(l, r)
  }
  return int(math.Max(float64(l),float64(r)))
}
