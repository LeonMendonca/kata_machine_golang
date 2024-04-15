package main

import ("fmt")

func linearSearch(arr []int, el int) bool {
    for _,eachEl := range arr {
        if(eachEl == el) {
            println("found value",el)
            return true
        }
    }
    return false
}

func main() {
    arr := []int{2,4,1,7,5}
    result := linearSearch(arr,7)
    fmt.Println(result)
}
