package main

import ("fmt")

func binarySearch(arr []int, el int) bool {
    var low, mid, high int
    low = 0 //first index
    high = len(arr)-1 //last index
    for true {
        //mid = (low+high) / 2
        mid = low + (high - low) / 2
        if( arr[mid] == el) {
            println("found element",el)
            return true
        } else if ( arr[mid] < el ) {
            low = mid+1
        } else {
            high = mid-1
        }    

        //it may cross or have equal index
        if(low >= high) { break }
    }
    return false
}

func main() {
    arr := []int{1,4,6,9,15}
    result := binarySearch(arr,8)
    fmt.Println(result)
}
