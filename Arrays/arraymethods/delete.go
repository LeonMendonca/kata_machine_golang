package arraymethods

func DeleteElement(arr []int, index int) (int, bool) {
    length := len(arr)
    if( index < length ) {
      var deletedElement = arr[index]
      arr[index] = 0
      return deletedElement,true
    }
    println("out of bound")
    return 0,false
}
