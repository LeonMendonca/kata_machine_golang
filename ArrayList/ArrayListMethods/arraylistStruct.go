package arraylist

type arrayList struct {
    array []int
    Length, Capacity int
}

//constructor
func ArrayList() *arrayList {
    //intial : length = 0; capacity = 5

    /*
    You could also get the capacity by len(), but it iterates on each element
    Its better to intialize the length
    */
    return &arrayList{
      array:make([]int,5), 
      Length:0, 
      Capacity:5,
    }
}


