package treemethods

func Btcomp(treeA, treeB *Treenode) bool {
  //println(treeA.Data,treeB.Data)
  if(treeA == nil && treeB == nil) {
    println("true")
    return true
  }
  if(treeA == nil || treeB == nil) {
    println("false 1")
    return false
  }
  if(treeA.Data != treeB.Data) {
    println("false 2")
    return false
  }
  return Btcomp(treeA.TLeft,treeB.TLeft) && Btcomp(treeA.TRight, treeB.TRight)
}
