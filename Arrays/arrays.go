package main

import ("fmt")

func IndexCheck(in, length int) (bool) {
    //if index is greater than equal to length then false
    if(in >= length) {
        return true
    }
    return false
}

func getIndex(arr []int, index int) (int, bool) {
    length := len(arr)
    if( IndexCheck(index, length) ) {
        println("out of bound")
        return 0, false
    }
    el := arr[index] 
    return el, true
}

//Nothing as insertion, but update :)
//mutable array in GO, so no pointers required
func updateElement(arr []int, index, newval int) (bool) {
    length := len(arr)
    if( IndexCheck(index, length) ) {
        println("out of bound")
        return false
    }
    arr[index] = newval
    return true
}

func deletionElement(arr []int, index int) (bool) {
    length := len(arr)
    if( IndexCheck(index, length) ) {
        println("out of bound")
        return false
    }
   
    arr[index] = 0
    return true
}

func main() {
    arr := []int{1,2,3,4}
    //value,result := getIndex(arr,4)
    //fmt.Println(value,result)
    result := updateElement(arr,1,77)
    fmt.Println(result)
    fmt.Println(arr)
}

