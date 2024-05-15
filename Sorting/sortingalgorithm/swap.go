package sortingalgorithm

func Swap(arr []int, leftIndex, rightIndex int) {
    temp := arr[rightIndex]
    arr[rightIndex] = arr[leftIndex]
    arr[leftIndex] = temp
}
