package main

import "fmt"

var HeapSize = 15

func main() {
	data := [15]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	maxHeap(&data)
	fmt.Println(data)
	for range data{
		deleteMaxHeap(&data)
	}
}

// output tree
//                  15
//        11                     14
//   9          10          13          7
// 8    4     2    5     12    6      3   1

func maxHeap(arr *[15]int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapify(arr, i)
	}
}

func heapify(arr *[15]int, i int) {
	if i > HeapSize/2-1 {
		return
	}

	lc := arr[2*i+1]
	rc := arr[2*i+2]
	if arr[i] > lc && arr[i] > rc {
		return
	}

	if arr[i] < lc && lc > rc {
		temp := arr[i]
		arr[i] = lc
		arr[2*i+1] = temp
		heapify(arr, 2*i+1)
	} else if arr[i] < rc && rc > lc {
		temp := arr[i]
		arr[i] = rc
		arr[2*i+2] = temp
		heapify(arr, 2*i+2)
	}
}

func deleteMaxHeap(arr *[15]int){
	fmt.Printf("%d -> ", arr[0])
	HeapSize--

	if HeapSize == 0 {
		return
	}

	arr[0] = arr[HeapSize]
	heapify(arr, 0)
}