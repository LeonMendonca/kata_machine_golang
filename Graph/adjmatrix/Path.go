package adjmatrix

func Path(prev []int, vertex int) {
  if(vertex == -1) {
    return
  }
  Path(prev,prev[vertex])
  print(vertex," ")
}
