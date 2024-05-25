package arraylist

func (arrlt *arrayList) Push(item int) {
    if(arrlt.Capacity == arrlt.Length) {
        temparr := arrlt.array
        arrlt.Capacity = arrlt.Capacity+5
        arrlt.array = make([]int, arrlt.Capacity)
        copy(arrlt.array,temparr) 
    }
    arrlt.array[arrlt.Length] = item
    arrlt.Length++
}
