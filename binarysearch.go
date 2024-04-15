package main

import ("fmt")

func binarySearch(arr []int, el int) bool {
    var low, mid, high int
    low = 0 //first index
    high = len(arr)-1 //last index
    for {
        //mid = (low+high) / 2
        mid = low + (high - low) / 2
        fmt.Println("low:",low,"mid:",mid,arr[mid],"high",high)
        if( arr[mid] == el) {
            println("found element",el)
            return true
        } else if ( arr[mid] < el ) {
            low = mid+1
        } else {
            high = mid-1
        }    

        //it may cross or have equal index
        if(low > high) { break }
    }
    return false
}

func main() {
    arr := []int{1,4,6,9,15,17,20,22,27,32,35,37,38,49,52,105,107,171,200,201,202,301,444,555}
    result := binarySearch(arr,35)
    fmt.Println(len(arr))
    fmt.Println(result)
}
