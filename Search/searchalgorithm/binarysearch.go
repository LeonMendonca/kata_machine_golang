package searchalgorithm

func BinarySearch(arr []int, el int) bool {
    var low, mid, high int
    low = 0 //first index
    high = len(arr)-1 //last index
    for {
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

        //if they both cross
        if(low > high) { break }
    }
    return false
}
