/*
terms used :
length - count of elements
capacity - length of the array
*/
package main

import("fmt")

type arrayList struct {
    array []int
    length, capacity int
}

func (arrlt *arrayList) printArray() {
    for i:=0 ; i<arrlt.length ; i++ {
        fmt.Print(arrlt.array[i]," ")
    }
    println()
}

func arrayListConstructor() *arrayList {
    //intial : length = 0; capacity = 5

    /*
    You could also get the capacity by len(), but it iterates on each element
    Its better to define the value
    */
    return &arrayList{ make([]int,5), 0, 5}
}

func (arrlt *arrayList) get(idx int) int {
    if(idx >= arrlt.length) {
        return -1 
    }
    return arrlt.array[idx]
}

func (arrlt *arrayList) push(item int) {
    if(arrlt.capacity == arrlt.length) {
        temparr := arrlt.array
        arrlt.capacity = arrlt.capacity+5
        arrlt.array = make([]int, arrlt.capacity)
        copy(arrlt.array,temparr) 
    }
    arrlt.array[arrlt.length] = item
    arrlt.length++
}

func (arrlt *arrayList) pop() int {
    if(arrlt.length == 0) {
        return -1
    }
    arrlt.length--
    delElement := arrlt.array[arrlt.length]
    arrlt.array[arrlt.length] = 0
    return delElement
}

func main() {
    arraylist := arrayListConstructor() //creates an array with length 5, and no of elements 0
    //pushing elements
    for i:=1 ; i<=7 ; i++ {
        arraylist.push(i)
    }
    //fmt.Println(arraylist.pop())
    arraylist.printArray()
    fmt.Println("length:",arraylist.length,"capacity:",arraylist.capacity)
    //fmt.Println(arraylist.get(9))
}
