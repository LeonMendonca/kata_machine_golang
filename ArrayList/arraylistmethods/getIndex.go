package arraylist

//get element by the index
func (arrlt *arrayList) GetIndex(idx int) int {
    if(idx >= arrlt.Length) {
        return -1 
    }
    return arrlt.array[idx]
}
