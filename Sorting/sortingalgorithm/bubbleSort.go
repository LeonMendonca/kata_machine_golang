package sortingalgorithm



//you can use this instead of Swap()
/*
temp := arr[j]
arr[j] = arr[j+1]
arr[j+1] = temp
*/
 
func BubbleSort(arr []int) {
    size := len(arr)
    for i:=0; i<size ; i++ {
        for j:=0; j<size-1-i ; j++ {
            if(arr[j] > arr[j+1]) {
                Swap(arr, j, j+1)
           }
        }
    }
}
