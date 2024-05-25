package searchalgorithm

func LinearSearch(arr []int, el int) bool {
    for _,eachEl := range arr {
        if(eachEl == el) {
            println("found value",el)
            return true
        }
    }
    return false
}
