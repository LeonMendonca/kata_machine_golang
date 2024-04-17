package main

import ("fmt";"math")

func twoCB(arr []int) int {
    size := len(arr) 
    jumplen := int(math.Round(math.Sqrt(float64(size))))
    /*
    As said we need Square root of the length of the array, but GO expects a 
    float number, hence typecast it to float64
    After this, use Round to exclude the decimal digits
    */
    i := 0 //declared in a global scope

    //first floor is 1
    if( arr[0] == 1) {
        return 0        
    }

    for ; i<size ; i+=jumplen { 
        if( arr[i] == 1 ) {
            break
        }
    }
   
    i-=jumplen

    for j := 0 ; j<=jumplen && i < size ; {
        if( arr[i] == 1 ) {
            return i
        }
        i++
        j++
    }

    return -1

}

func main() {
    var bldg = []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1}
    floor := twoCB(bldg) 
    fmt.Println(floor)
}
