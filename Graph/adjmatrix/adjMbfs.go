package adjmatrix

import (
  //"fmt"
)

func (adjmat *AdjMatrix) Bfs (source, needle int) {
  length := len(adjmat.Vertices)
  seen := make([]bool, length) //default false
  prev := make([]int, length)
  for idx := range prev {
    prev[idx] = -1
  }
  seen[source] = true 
  //create a queue
  q := []int{source} //push the first vertex
  for len(q) > 0 {
    currVertex := q[0]
    q = q[1:] // pop the first vertex
    if(currVertex == needle) {
      Path(prev,needle)
      break
    }
    //if vertex not found
    currVertexRel := adjmat.Vertices[currVertex]
    for idx := range currVertexRel {
      if( currVertexRel[idx] == 0 || seen[idx] ) {
        
        continue
      }
      seen[idx] = true
      prev[idx] = currVertex
      q = append(q,idx)
    }
  }
}
