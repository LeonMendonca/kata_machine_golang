//Ive used Recursion to find the path backwardly
package main

import (
  "fmt"
  "graph/adjmatrix"
)

func main() {
  fmt.Println("graph")
  /*
  Graph used here (uni directed)
  0 ---- 1
  |      |
  |      |
  3 ---- 2
  */
  AdjMat := &adjmatrix.AdjMatrix{Vertices:[][]int{
    {0, 1, 0, 1} , //0 vertex
    {1, 0, 1, 0} , //1 vertex
    {0, 1, 0, 1} , //2 vertex
    {1, 0, 1, 0} , //3 vertex
  }}
  fmt.Println(AdjMat)
  //AdjMat.Bfs(0,2)
  AdjMat.Dfs(0,0)

}
