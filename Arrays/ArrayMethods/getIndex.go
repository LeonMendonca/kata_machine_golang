package array

func GetIndex(arr []int, index int) (int, bool) {
    length := len(arr)
    if( index < length ) {
      var Element = arr[index] 
      return Element, true
    }
    println("out of bound")
    return 0, false
}
