//Bonus
/*
with bubble sort, I feel that selection sort is also equally important. Since, 
they are kinda identical and have same time complexity O(n^2)

Do refer this site - https://www.geeksforgeeks.org/selection-sort/

Happy Coding! :)
*/
package main

import ("fmt")

/*
You can try to understand this first!

this finds the smallest value in the array and returns the index of it 
(inner loop of the selection sort)
*/
func small(arr []int) int {
    small := 0 ; size := len(arr)
    for i:=1 ; i<size ; i++ {
        if(arr[small] > arr[i]) {
            small = i
        }
    }
    return small
}

func Swap(arr []int, leftIndex, rightIndex int) {
    temp := arr[rightIndex]
    arr[rightIndex] = arr[leftIndex]
    arr[leftIndex] = temp
}

func selectionSort(arr []int) {
    size := len(arr)
    for i:=0 ; i<size ; i++ {
        small := i
        for j:=1+i ; j<size ; j++ {
            if(arr[small] > arr[j]) {
                small = j
            }
        }
        /* i and small are swapped but if both have same index, there is no point
        to swap */
        if(small != i) {
            Swap(arr,small,i)
        }
   }
}

func main() {
    arr := []int{8,5,1,2,7}
    //fmt.Println(small(arr))
    selectionSort(arr)
    fmt.Println(arr)
}
