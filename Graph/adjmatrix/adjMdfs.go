package adjmatrix

func (adjmat *AdjMatrix) Dfs (vertex, needle int) bool {
  length := len(adjmat.Vertices)
  seen := make([]bool, length) //default false
  prev := make([]int, length)
  for idx := range prev {
    prev[idx] = -1
  }

  return Dfs(adjmat, vertex, needle, seen, prev)
}

func Dfs (adjmat *AdjMatrix, vertex, needle int, seen []bool, prev []int) bool {

  seen[vertex] = true

  if(vertex == needle) {
    Path(prev,needle)
    return true
  }
   
  currVertexRel := adjmat.Vertices[vertex]
  for idx := range currVertexRel {
    if( currVertexRel[idx] == 0 || seen[idx] ) {
      continue
    }
    prev[idx] = vertex
    if( Dfs(adjmat, idx, needle, seen, prev) ) {
      return true
    }

  }
  return false
}
