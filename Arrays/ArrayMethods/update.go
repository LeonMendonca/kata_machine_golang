package array

//Nothing as insertion, but update :)
func UpdateElement(arr []int, index, newval int) (bool) {
    length := len(arr)
    if( index < length ) {
      arr[index] = newval
      return true
    }
    println("out of bound")
    return false
}
