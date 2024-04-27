//Primeagen referred curr and end with Point datatype, i.e curr.x, curr.y and end.x, end.y 
//i dont know about Point in GO lang, therefore ive preferred row, col for the curr(start), and end
package main

import (
 "fmt"
)

func fourcheck(arr [][]string, sr,sc,er,ec int, seen *[][]int, path *[]string) bool {
  //up down left right
  dir := [][]int{ {-1,0},{1,0},{0,-1},{0,1} }
  //outside the map
  //sr start row - single array inside an outer array
  //sc start col - each element of an array inside an outer array

  //BASE CASES
  if( sr<0 || sr >= len(arr) || sc<0 || sc >=len(arr[0])) {
    println("out of the map")
    return false
  }
  if ( arr[sr][sc] == "#" ) {
    println("current on the wall")
    return false
  }
  if ( sr==er && sc==ec ) {
    //assign the end we visit
    (*seen)[er][ec] = 1
    (*path) = append((*path),arr[er][ec])
    //println("reached end")
    return true
  }
  if ( (*seen)[sr][sc] == 1 ) {
    println("already visited")
    return false
  }

  //main recursion
  (*seen)[sr][sc] = 1
  //push the main path and the not needed
  (*path) = append((*path),arr[sr][sc])
  for i:=0 ; i<len(dir) ; i++ {
    dirin := dir[i]
    if(fourcheck(arr,sr+dirin[0],sc+dirin[1],er,ec,seen,path)) {
       return true
    }
  }
  //use slice to pop the not needed
  (*path) = (*path)[:len(*path)-1]
  return false
}

func main() {
  maze := [][]string{
  //#:walls x:path not needed o:main path
	{"#","#","E"},
	{"o","o","o"},
	{"S","x","#"},
  }
  row := len(maze)
  col := len(maze[0])
  fmt.Println(row,col)

  point := []string{}
  //add false
  boolean := make([][]int,row)
  for i := range boolean {
    boolean[i] = make([]int, col)
    for j := range boolean[i] {
       boolean[i][j] = 0
    }
  }

  //print maze
  for _,val := range maze {
    for _,inval := range val {
      fmt.Print(inval," ")
    }
    println()
  }
  //1st function call
  if(fourcheck(maze,2,0,0,2,&boolean,&point)) {
    fmt.Println("Found End")
  } else {
    fmt.Println("Fail :(")
  }
  for i := range point {
    fmt.Print(point[i]," ")
  }
  println()


  println("visited path")
  for _,val := range boolean {
    for _,inval := range val {
      fmt.Print(inval," ")
    }
    println()
  }

}
