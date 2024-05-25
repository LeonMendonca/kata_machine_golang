package treemethods

func (treeNode *Treenode) Bfs(searchVal int) bool {
  if(treeNode == nil) {
    return false
  }
  //declare an array of treenode struct type
  q := []*Treenode{treeNode}
  for len(q) > 0 {
    node := q[0] // save the first node
    q = q[1:] //remove the first node using slice
    print(node.Data," ")
    if(node.Data == searchVal) {
      return true
    }
    if(node.TLeft != nil) {
      q = append(q,node.TLeft)
    }
    if(node.TRight != nil) {
      q = append(q,node.TRight)
    }
  }
  return true
}
