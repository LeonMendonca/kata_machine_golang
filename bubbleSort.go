package main

import ("fmt")

func Swap(arr []int, leftIndex, rightIndex int) {
    temp := arr[rightIndex]
    arr[rightIndex] = arr[leftIndex]
    arr[leftIndex] = temp
}

//you can use this instead of Swap()
/*
temp := arr[j]
arr[j] = arr[j+1]
arr[j+1] = temp
*/
 
func bubbleSort(arr []int) {
    size := len(arr)
    for i:=0; i<size ; i++ {
        for j:=0; j<size-1-i ; j++ {
            if(arr[j] > arr[j+1]) {
                Swap(arr, j, j+1)
           }
        }
    }
}

func main() {
    arr:=[]int{7,3,5,2,8}
    bubbleSort(arr)
    fmt.Println(arr)
}
