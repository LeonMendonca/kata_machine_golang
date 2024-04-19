package arraylist

func (arrlt *arrayList) Pop() int {
    if(arrlt.Length == 0) {
        return -1
    }
    arrlt.Length--
    delElement := arrlt.array[arrlt.Length]
    arrlt.array[arrlt.Length] = 0
    return delElement
}
