package treemethods

type Treenode struct {
  Data int
  TLeft, TRight *Treenode
}

//constructor - it builds a tree
func BinaryTree(nodes []int, idx, length int) *Treenode {
  if idx >= length {
    return nil
  }
  //pre
  node := &Treenode{Data:nodes[idx]}

  //recurse
  node.TLeft = BinaryTree(nodes,2*idx+1, length)
  node.TRight = BinaryTree(nodes,2*idx+2, length)

  //post
  return node
  //return &Treenode{Data:0,TLeft:nil,TRight:nil}
}
